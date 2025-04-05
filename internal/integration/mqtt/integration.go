package mqtt

import (
	"bytes"
	"context"
	"sync"
	"text/template"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"mioty-bssci-adapter/internal/integration/auth"

	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/api/rsp"
	"mioty-bssci-adapter/internal/common"
	"mioty-bssci-adapter/internal/config"
)

const (
	stateTopicTemplate    = "basestation/{{ .BsEui }}/state/"
	eventTopicTemplate    = "basestation/{{ .BsEui }}/event/{{ .EventSource }}/{{ .EventType }}"
	commandTopicTemplate  = "basestation/{{ .BsEui }}/command/#"
	responseTopicTemplate = "basestation/{{ .BsEui }}/response/#"
)

const (
	eventSourceEndpoint    = "ep"
	eventSourceBasestation = "bs"
)

// Integration implements a MQTT Integration.
type Integration struct {
	auth auth.Authentication

	conn       paho.Client
	connMux    sync.RWMutex
	connClosed bool
	clientOpts *paho.ClientOptions

	serverCommandHandler  func(*cmd.ProtoCommand)
	serverResponseHandler func(*rsp.ProtoResponse)

	basestationsMux           sync.RWMutex
	basestations              map[common.EUI64]struct{}
	basestationsSubscribedMux sync.Mutex
	basestationsSubscribed    map[common.EUI64]struct{}
	terminateOnConnectError   bool
	stateRetained             bool
	maxTokenWait              time.Duration

	qos uint8

	eventTopicTemplate    *template.Template
	stateTopicTemplate    *template.Template
	commandTopicTemplate  *template.Template
	responseTopicTemplate *template.Template

	marshal   func(msg proto.Message) ([]byte, error)
	unmarshal func(b []byte, msg proto.Message) error
}

// NewIntegration creates a new Integration.
func NewIntegration(conf config.Config) (*Integration, error) {
	var err error

	integ := Integration{
		qos:                     conf.Integration.MQTT_V3.Auth.Generic.QOS,
		terminateOnConnectError: conf.Integration.MQTT_V3.TerminateOnConnectError,
		clientOpts:              paho.NewClientOptions(),
		basestations:            make(map[common.EUI64]struct{}),
		basestationsSubscribed:  make(map[common.EUI64]struct{}),
		stateRetained:           conf.Integration.MQTT_V3.StateRetained,
		maxTokenWait:            conf.Integration.MQTT_V3.MaxTokenWait,
	}

	// set authentication
	switch conf.Integration.MQTT_V3.Auth.Type {
	case "generic":
		integ.auth, err = auth.NewGenericAuthentication(conf)
		if err != nil {
			return nil, errors.Wrap(err, "new generic authentication error")
		}
	default:
		return nil, errors.Errorf("unknown auth type: %s", conf.Integration.MQTT_V3.Auth.Type)
	}

	// set marshaler
	switch conf.Integration.Marshaler {
	case "json":
		integ.marshal = func(msg proto.Message) ([]byte, error) {
			return protojson.MarshalOptions{
				EmitUnpopulated: false,
			}.Marshal(msg)
		}
		integ.unmarshal = func(b []byte, msg proto.Message) error {
			return protojson.UnmarshalOptions{
				DiscardUnknown: true,
				AllowPartial:   true,
			}.Unmarshal(b, msg)
		}
	case "protobuf":
		integ.marshal = func(msg proto.Message) ([]byte, error) {
			return proto.Marshal(msg)
		}

		integ.unmarshal = func(b []byte, msg proto.Message) error {
			return proto.Unmarshal(b, msg)
		}
	default:
		return nil, errors.Errorf("unknown marshaler: %s", conf.Integration.Marshaler)
	}

	// set topic templates
	integ.eventTopicTemplate, err = template.New("state").Parse(eventTopicTemplate)
	if err != nil {
		return nil, errors.Wrap(err, "parse event topic template error")
	}
	integ.stateTopicTemplate, err = template.New("state").Parse(stateTopicTemplate)
	if err != nil {
		return nil, errors.Wrap(err, "parse state topic template error")
	}
	integ.commandTopicTemplate, err = template.New("command").Parse(commandTopicTemplate)
	if err != nil {
		return nil, errors.Wrap(err, "parse command topic template error")
	}
	integ.responseTopicTemplate, err = template.New("state").Parse(responseTopicTemplate)
	if err != nil {
		return nil, errors.Wrap(err, "parse response topic template error")
	}

	// set mqtt parameters
	integ.clientOpts.SetProtocolVersion(4)
	integ.clientOpts.SetAutoReconnect(true) // this is required for buffering messages in case offline!
	integ.clientOpts.SetOnConnectHandler(integ.onConnected)
	integ.clientOpts.SetConnectionLostHandler(integ.onConnectionLost)
	integ.clientOpts.SetKeepAlive(conf.Integration.MQTT_V3.KeepAlive)
	integ.clientOpts.SetMaxReconnectInterval(conf.Integration.MQTT_V3.MaxReconnectInterval)

	if err = integ.auth.Init(integ.clientOpts); err != nil {
		return nil, errors.Wrap(err, "mqtt: init authentication error")
	}

	if bsEui := integ.auth.GetBasestationEui(); bsEui != nil {
		logger := log.With().Str("bs_eui", bsEui.String()).Logger()

		logger.Info().Msg("basestation EUI provided by authentication method")

		// Add basestation EUI to list of gateways we must subscribe to.
		integ.basestations[*bsEui] = struct{}{}

		// set last will and testament.
		pl := msg.ProtoBasestationState{
			BsEui: bsEui.ToUnsignedInt(),
			State: msg.ConnectionState_OFFLINE,
		}
		bb, err := integ.marshal(&pl)
		if err != nil {
			return nil, errors.Wrap(err, "marshal error")
		}

		topic := bytes.NewBuffer(nil)
		if err := integ.stateTopicTemplate.Execute(topic, struct {
			BsEui common.EUI64
		}{*bsEui}); err != nil {
			return nil, errors.Wrap(err, "execute state template error")
		}
		topicStr := topic.String()

		logger.Info().Str("topic", topicStr).Msg("setting last will and testament")

		integ.clientOpts.SetBinaryWill(topicStr, bb, integ.qos, true)
	}

	return &integ, nil
}

