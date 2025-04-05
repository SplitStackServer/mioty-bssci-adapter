package messages

import (
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/backend/events"
	"mioty-bssci-adapter/internal/common"

	"github.com/tinylib/msgp/msgp"
)

// Each message must implement this
type Message interface {
	// get the opId
	GetOpId() int64
	// get the name of this message type
	GetCommand() structs.Command
	// message pack interfaces
	msgp.Encodable
	msgp.Marshaler
	msgp.Unmarshaler
	msgp.Decodable
}

type EndnodeMessage interface {
	Message
	GetEventType() events.EventType
	IntoProto(bsEui common.EUI64) *msg.ProtoEndnodeMessage
}

type BasestationMessage interface {
	Message
	GetEventType() events.EventType
	IntoProto(bsEui *common.EUI64) *msg.ProtoBasestationMessage
}

type ServerMessage interface {
	Message
	SetOpId(opId int64)
}
