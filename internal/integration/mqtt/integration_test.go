package mqtt

import (
	"os"
	"strings"
	"testing"
	"text/template"
	"time"

	"github.com/SplitStackServer/splitstack/api/go/v5/bs"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"google.golang.org/protobuf/proto"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	testStateTopicTemplate    = "test/bssci/{{ .BsEui }}/state"
	testEventTopicTemplate    = "test/bssci/{{ .BsEui }}/event/{{ .EventSource }}/{{ .EventType }}"
	testCommandTopicTemplate  = "test/bssci/{{ .BsEui }}/command/#"
	testResponseTopicTemplate = "test/bssci/{{ .BsEui }}/response/#"
)

type TestIntegrationSuite struct {
	suite.Suite

	mqttClient      paho.Client
	integration     *Integration
	basestationsEui common.EUI64
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(TestIntegrationSuite))

}

func (ts *TestIntegrationSuite) SetupSuite() {

	ts.T().Skip("fix later")
	assert := require.New(ts.T())
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	godotenv.Load("../../../.env")

	var server string
	var username string
	var password string

	if v := os.Getenv("TEST_MQTT_SERVER"); v != "" {
		server = v
	} else {
		ts.T().Skip("TEST_MQTT_SERVER is not set in env")
	}

	if v := os.Getenv("TEST_MQTT_USERNAME"); v != "" {
		username = v
	}
	if v := os.Getenv("TEST_MQTT_PASSWORD"); v != "" {
		password = v
	}

	opts := paho.NewClientOptions().AddBroker(server).SetUsername(username).SetPassword(password)
	ts.mqttClient = paho.NewClient(opts)
	token := ts.mqttClient.Connect()
	token.Wait()
	assert.NoError(token.Error())

	ts.basestationsEui = common.EUI64{8, 7, 6, 5, 4, 3, 2, 1}

	var conf config.Config
	conf.Integration.Marshaler = "json"
	conf.Integration.MQTTV3.StateRetained = true
	conf.Integration.MQTTV3.Auth.Type = "generic"
	conf.Integration.MQTTV3.Auth.Generic.Servers = []string{server}
	conf.Integration.MQTTV3.Auth.Generic.Username = username
	conf.Integration.MQTTV3.Auth.Generic.Password = password
	conf.Integration.MQTTV3.Auth.Generic.CleanSession = true
	conf.Integration.MQTTV3.Auth.Generic.ClientID = ts.basestationsEui.String()
	conf.Integration.MQTTV3.MaxTokenWait = time.Second

	var err error
	ts.integration, err = NewIntegration(conf)
	assert.NoError(err)

	ts.integration.stateTopicTemplate, _ = template.New("state").Parse(testStateTopicTemplate)
	ts.integration.eventTopicTemplate, _ = template.New("event").Parse(testEventTopicTemplate)
	ts.integration.commandTopicTemplate, _ = template.New("command").Parse(testCommandTopicTemplate)
	ts.integration.responseTopicTemplate, _ = template.New("response").Parse(testResponseTopicTemplate)

	assert.NoError(ts.integration.Start())

	// The subscribe loop runs every 100ms, we will wait twice the time to make
	// sure the subscription is set.
	time.Sleep(400 * time.Millisecond)
}

func (ts *TestIntegrationSuite) TearDownSuite() {
	ts.mqttClient.Disconnect(0)
	ts.integration.Stop()
}

func (ts *TestIntegrationSuite) TestIntegration_LastWill() {
	assert := require.New(ts.T())

	assert.True(ts.integration.clientOpts.WillEnabled)
	assert.Equal("test/bssci/0807060504030201/state", ts.integration.clientOpts.WillTopic)
	assert.Equal(`{"bsEui":"0807060504030201"}`, strings.ReplaceAll(string(ts.integration.clientOpts.WillPayload), " ", ""))
	assert.True(ts.integration.clientOpts.WillRetained)
}

func (ts *TestIntegrationSuite) TestIntegration_ConnStateOnline() {
	assert := require.New(ts.T())

	connStateChan := make(chan *bs.BasestationState)
	token := ts.mqttClient.Subscribe("test/bssci/0807060504030201/state", 0, func(c paho.Client, ms paho.Message) {
		var pl bs.BasestationState
		assert.NoError(ts.integration.unmarshal(ms.Payload(), &pl))
		connStateChan <- &pl
	})
	token.Wait()
	assert.NoError(token.Error())

	pl := <-connStateChan

	assert.True(proto.Equal(&bs.BasestationState{
		BsEui: ts.basestationsEui.String(),
		State: bs.BasestationState_ONLINE,
	}, pl))

	token = ts.mqttClient.Unsubscribe("test/bssci/0807060504030201/state")
	token.Wait()
	assert.NoError(token.Error())
}