// Start the integration.
func (integ *Integration) Start() error {
	integ.connectLoop()
	go integ.reconnectLoop()
	go integ.subscribeLoop()
	return nil
}

// Stop the integration.
func (integ *Integration) Stop() error {
	integ.connMux.Lock()
	defer integ.connMux.Unlock()

	integ.basestationsMux.Lock()
	defer integ.basestationsMux.Unlock()

	// setup ctx logger
	logger := log.With().Logger()
	ctx := context.Background()
	ctx = logger.WithContext(ctx)

	// Set gateway state to offline for all gateways.
	for bsEui := range integ.basestations {
		logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("bs_eui", bsEui.String())
		})

		pl := msg.ProtoBasestationState{
			BsEui: bsEui.ToUnsignedInt(),
			State: msg.ConnectionState_OFFLINE,
		}
		if err := integ.PublishState(ctx, bsEui, &pl); err != nil {
			logger.Error().Err(err).Msg("publish state error")
		}
	}

	integ.conn.Disconnect(250)
	integ.connClosed = true
	return nil
}

// Updates the subscription for the given EUI.
func (integ *Integration) SetBasestationSubscription(subscribe bool, bsEui common.EUI64) error {

	logger := log.With().Str("bs_eui", bsEui.String()).Bool("subscribe", subscribe).Logger()

	if id := integ.auth.GetBasestationEui(); id != nil && *id == bsEui {
		logger.Debug().Msg("ignoring SetBasestationSubscription as EUI is set by authentication")
		return nil
	}
	logger.Debug().Msg("updating basestation subscription")

	integ.basestationsMux.Lock()
	defer integ.basestationsMux.Unlock()

	if subscribe {
		integ.basestations[bsEui] = struct{}{}
	} else {
		delete(integ.basestations, bsEui)
	}

	logger.Info().Msg("basestation subscription updated")

	return nil
}

