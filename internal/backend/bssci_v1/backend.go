package bssci_v1

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	"net"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/api/rsp"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs/messages"
	"mioty-bssci-adapter/internal/backend/events"
	"mioty-bssci-adapter/internal/common"
	"mioty-bssci-adapter/internal/config"
)

type Backend struct {
	sync.RWMutex

	caCert  string
	tlsCert string
	tlsKey  string

	listener net.Listener
	isClosed bool

	basestations basestations

	statsInterval   time.Duration
	pingInterval    time.Duration
	keepAlivePeriod time.Duration
	writeTimeout    time.Duration

	basestationMessageHandler func(common.EUI64, events.EventType, *msg.ProtoBasestationMessage)
	endnodeMessageHandler     func(common.EUI64, events.EventType, *msg.ProtoEndnodeMessage)
}

// NewBackend creates a new Backend.
func NewBackend(conf config.Config) (backend *Backend, err error) {
	b := Backend{
		basestations: basestations{
			basestations: make(map[common.EUI64]*connection),
		},

		caCert:  conf.Backend.BssciV1.CACert,
		tlsCert: conf.Backend.BssciV1.TLSCert,
		tlsKey:  conf.Backend.BssciV1.TLSKey,

		statsInterval:   conf.Backend.BssciV1.StatsInterval,
		pingInterval:    conf.Backend.BssciV1.PingInterval,
		keepAlivePeriod: conf.Backend.BssciV1.KeepAlivePeriod,
		writeTimeout:    time.Second,
	}

	// create the listener
	b.listener, err = NewTcpKeepAliveListener(conf.Backend.BssciV1.Bind, b.keepAlivePeriod)
	if err != nil {
		return nil, errors.Wrap(err, "create tcp keep alive listener error")
	}

	// if the CA and TLS cert is configured, setup client certificate verification.
	if b.tlsCert != "" && b.tlsKey != "" && b.caCert != "" {
		rawCACert, err := os.ReadFile(b.caCert)
		if err != nil {
			return nil, errors.Wrap(err, "read ca cert error")
		}
		tlsCert, err := tls.LoadX509KeyPair(b.tlsCert, b.tlsKey)
		if err != nil {
			return nil, errors.Wrap(err, "read tls cert error")
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(rawCACert)

		// wrap the tcp listener in a tls listener
		b.listener = tls.NewListener(b.listener, &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			ClientCAs:    caCertPool,
			ClientAuth:   tls.RequireAndVerifyClientCert,
		})

	} else {
		log.Warn().Msg("config does not provide a TLS certificate, generating one")
		tlsCert, err := common.GenX509KeyPair()
		if err != nil {
			return nil, errors.Wrap(err, "generate tls cert error")
		}

		b.listener = tls.NewListener(b.listener, &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
		})
	}

	backend = &b
	return
}

// Handler for Subscribe events.
func (b *Backend) SetSubscribeEventHandler(f func(events.Subscribe)) {
	b.basestations.subscribeEventHandler = f
}

// Handler for connection messages from basestations
func (b *Backend) SetBasestationMessageHandler(f func(common.EUI64, events.EventType, *msg.ProtoBasestationMessage)) {
	b.basestationMessageHandler = f
}

// Handler for uplink messages from endnodes
func (b *Backend) SetEndnodeMessageHandler(f func(common.EUI64, events.EventType, *msg.ProtoEndnodeMessage)) {
	b.endnodeMessageHandler = f
}