func (ts *TestIntegrationSuite) TestIntegration_SubscribeBasestation() {
	assert := require.New(ts.T())

	bsEui := common.EUI64{1, 2, 3, 4, 5, 6, 7, 8}
	connStateChan := make(chan *bs.BasestationState)

	assert.NoError(ts.integration.SetBasestationSubscription(true, bsEui))
	_, ok := ts.integration.basestations[bsEui]
	assert.True(ok)

	// Wait 400ms to make sure that the subscribe loop has picked up the
	// change and set the ConnState. If we subscribe too early, it is
	// possible that we get an (old) OFFLINE retained message.
	time.Sleep(1000 * time.Millisecond)

	token := ts.mqttClient.Subscribe("test/bssci/0102030405060708/state", 0, func(c paho.Client, ms paho.Message) {
		var pl bs.BasestationState
		assert.NoError(ts.integration.unmarshal(ms.Payload(), &pl))
		connStateChan <- &pl
	})
	token.Wait()
	assert.NoError(token.Error())

	pl := <-connStateChan

	assert.True(proto.Equal(&bs.BasestationState{
		BsEui: bsEui.String(),
		State: bs.BasestationState_ONLINE,
	}, pl))

	ts.T().Run("Unsubscribe", func(t *testing.T) {
		assert := require.New(t)

		assert.NoError(ts.integration.SetBasestationSubscription(false, bsEui))
		_, ok := ts.integration.basestations[bsEui]
		assert.False(ok)

		pl := <-connStateChan

		assert.True(proto.Equal(&bs.BasestationState{
			BsEui: bsEui.String(),
			State: bs.BasestationState_OFFLINE,
		}, pl))
	})

	token = ts.mqttClient.Unsubscribe("test/bssci/0102030405060708/state")
	token.Wait()
	assert.NoError(token.Error())
}

func (ts *TestIntegrationSuite) TestIntegration_PublishPublishEndnodeEvent() {
	assert := require.New(ts.T())

	pb := bs.EndnodeUplink{
		BsEui: "test_bs",
		Message: &bs.EndnodeUplink_UlData{
			UlData: &bs.EndnodeUlDataMessage{
				EpEui: "test_ep",
				Data:  []byte{0, 1, 2, 3},
			},
		},
	}

	uplinkFrameChan := make(chan *bs.EndnodeUplink)
	token := ts.mqttClient.Subscribe("test/bssci/0807060504030201/event/ep/ul", 0, func(c paho.Client, ms paho.Message) {
		var pl bs.EndnodeUplink
		assert.NoError(ts.integration.unmarshal(ms.Payload(), &pl))
		uplinkFrameChan <- &pl
	})
	token.Wait()
	assert.NoError(token.Error())

	assert.NoError(ts.integration.PublishEndnodeEvent(ts.basestationsEui, "ul", &pb))
	uplinkReceived := <-uplinkFrameChan
	assert.True(proto.Equal(&pb, uplinkReceived))
}

func (ts *TestIntegrationSuite) TestIntegration_PublishBasestationEvent() {
	assert := require.New(ts.T())

	pb := bs.BasestationUplink{
		BsEui: "test_bs",
		Message: &bs.BasestationUplink_VmStatus{
			VmStatus: &bs.BasestationVariableMacStatus{
				MacTypes: []uint32{1, 2, 3},
			},
		},
	}

	uplinkFrameChan := make(chan *bs.BasestationUplink)
	token := ts.mqttClient.Subscribe("test/bssci/0807060504030201/event/ep/vm", 0, func(c paho.Client, ms paho.Message) {
		var pl bs.BasestationUplink
		assert.NoError(ts.integration.unmarshal(ms.Payload(), &pl))
		uplinkFrameChan <- &pl
	})
	token.Wait()
	assert.NoError(token.Error())

	assert.NoError(ts.integration.PublishBasestationEvent(ts.basestationsEui, "vm", &pb))
	uplinkReceived := <-uplinkFrameChan
	assert.True(proto.Equal(&pb, uplinkReceived))
}

func (ts *TestIntegrationSuite) TestIntegration_HandleServerResponse() {
	assert := require.New(ts.T())
	rspChan := make(chan *bs.ServerResponse, 1)
	ts.integration.SetServerResponseHandler(func(pl *bs.ServerResponse) {
		rspChan <- pl
	})

	response := bs.ServerResponse{
		OpId:  1,
		BsEui: "test_bs",
		Response: &bs.ServerResponse_Err{
			Err: &bs.ErrorResponse{
				Message: "test",
			},
		},
	}

	b, err := ts.integration.marshal(&response)
	assert.NoError(err)

	token := ts.mqttClient.Publish("test/bssci/0807060504030201/response/", 0, false, b)
	token.Wait()
	assert.NoError(token.Error())

	receivedResponse := <-rspChan
	assert.True(proto.Equal(&response, receivedResponse))
}

func (ts *TestIntegrationSuite) TestIntegration_HandleServerCommand() {
	assert := require.New(ts.T())
	cmdChan := make(chan *bs.ServerCommand, 1)
	ts.integration.SetServerCommandHandler(func(pl *bs.ServerCommand) {
		cmdChan <- pl
	})

	command := bs.ServerCommand{
		BsEui: "test_bs",
		Command: &bs.ServerCommand_VmStatus{
			VmStatus: &bs.RequestVariableMacStatus{},
		},
	}

	b, err := ts.integration.marshal(&command)
	assert.NoError(err)

	token := ts.mqttClient.Publish("test/bssci/0807060504030201/command/", 0, false, b)
	token.Wait()
	assert.NoError(token.Error())

	receivedResponse := <-cmdChan
	assert.True(proto.Equal(&command, receivedResponse))
}