// Set handler for server command messages
func (integ *Integration) SetServerCommandHandler(f func(*cmd.ProtoCommand)) {
	integ.serverCommandHandler = f
}

// Set handler for server command messages
func (integ *Integration) SetServerResponseHandler(f func(*rsp.ProtoResponse)) {
	integ.serverResponseHandler = f
}

// Publish basestation messages.
func (integ *Integration) PublishState(ctx context.Context, bsEui common.EUI64, pb *msg.ProtoBasestationState) error {
	logger := zerolog.Ctx(ctx)

	if integ.stateTopicTemplate == nil {
		logger.Debug().Msg("ignoring publish state, no stateTopicTemplate configured")
		return nil
	}

	mqttStateCounter().Inc()

	topic := bytes.NewBuffer(nil)
	if err := integ.stateTopicTemplate.Execute(topic, struct {
		BsEui common.EUI64
	}{bsEui}); err != nil {
		return errors.Wrap(err, "execute state template error")
	}
	topicStr := topic.String()

	bytes, err := integ.marshal(pb)
	if err != nil {
		return errors.Wrap(err, "marshal message error")
	}
	logger.Info().Str("topic", topicStr).Uint8("qos", integ.qos).Msg("publishing state")

	if err := tokenWrapper(integ.conn.Publish(topicStr, integ.qos, integ.stateRetained, bytes), integ.maxTokenWait); err != nil {
		return errors.Wrap(err, "publish state error")
	}
	logger.Debug().Str("topic", topicStr).Uint8("qos", integ.qos).Any("data", pb).Msg("published state")

	return nil
}

// Publish endnode messages.
func (integ *Integration) PublishEndnodeEvent(bsEui common.EUI64, event string, pb *msg.ProtoEndnodeMessage) error {
	// setup ctx logger
	logger := log.With().Str("bs_eui", bsEui.String()).Str("event", event).Str("source", eventSourceEndpoint).Logger()
	ctx := context.Background()
	ctx = logger.WithContext(ctx)

	return integ.publishEvent(ctx, bsEui, eventSourceEndpoint, event, pb)
}

// Publish basestation messages.
func (integ *Integration) PublishBasestationEvent(bsEui common.EUI64, event string, pb *msg.ProtoBasestationMessage) error {
	// setup ctx logger
	logger := log.With().Str("bs_eui", bsEui.String()).Str("event", event).Str("source", eventSourceBasestation).Logger()
	ctx := context.Background()
	ctx = logger.WithContext(ctx)

	return integ.publishEvent(ctx, bsEui, eventSourceBasestation, event, pb)
}

func (integ *Integration) publishEvent(ctx context.Context, bsEui common.EUI64, source string, event string, pb proto.Message) error {
	logger := zerolog.Ctx(ctx)

	mqttEventCounter(bsEui.String(), source, event)

	topic := bytes.NewBuffer(nil)
	if err := integ.eventTopicTemplate.Execute(topic, struct {
		BsEui       common.EUI64
		EventSource string
		EventType   string
	}{bsEui, source, event}); err != nil {
		return errors.Wrap(err, "execute event template error")
	}
	topicStr := topic.String()

	bytes, err := integ.marshal(pb)
	if err != nil {
		return errors.Wrap(err, "marshal message error")
	}

	logger.Info().Str("topic", topicStr).Uint8("qos", integ.qos).Msg("publishing event")

	if err := tokenWrapper(integ.conn.Publish(topicStr, integ.qos, false, bytes), integ.maxTokenWait); err != nil {
		return err
	}
	logger.Debug().Str("topic", topicStr).Uint8("qos", integ.qos).Any("data", pb).Msg("published event")
	return nil
}

func (integ *Integration) connect() error {
	integ.connMux.Lock()
	defer integ.connMux.Unlock()

	if err := integ.auth.Update(integ.clientOpts); err != nil {
		return errors.Wrap(err, "update authentication error")
	}

	if integ.conn != nil {
		integ.conn.Disconnect(250)
	}
	integ.conn = paho.NewClient(integ.clientOpts)
	if err := tokenWrapper(integ.conn.Connect(), integ.maxTokenWait); err != nil {
		return err
	}

	return nil
}

