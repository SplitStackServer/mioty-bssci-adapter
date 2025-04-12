package mqtt

import (
	"os"
	"strings"
	"testing"
	"text/template"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"google.golang.org/protobuf/proto"

	"mioty-bssci-adapter/internal/common"
	"mioty-bssci-adapter/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/api/rsp"
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

	// override topics to make cleanup easier
	testStateTopicTemplate    := "test/" + stateTopicTemplate
	testEventTopicTemplate    := "test/" + eventTopicTemplate
	testCommandTopicTemplate  := "test/" + commandTopicTemplate
	testResponseTopicTemplate := "test/" + responseTopicTemplate

	ts.integration.stateTopicTemplate, _ = template.New("state").Parse(testStateTopicTemplate)
	ts.integration.eventTopicTemplate, _ = template.New("state").Parse(testEventTopicTemplate)
	ts.integration.commandTopicTemplate, _ = template.New("command").Parse(testCommandTopicTemplate)
	ts.integration.responseTopicTemplate, _ = template.New("state").Parse(testResponseTopicTemplate)



	assert.NoError(ts.integration.Start())

	// The subscribe loop runs every 100ms, we will wait twice the time to make
	// sure the subscription is set.
	time.Sleep(400 * time.Millisecond)
}

func (ts *TestIntegrationSuite) TearDownSuite() {
	ts.mqttClient.Disconnect(0)
	ts.integration.Stop()
}

func (ts *TestIntegrationSuite) TestLastWill() {
	assert := require.New(ts.T())

	assert.True(ts.integration.clientOpts.WillEnabled)
	assert.Equal("test/bssci/0807060504030201/state", ts.integration.clientOpts.WillTopic)
	assert.Equal(`{"bsEui":"0807060504030201"}`, strings.ReplaceAll(string(ts.integration.clientOpts.WillPayload), " ", ""))
	assert.True(ts.integration.clientOpts.WillRetained)
}

func (ts *TestIntegrationSuite) TestConnStateOnline() {
	assert := require.New(ts.T())

	connStateChan := make(chan *msg.ProtoBasestationState)
	token := ts.mqttClient.Subscribe("test/bssci/0807060504030201/state", 0, func(c paho.Client, ms paho.Message) {
		var pl msg.ProtoBasestationState
		assert.NoError(ts.integration.unmarshal(ms.Payload(), &pl))
		connStateChan <- &pl
	})
	token.Wait()
	assert.NoError(token.Error())

	pl := <-connStateChan

	assert.True(proto.Equal(&msg.ProtoBasestationState{
		BsEui: ts.basestationsEui.String(),
		State: msg.ConnectionState_ONLINE,
	}, pl))

	token = ts.mqttClient.Unsubscribe("test/bssci/0807060504030201/state")
	token.Wait()
	assert.NoError(token.Error())
}

func (ts *TestIntegrationSuite) TestSubscribeBasestation() {
	assert := require.New(ts.T())

	bsEui := common.EUI64{1, 2, 3, 4, 5, 6, 7, 8}
	connStateChan := make(chan *msg.ProtoBasestationState)

	assert.NoError(ts.integration.SetBasestationSubscription(true, bsEui))
	_, ok := ts.integration.basestations[bsEui]
	assert.True(ok)

	// Wait 400ms to make sure that the subscribe loop has picked up the
	// change and set the ConnState. If we subscribe too early, it is
	// possible that we get an (old) OFFLINE retained message.
	time.Sleep(1000 * time.Millisecond)

	token := ts.mqttClient.Subscribe("test/bssci/0102030405060708/state", 0, func(c paho.Client, ms paho.Message) {
		var pl msg.ProtoBasestationState
		assert.NoError(ts.integration.unmarshal(ms.Payload(), &pl))
		connStateChan <- &pl
	})
	token.Wait()
	assert.NoError(token.Error())

	pl := <-connStateChan

	assert.True(proto.Equal(&msg.ProtoBasestationState{
		BsEui: bsEui.String(),
		State: msg.ConnectionState_ONLINE,
	}, pl))

	ts.T().Run("Unsubscribe", func(t *testing.T) {
		assert := require.New(t)

		assert.NoError(ts.integration.SetBasestationSubscription(false, bsEui))
		_, ok := ts.integration.basestations[bsEui]
		assert.False(ok)

		pl := <-connStateChan

		assert.True(proto.Equal(&msg.ProtoBasestationState{
			BsEui: bsEui.String(),
			State: msg.ConnectionState_OFFLINE,
		}, pl))
	})

	token = ts.mqttClient.Unsubscribe("test/bssci/0102030405060708/state")
	token.Wait()
	assert.NoError(token.Error())
}