// Handler for server commands
func (b *Backend) HandleServerCommand(pb *cmd.ProtoCommand) error {
	if pb == nil {
		return errors.New("empty protobuf command")
	}

	bsEui, err := common.Eui64FromHexString(pb.BsEui)
	if err != nil {
		return errors.New("invalid eui64 hex string")
	}

	logger := log.With().Str("bs_eui", bsEui.String()).Logger()

	var msg messages.ServerMessage

	switch pb.V1.(type) {
	case *cmd.ProtoCommand_DlDataQue:
		command := pb.GetDlDataQue()
		msgA, err := messages.NewDlDataQueFromProto(0, command)
		if err != nil {
			return err
		}
		msg = msgA

		logger.Debug().Str("proto", "ProtoCommand_DlDataQue").Str("ep_eui", msgA.EpEui.String()).Uint64("que_id", msgA.QueId).Msg("queing downlink")

	case *cmd.ProtoCommand_DlDataRev:
		command := pb.GetDlDataRev()
		msgA, err := messages.NewDlDataRevFromProto(0, command)
		if err != nil {
			return err
		}
		msg = msgA

		logger.Debug().Str("proto", "ProtoCommand_DlDataRev").Str("ep_eui", msgA.EpEui.String()).Uint64("que_id", msgA.QueId).Msg("revoking downlink")

	case *cmd.ProtoCommand_DlRxStatQry:
		command := pb.GetDlRxStatQry()
		msgA, err := messages.NewDlRxStatQryFromProto(0, command)
		if err != nil {
			return err
		}
		msg = msgA

		logger.Debug().Str("proto", "ProtoCommand_DlRxStatQry").Str("ep_eui", msgA.EpEui.String()).Msg("requesting downlink status ")

	case *cmd.ProtoCommand_AttPrp:
		command := pb.GetAttPrp()
		msgA, err := messages.NewAttPrpFromProto(0, command)
		if err != nil {
			return err
		}
		msg = msgA

		logger.Debug().Str("proto", "ProtoCommand_AttPrp").Str("ep_eui", msgA.EpEui.String()).Msg("propagate attaching endnode")

	case *cmd.ProtoCommand_DetPrp:
		command := pb.GetDetPrp()
		msgA, err := messages.NewDetPrpFromProto(0, command)
		if err != nil {
			return err
		}
		msg = msgA

		logger.Debug().Str("proto", "ProtoCommand_DetPrp").Str("ep_eui", msgA.EpEui.String()).Msg("propagate detaching endnode")

	case *cmd.ProtoCommand_ReqStatus:
		command := pb.GetReqStatus()
		msgA, err := messages.NewStatusFromProto(0, command)
		if err != nil {
			return err
		}
		msg = msgA

		logger.Debug().Str("proto", "ProtoCommand_ReqStatus").Msg("requesting basestation status")

	case *cmd.ProtoCommand_VmActivate:
		command := pb.GetVmActivate()
		msgA, err := messages.NewVmActivateFromProto(0, command)
		if err != nil {
			return err
		}
		msg = msgA

		logger.Debug().Str("proto", "ProtoCommand_VmActivate").Msgf("requesting variable mac activation: %v", msgA.MacType)

	case *cmd.ProtoCommand_VmDeactivate:
		command := pb.GetVmDeactivate()
		msgA, err := messages.NewVmDeactivateFromProto(0, command)
		if err != nil {
			return err
		}
		msg = msgA

		logger.Debug().Str("proto", "ProtoCommand_VmDeactivate").Msgf("requesting variable mac deactivation: %v", msgA.MacType)

	case *cmd.ProtoCommand_VmStatus:
		command := pb.GetVmStatus()
		msgA, err := messages.NewVmStatusFromProto(0, command)
		if err != nil {
			return err
		}
		msg = msgA

		logger.Debug().Str("proto", "ProtoCommand_VmStatus").Msg("requesting variable mac status")

	default:
		return errors.New("empty protobuf command")
	}

	return b.sendServerMessageToBasestation(bsEui, msg)
}

// Handler for server response messages
func (b *Backend) HandleServerResponse(pb *rsp.ProtoResponse) error {
	if pb == nil {
		return errors.New("empty protobuf command")
	}

	opId := pb.OpId
	bsEui, err := common.Eui64FromHexString(pb.BsEui)
	if err != nil {
		return err
	}

	var msg messages.Message

	switch pb.V1.(type) {
	case *rsp.ProtoResponse_DetRsp:
		command := pb.GetDetRsp()
		msgA, err := messages.NewDetRspFromProto(opId, command)
		if err != nil {
			return err
		}
		msg = msgA
		log.Debug().Str("proto", "ProtoResponse_DetRsp").Int64("op_id", opId).Msgf("detaching endnode %v from basestation %v", command.EndnodeEui, bsEui.String())
	case *rsp.ProtoResponse_AttRsp:
		command := pb.GetAttRsp()
		msgA, err := messages.NewAttRspFromProto(opId, command)
		if err != nil {
			return err
		}
		msg = msgA
		log.Debug().Str("proto", "ProtoResponse_AttRsp").Int64("op_id", opId).Msgf("attaching endnode %v to basestation %v", command.EndnodeEui, bsEui.String())

	case *rsp.ProtoResponse_Err:
		command := pb.GetErr()
		msgA := messages.NewBssciError(opId, 5, command.GetMessage())

		msg = &msgA
		log.Warn().Str("proto", "ProtoResponse_Err").Int64("op_id", opId).Msgf("server responded with error: %s", command.GetMessage())

	default:
		return errors.New("empty protobuf command")
	}

	return b.sendServerResponseToBasestation(bsEui, msg)
}