func (integ *Integration) disconnect() error {
	mqttDisconnectCounter().Inc()

	integ.connMux.Lock()
	defer integ.connMux.Unlock()

	integ.conn.Disconnect(250)
	return nil
}

func (integ *Integration) onConnected(c paho.Client) {
	mqttConnectCounter().Inc()
	log.Info().Msg("connected to mqtt broker")

	integ.basestationsSubscribedMux.Lock()
	defer integ.basestationsSubscribedMux.Unlock()

	integ.basestationsSubscribed = make(map[common.EUI64]struct{})
}

func (integ *Integration) onConnectionLost(c paho.Client, err error) {
	if integ.terminateOnConnectError {
		log.Fatal().Err(err).Msg("mqtt connection lost")
	}
	mqttDisconnectCounter().Inc()
	log.Error().Err(err).Msg("mqtt connection lost")
}

// connectLoop blocks until the client is connected
func (integ *Integration) connectLoop() {
	for {
		if err := integ.connect(); err != nil {
			if integ.terminateOnConnectError {
				log.Fatal().Err(err).Msg("mqtt connection error")
			}

			log.Error().Err(err).Msg("mqtt connection error")
			time.Sleep(time.Second * 2)

		} else {
			break
		}
	}
}

func (integ *Integration) reconnectLoop() {
	if integ.auth.ReconnectAfter() > 0 {
		for {
			if integ.isClosed() {
				break
			}

			time.Sleep(integ.auth.ReconnectAfter())
			log.Info().Msg("re-connect triggered")

			mqttReconnectCounter().Inc()

			integ.disconnect()
			integ.connectLoop()
		}
	}
}

func (integ *Integration) subscribeLoop() {
	for {
		time.Sleep(time.Millisecond * 100)

		if integ.isClosed() {
			break
		}

		if !integ.conn.IsConnected() {
			continue
		}

		var subscribe []common.EUI64
		var unsubscribe []common.EUI64

		integ.basestationsMux.RLock()
		integ.basestationsSubscribedMux.Lock()

		// subscribe
		for bsEui := range integ.basestations {
			if _, ok := integ.basestationsSubscribed[bsEui]; !ok {
				subscribe = append(subscribe, bsEui)
			}
		}

		// unsubscribe
		for bsEui := range integ.basestationsSubscribed {
			if _, ok := integ.basestations[bsEui]; !ok {
				unsubscribe = append(unsubscribe, bsEui)
			}
		}

		// unlock gatewaysMux so that SetGatewaySubscription can write again
		// to the map, in which case changes are picked up in the next run
		integ.basestationsMux.RUnlock()

		// setup ctx logger
		logger := log.With().Logger()
		ctx := context.Background()
		ctx = logger.WithContext(ctx)

		// subscribe to all basestations
		for _, bsEui := range subscribe {
			logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("bs_eui", bsEui.String())
			})

			pl := msg.ProtoBasestationState{
				BsEui: bsEui.ToUnsignedInt(),
				State: msg.ConnectionState_ONLINE,
			}

			if err := integ.subscribeBasestation(ctx, bsEui); err != nil {
				logger.Error().Err(err).Msg("mqtt subscribe basestation error")
			} else {
				if err := integ.PublishState(ctx, bsEui, &pl); err != nil {
					logger.Error().Err(err).Msg("publish basestation error")
				} else {
					integ.basestationsSubscribed[bsEui] = struct{}{}
				}
			}
		}

		// unsubscribe from all remaining basestations
		for _, bsEui := range unsubscribe {
			logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("bs_eui", bsEui.String())
			})

			pl := msg.ProtoBasestationState{
				BsEui: bsEui.ToUnsignedInt(),
				State: msg.ConnectionState_OFFLINE,
			}

			if err := integ.unsubscribeBasestation(ctx, bsEui); err != nil {
				logger.Error().Err(err).Msg("mqtt unsubscribe basestation error")
			} else {
				if err := integ.PublishState(ctx, bsEui, &pl); err != nil {
					logger.Error().Err(err).Msg("publish basestation error")
				} else {
					delete(integ.basestationsSubscribed, bsEui)
				}
			}
		}

		integ.basestationsSubscribedMux.Unlock()
	}
}