func (ts *TestIntegrationSuite) TestPublishPublishEndnodeEvent() {
	assert := require.New(ts.T())

	pb := msg.ProtoEndnodeMessage{
		BsEui:      "test_bs",
		EndnodeEui: "test_ep",
		V1: &msg.ProtoEndnodeMessage_UlData{
			UlData: &msg.EndnodeUlDataMessage{
				Data: []byte{0, 1, 2, 3},
			},
		},
	}

	uplinkFrameChan := make(chan *msg.ProtoEndnodeMessage)
	token := ts.mqttClient.Subscribe("test/bssci/0807060504030201/event/ep/ul", 0, func(c paho.Client, ms paho.Message) {
		var pl msg.ProtoEndnodeMessage
		assert.NoError(ts.integration.unmarshal(ms.Payload(), &pl))
		uplinkFrameChan <- &pl
	})
	token.Wait()
	assert.NoError(token.Error())

	assert.NoError(ts.integration.PublishEndnodeEvent(ts.basestationsEui, "ul", &pb))
	uplinkReceived := <-uplinkFrameChan
	assert.True(proto.Equal(&pb, uplinkReceived))
}

func (ts *TestIntegrationSuite) TestPublishBasestationEvent() {
	assert := require.New(ts.T())

	pb := msg.ProtoBasestationMessage{
		BsEui: "test_bs",

		V1: &msg.ProtoBasestationMessage_VmStatus{
			VmStatus: &msg.BasestationVariableMacStatus{
				MacTypes: []int64{1, 2, 3},
			},
		},
	}

	uplinkFrameChan := make(chan *msg.ProtoBasestationMessage)
	token := ts.mqttClient.Subscribe("test/bssci/0807060504030201/event/ep/vm", 0, func(c paho.Client, ms paho.Message) {
		var pl msg.ProtoBasestationMessage
		assert.NoError(ts.integration.unmarshal(ms.Payload(), &pl))
		uplinkFrameChan <- &pl
	})
	token.Wait()
	assert.NoError(token.Error())

	assert.NoError(ts.integration.PublishBasestationEvent(ts.basestationsEui, "vm", &pb))
	uplinkReceived := <-uplinkFrameChan
	assert.True(proto.Equal(&pb, uplinkReceived))
}

func (ts *TestIntegrationSuite) TestHandleServerResponse() {
	assert := require.New(ts.T())
	rspChan := make(chan *rsp.ProtoResponse, 1)
	ts.integration.SetServerResponseHandler(func(pl *rsp.ProtoResponse) {
		rspChan <- pl
	})

	response := rsp.ProtoResponse{
		OpId:  1,
		BsEui: "test_bs",
		V1: &rsp.ProtoResponse_Err{
			Err: &rsp.ErrorResponse{
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

func (ts *TestIntegrationSuite) TestHandleServerCommand() {
	assert := require.New(ts.T())
	cmdChan := make(chan *cmd.ProtoCommand, 1)
	ts.integration.SetServerCommandHandler(func(pl *cmd.ProtoCommand) {
		cmdChan <- pl
	})

	command := cmd.ProtoCommand{
		BsEui: "test_bs",
		V1: &cmd.ProtoCommand_VmStatus{
			VmStatus: &cmd.RequestVariableMacStatus{},
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