// Stops the backend.
func (b *Backend) Stop() error {
	b.isClosed = true
	return b.listener.Close()
}

// Starts the backend.
func (b *Backend) Start() error {

	log.Info().Str("addr", b.listener.Addr().String()).Msg("STARTING SERVICE")

	go func() {
		for !b.isClosed {
			// accept a new connection
			conn, err := b.listener.Accept()

			if err != nil {
				log.Error().Err(err).Msg("connection accept failed")
			}
			// defer conn.Close()

			logger := log.With().Str("remote", conn.RemoteAddr().String()).Logger()
			logger.Info().Msg("accepted new connection")

			// try to read Con message
			conn.SetReadDeadline(time.Now().Add(time.Minute))
			cmdHeader, raw, err := ReadBssciMessage(conn)

			if err != nil {
				logger.Error().Err(err).Msg("codec error")
			} else {
				// first message after connecting should always be Con
				cmd := cmdHeader.GetCommand()

				if cmd == structs.MsgCon {
					var con messages.Con
					_, err = con.UnmarshalMsg(raw)
					if err != nil {
						logger.Error().Err(err).Str("command", string(cmd)).Msg("unmarshal msgp error")
					} else {
						logger.Info().Str("command", string(cmd)).Msg("initializing basestation connection")
						ctx := context.Background()
						ctx = logger.WithContext(ctx)
						// handle the basestation in a new goroutine
						go b.initBasestation(ctx, con, conn, b.handleBasestationMessages)
					}
				} else {
					logger.Error().Str("command", string(cmd)).Msg("expected con command")

				}
			}
		}
	}()
	return nil
}

func (b *Backend) initBasestation(ctx context.Context, con messages.Con, conn net.Conn, handler func(ctx context.Context, eui common.EUI64, conn *connection) error) error {
	defer conn.Close()

	eui := con.GetEui()

	logger := zerolog.Ctx(ctx)
	logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("bs_eui", eui.String())
	})

	// keep track of new connections
	connectCounter(eui.String()).Inc()
	b.forwardBasestationMessage(ctx, eui, &con)

	bsConnection := newConnection(conn, con.SnBsUuid)
	conRsp := messages.NewConRsp(con.OpId, con.Version, bsConnection.SnScUuid)

	// set the gateway connection
	if err := b.basestations.set(eui, &bsConnection); err != nil {
		logger.Error().Err(err).Msg("failed to set connection")
	}

	logger.Info().Msg("basestation connected")

	// setup recurring tasks
	done := make(chan struct{})

	// remove the basestation on return
	defer func() {
		done <- struct{}{}
		b.basestations.remove(eui)
		bsConnection.conn.Close()
		disconnectCounter(eui.String()).Inc()
		logger.Info().Msg("basestation disconnected")
	}()

	// setup ping and status tickers
	pingTicker := time.NewTicker(b.pingInterval)
	defer pingTicker.Stop()
	statusTicker := time.NewTicker(b.statsInterval)
	defer statusTicker.Stop()

	go func() {
		logger.Debug().Msg("scheduling status messages")
		for {
			select {
			case <-pingTicker.C:
				opId := bsConnection.GetAndDecrementOpId()
				msg := messages.NewPing(opId)

				err := bsConnection.Write(&msg, b.writeTimeout)
				if err != nil {
					logger.Error().Err(err).Str("command", string(msg.GetCommand())).Msg("failed to send scheduled ping request")
					return
				}

				logger.Debug().Msg("sent scheduled ping request")
				pingPongCounter("server", eui.String())

			case <-statusTicker.C:
				opId := bsConnection.GetAndDecrementOpId()
				msg := messages.NewStatus(opId)

				err := bsConnection.Write(&msg, b.writeTimeout)
				if err != nil {
					logger.Error().Err(err).Str("command", string(msg.GetCommand())).Msg("failed to send scheduled status request")
					return
				}

				logger.Debug().Msg("sent scheduled status request")
				messageSendCounter(eui.String(), string(msg.GetCommand()))
			case <-done:
				logger.Debug().Msg("stopping scheduled status messages")
				return
			}
		}
	}()

	// send ConRsp
	err := bsConnection.Write(&conRsp, b.writeTimeout)
	if err != nil {
		logger.Error().Err(err).Str("command", string(conRsp.GetCommand())).Msg("failed to send message")
		// terminate this connection on error
		return err
	}
	messageSendCounter(eui.String(), string(conRsp.GetCommand()))

	// start the message handler
	err = handler(ctx, eui, &bsConnection)
	return err
}

