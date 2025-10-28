package messages

import (

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

//go:generate msgp
//msgp:shim common.EUI64 as:uint64 using:common.Eui64toUnsignedInt/common.Eui64FromUnsignedInt

// Propagate Acknowledge
//
// # Wrapper for attPrpRsp, detPrpRsp or error messages in response to attPrp or detPrp
//
// BSSCI Bridge -> Service Center
type PrpAck struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation (just for compatibility, same as in attPrp/detPrp)
	OpId int64 `msg:"opId" json:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui" json:"epEui"`
	// True if propagation was successful
	Success bool `msg:"dualChan" json:"dualChan"`
	// True if attPrp, false if detPrp
	Attach bool `msg:"repetition" json:"repetition"`
}

func NewPrpAck(opId int64, epEui common.EUI64, success bool, attach bool,
) PrpAck {
	return PrpAck{
		Command:   structs.MsgPrpAck,
		OpId:      opId,
		EpEui:     epEui,
		Success:   success,
		Attach:    attach,
	}
}

func (m *PrpAck) GetOpId() int64 {
	return m.OpId
}

func (m *PrpAck) GetCommand() structs.Command {
	return structs.MsgPrpAck
}

// implements BasestationMessage.GetEventType()
func (m *PrpAck) GetEventType() events.EventType {
	return events.EventTypeBsPrpAck
}

// implements BasestationMessage.IntoProto()
func (m *PrpAck) IntoProto(bsEui *common.EUI64) *bs.BasestationUplink {

	var message bs.BasestationUplink

	if m != nil && bsEui != nil {
		bsEuiB := bsEui.String()
		epEuiB := m.EpEui.String()

		now := getNow().UnixNano()
		ts := TimestampNsToProto(now)

		var state bs.BasestationPropagationAck_State 
		if m.Attach {
			state = bs.BasestationPropagationAck_ATTACH
		} else {
			state = bs.BasestationPropagationAck_DETACH
		}

		message = bs.BasestationUplink{
			Ts:    ts,
			BsEui: bsEuiB,
			OpId:  m.OpId,
			Message: &bs.BasestationUplink_PrpAck{
				PrpAck: &bs.BasestationPropagationAck{
					EpEui:     epEuiB,
					Success:   m.Success,
					State:     state,
				},
			},
		}
	}

	return &message
}