func (integ *Integration) subscribeBasestation(ctx context.Context, bsEui common.EUI64) error {
	logger := zerolog.Ctx(ctx)

	// subscribe to server commands
	topic := bytes.NewBuffer(nil)
	if err := integ.commandTopicTemplate.Execute(topic, struct{ BsEui common.EUI64 }{bsEui}); err != nil {
		return errors.Wrap(err, "execute command topic template error")
	}
	topicStr := topic.String()
	logger.Info().Str("topic", topicStr).Uint8("qos", integ.qos).Msg("subscribing to topic")

	if err := tokenWrapper(integ.conn.Subscribe(topicStr, integ.qos, integ.handleServerCommand), integ.maxTokenWait); err != nil {
		return errors.Wrap(err, "subscribe command topic error")
	}
	logger.Debug().Str("topic", topicStr).Uint8("qos", integ.qos).Msg("subscribed to topic")

	// subscribe to server responses
	topic = bytes.NewBuffer(nil)
	if err := integ.responseTopicTemplate.Execute(topic, struct{ BsEui common.EUI64 }{bsEui}); err != nil {
		return errors.Wrap(err, "execute response topic template error")
	}
	topicStr = topic.String()
	logger.Info().Str("topic", topicStr).Uint8("qos", integ.qos).Msg("subscribing to topic")

	if err := tokenWrapper(integ.conn.Subscribe(topicStr, integ.qos, integ.handleServerResponse), integ.maxTokenWait); err != nil {
		return errors.Wrap(err, "subscribe response topic error")
	}
	logger.Debug().Str("topic", topicStr).Uint8("qos", integ.qos).Msg("subscribed to topic")

	return nil
}

func (integ *Integration) handleServerCommand(c paho.Client, msg paho.Message) {
	var pb cmd.ProtoCommand

	if err := integ.unmarshal(msg.Payload(), &pb); err != nil {
		log.Error().Err(err).Msg("unmarshal server command error")
		return
	}
	mqttCommandCounter(common.Eui64FromUnsignedInt(pb.BsEui).String())
	integ.serverCommandHandler(&pb)

}

func (integ *Integration) handleServerResponse(c paho.Client, msg paho.Message) {
	var pb rsp.ProtoResponse

	if err := integ.unmarshal(msg.Payload(), &pb); err != nil {
		log.Error().Err(err).Msg("unmarshal server response error")
		return
	}
	mqttResponseCounter(common.Eui64FromUnsignedInt(pb.BsEui).String())
	integ.serverResponseHandler(&pb)

}

func (b *Integration) unsubscribeBasestation(ctx context.Context, bsEui common.EUI64) error {
	logger := zerolog.Ctx(ctx)

	topic := bytes.NewBuffer(nil)
	if err := b.commandTopicTemplate.Execute(topic, struct{ BsEui common.EUI64 }{bsEui}); err != nil {
		return errors.Wrap(err, "execute command topic template error")
	}
	topicStr := topic.String()

	logger.Info().Str("topic", topicStr).Msg("unsubscribing from topic")

	if err := tokenWrapper(b.conn.Unsubscribe(topicStr), b.maxTokenWait); err != nil {
		return errors.Wrap(err, "unsubscribe topic error")
	}

	logger.Debug().Str("topic", topicStr).Msg("unsubscribed from topic")

	return nil
}

// isClosed returns true when the integration is shutting down.
func (integ *Integration) isClosed() bool {
	integ.connMux.RLock()
	defer integ.connMux.RUnlock()
	return integ.connClosed
}

func tokenWrapper(token paho.Token, timeout time.Duration) error {
	if !token.WaitTimeout(timeout) {
		return errors.New("token wait timeout error")
	}
	return token.Error()
}