// handle all messages coming from a client
func (b *Backend) handleBasestationMessages(ctx context.Context, eui common.EUI64, connection *connection) error {
	logger := zerolog.Ctx(ctx)
	for {

		cmdHeader, raw, err := connection.Read()

		if err != nil {
			logger.Error().Err(err).Msg("failed to read message")
			// terminate this connection
			return err
		}

		opId := cmdHeader.GetOpId()
		cmd := cmdHeader.GetCommand()

		// update logging context
		logger := logger.With().Str("command", string(cmd)).Int64("op_id", opId).Logger()

		logger.Debug().Hex("raw", raw).Msg("received message")

		messageReceiveCounter(eui.String(), string(cmd))

		var response messages.Message
		// only match ClientMsg... messages
		switch cmd {
		case structs.ClientMsgCon:
			// handle con message
			var msg messages.Con
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				response = log_and_notify_msgp_error(logger, err, opId)
				break
			}
			response = b.handleConMessage(ctx, connection, msg)
		case structs.ClientMsgAtt:
			// handle attach message
			var msg messages.Att
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				response = log_and_notify_msgp_error(logger, err, opId)
				break
			}
			response = b.handleAttMessage(ctx, eui, &msg)
		case structs.ClientMsgDet:
			// handle detach message
			var msg messages.Det
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				response = log_and_notify_msgp_error(logger, err, opId)
				break
			}
			response = b.handleDetMessage(ctx, eui, &msg)
		case structs.ClientMsgUlData:
			// handle uplink data message
			var msg messages.UlData
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				response = log_and_notify_msgp_error(logger, err, opId)
				break
			}
			response = b.handleUlDataMessage(ctx, eui, &msg)
		case structs.ClientMsgVmUlData:
			// handle variable mac uplink data message
			var msg messages.VmUlData
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				response = log_and_notify_msgp_error(logger, err, opId)
				break
			}
			response = b.handleVmUlDataMessage(ctx, eui, &msg)
		case structs.ClientMsgDlDataRes:
			// handle downlink data result response
			var msg messages.DlDataRes
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				response = log_and_notify_msgp_error(logger, err, opId)
				break
			}
			response = b.handleDlDataResMessage(ctx, eui, &msg)
		case structs.ClientMsgDlRxStat:
			// handle downlink rx status data message
			var msg messages.DlRxStat
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				response = log_and_notify_msgp_error(logger, err, opId)
				break
			}
			response = b.handleDlRxStatMessage(ctx, eui, &msg)
		case structs.ClientMsgStatusRsp:
			// handle status response message
			var msg messages.StatusRsp
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				response = log_and_notify_msgp_error(logger, err, opId)
				break
			}
			response = b.handleStatusRspMessage(ctx, eui, &msg)
		case structs.ClientMsgPing:
			// handle ping message
			pingPongCounter("client", eui.String())
			defaultResponse := messages.NewPingRsp(opId)
			response = &defaultResponse
		case structs.ClientMsgPingRsp:
			// handle ping response (pong) message
			defaultResponse := messages.NewPingCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgDlDataRevRsp:
			// handle downlink data revoke response message
			defaultResponse := messages.NewDlDataRevCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgDlDataQueRsp:
			// handle downlink data queue response message
			defaultResponse := messages.NewDlDataQueCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgDlRxStatQryRsp:
			// handle downlink rx status query response message
			defaultResponse := messages.NewDlRxStatQryCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgAttPrpRsp:
			// handle attach propagate response message
			defaultResponse := messages.NewAttPrpCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgDetPrpRsp:
			// handle detach propagate response message
			defaultResponse := messages.NewDetPrpCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgVmActivateRsp:
			// handle variable mac activate response message
			defaultResponse := messages.NewVmActivateCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgVmDeactivateRsp:
			// handle variable mac deactivate response message
			defaultResponse := messages.NewVmDeactivateCmp(opId)
			response = &defaultResponse
		case structs.ClientMsgVmStatusRsp:
			// handle variable mac status response message
			var msg messages.VmStatusRsp
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				response = log_and_notify_msgp_error(logger, err, opId)
				break
			}
			response = b.handleVmStatusRspMessage(ctx, eui, &msg)
		case structs.ClientMsgError:
			// handle error message
			var msg messages.BssciError
			_, err = msg.UnmarshalMsg(raw)
			if err != nil {
				response = log_and_notify_msgp_error(logger, err, opId)
				break
			}
			logger.Warn().Uint32("err_code", msg.Code).Str("err_msg", msg.Message).Msg("received bssci error message")
			defaultResponse := messages.NewBssciErrorAck(opId)
			response = &defaultResponse
		case structs.ClientMsgErrorAck:
			// Equivalent to ...Cmp message
			continue
		case structs.ClientMsgUlDataCmp:
			// ...Cmp messages need no further handling
			continue
		case structs.ClientMsgPingCmp:
			continue
		case structs.ClientMsgConCmp:
			continue
		case structs.ClientMsgDlDataResCmp:
			continue
		case structs.ClientMsgAttCmp:
			continue
		case structs.ClientMsgDetCmp:
			continue
		case structs.ClientMsgVmUlDataCmp:
			continue
		case structs.ClientMsgDlRxStatCmp:
			continue

		default:
			logger.Warn().Msg("unsupported message type")
			bssciError := messages.NewBssciError(opId, 5, "unsupported message type")
			response = &bssciError
		}

		if response != nil {
			err := connection.Write(response, b.writeTimeout)
			if err != nil {
				logger.Error().Err(err).Msg("failed to write message")
				// terminate this connection
				return err
			}
			messageSendCounter(eui.String(), string(response.GetCommand()))
			logger.Debug().Any("json", response).Msg("sent response")
		}
	}
}

