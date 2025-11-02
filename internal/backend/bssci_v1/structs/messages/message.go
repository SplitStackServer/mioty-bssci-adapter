package messages

import (
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v5/bs"
	"github.com/tinylib/msgp/msgp"
)

// Each message must implement this
type Message interface {
	// get the opId
	GetOpId() int64
	// get the name of this message type
	GetCommand() structs.Command
}

// Each message must implement this
type MessageMsgp interface {
	Message
	// message pack interfaces
	msgp.Encodable
	msgp.Marshaler
	msgp.Unmarshaler
	msgp.Decodable
}

type EndnodeMessage interface {
	Message
	GetEventType() events.EventType
	IntoProto(bsEui *common.EUI64) *bs.EndnodeUplink
}

type BasestationMessage interface {
	Message
	GetEventType() events.EventType
	IntoProto(bsEui *common.EUI64) *bs.BasestationUplink
}

type ServerMessage interface {
	MessageMsgp
	SetOpId(opId int64)
}