// upstream messages from basestations
func (b *Backend) forwardBasestationMessage(ctx context.Context, eui common.EUI64, msg messages.BasestationMessage) messages.Message {
	logger := zerolog.Ctx(ctx)

	if b.basestationMessageHandler != nil {
		data := msg.IntoProto(&eui)
		b.basestationMessageHandler(eui, msg.GetEventType(), data)
		return nil
	}

	logger.Warn().Msg("basestationConnectionMessageHandler not set")
	response := messages.NewBssciError(msg.GetOpId(), 5, "server unable to handle message")
	return &response

}

// upstream messages from endnodes
func (b *Backend) forwardEndnodeMessage(ctx context.Context, eui common.EUI64, msg messages.EndnodeMessage) messages.Message {
	logger := zerolog.Ctx(ctx)

	if b.endnodeMessageHandler != nil {
		data := msg.IntoProto(eui)
		b.endnodeMessageHandler(eui, msg.GetEventType(), data)
		return nil
	}

	logger.Warn().Msg("endnodeMessageHandler not set")
	response := messages.NewBssciError(msg.GetOpId(), 5, "server unable to handle message")
	return &response
}

func (b *Backend) handleConMessage(ctx context.Context, conn *connection, msg messages.Con) messages.Message {
	logger := zerolog.Ctx(ctx)

	error_response := b.forwardBasestationMessage(ctx, msg.BsEui, &msg)
	if error_response == nil {
		resume, snScUuid := conn.ResumeConnection(msg.SnBsUuid.ToUuid(), msg.SnScOpId)

		conRsp := messages.NewConRsp(msg.GetOpId(), msg.Version, snScUuid)

		// check if session uuid is identical to current session
		if resume {
			logger.Info().Msg("Resuming current session")
			// set resume flag
			conRsp.SnResume = true
		} else {
			logger.Warn().Msg("Failed to resume current session")
		}
		return &conRsp
	}
	return error_response

}

func (b *Backend) handleStatusRspMessage(ctx context.Context, eui common.EUI64, msg *messages.StatusRsp) messages.Message {
	error_response := b.forwardBasestationMessage(ctx, eui, msg)
	if error_response == nil {
		response := messages.NewStatusCmp(msg.GetOpId())
		return &response
	}
	return error_response
}

func (b *Backend) handleVmStatusRspMessage(ctx context.Context, eui common.EUI64, msg *messages.VmStatusRsp) messages.Message {
	error_response := b.forwardBasestationMessage(ctx, eui, msg)
	if error_response == nil {
		response := messages.NewVmStatusCmp(msg.GetOpId())
		return &response
	}
	return error_response
}

func (b *Backend) handleAttMessage(ctx context.Context, eui common.EUI64, msg *messages.Att) messages.Message {
	// Att has to be handled by downstream application
	return b.forwardEndnodeMessage(ctx, eui, msg)
}

func (b *Backend) handleDetMessage(ctx context.Context, eui common.EUI64, msg *messages.Det) messages.Message {
	// Det has to be handled by downstream application
	return b.forwardEndnodeMessage(ctx, eui, msg)
}

func (b *Backend) handleUlDataMessage(ctx context.Context, eui common.EUI64, msg *messages.UlData) messages.Message {
	error_response := b.forwardEndnodeMessage(ctx, eui, msg)
	if error_response == nil {
		response := messages.NewUlDataRsp(msg.GetOpId())
		return &response
	}
	return error_response
}

func (b *Backend) handleVmUlDataMessage(ctx context.Context, eui common.EUI64, msg *messages.VmUlData) messages.Message {
	error_response := b.forwardEndnodeMessage(ctx, eui, msg)
	if error_response == nil {
		response := messages.NewVmUlDataRsp(msg.GetOpId())
		return &response
	}
	return error_response
}

func (b *Backend) handleDlRxStatMessage(ctx context.Context, eui common.EUI64, msg *messages.DlRxStat) messages.Message {
	error_response := b.forwardEndnodeMessage(ctx, eui, msg)
	if error_response == nil {
		response := messages.NewDlRxStatRsp(msg.GetOpId())
		return &response
	}
	return error_response
}

func (b *Backend) handleDlDataResMessage(ctx context.Context, eui common.EUI64, msg *messages.DlDataRes) messages.Message {
	error_response := b.forwardEndnodeMessage(ctx, eui, msg)
	if error_response == nil {
		response := messages.NewDlDataResRsp(msg.GetOpId())
		return &response
	}
	return error_response
}

// sends a server response to a basestation
func (b *Backend) sendServerResponseToBasestation(bsEui common.EUI64, msg messages.Message) error {
	if msg != nil {
		b.Lock()
		defer b.Unlock()
		logger := log.With().Str("bs_eui", bsEui.String()).Str("command", string(msg.GetCommand())).Int64("op_id", msg.GetOpId()).Logger()

		bsConnection, err := b.basestations.get(bsEui)
		if err != nil {
			logger.Error().Err(err).Msg("basestation does not exist")
			return err
		}

		err = bsConnection.Write(msg, b.writeTimeout)
		if err != nil {
			logger.Error().Err(err).Msg("failed to send to basestation")
			return err
		}

		messageSendCounter(bsEui.String(), string(msg.GetCommand()))
		logger.Debug().Int64("op_id", msg.GetOpId()).Any("json", msg).Msg("sent server response")

	}
	return nil
}

// sends a server message to a basestation
func (b *Backend) sendServerMessageToBasestation(bsEui common.EUI64, msg messages.ServerMessage) error {
	if msg != nil {
		b.Lock()
		defer b.Unlock()
		logger := log.With().Str("bs_eui", bsEui.String()).Str("command", string(msg.GetCommand())).Logger()

		bsConnection, err := b.basestations.get(bsEui)
		if err != nil {
			logger.Error().Err(err).Msg("basestation does not exist")
			return err
		}

		// get a new opId
		opId := bsConnection.GetAndDecrementOpId()
		msg.SetOpId(opId)

		err = bsConnection.Write(msg, b.writeTimeout)
		if err != nil {
			logger.Error().Err(err).Msg("failed to send to basestation")
			return err
		}
		messageSendCounter(bsEui.String(), string(msg.GetCommand()))

		logger.Debug().Int64("op_id", msg.GetOpId()).Any("json", msg).Msg("sent server message")

		return nil

	}
	return nil
}

func log_and_notify_msgp_error(logger zerolog.Logger, err error, opId int64) messages.Message {
	logger.Error().Err(err).Msg("unmarshal msgp error")
	response := messages.NewBssciError(opId, 5, "message pack error")
	return &response
}
