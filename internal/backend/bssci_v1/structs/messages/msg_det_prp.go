package messages

import (
	"errors"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

//go:generate msgp
//msgp:shim common.EUI64 as:uint64 using:common.Eui64toUnsignedInt/common.Eui64FromUnsignedInt

// Detach propagate
//
// Service Center -> Basestation
//
// The detach propagate operation is initiated by the Service Center to propagate an End
// Point detachment to the Base Station.
type DetPrp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui" json:"epEui"`
}

func NewDetPrp(
	opId int64,
	epEui common.EUI64,
) DetPrp {
	return DetPrp{
		Command: structs.MsgDetPrp,
		OpId:    opId,
		EpEui:   epEui,
	}
}

func NewDetPrpFromProto(opId int64, pb *bs.DetachPropagate) (*DetPrp, error) {
	if pb != nil {
		epEui, err := common.Eui64FromHexString(pb.EndnodeEui)
		if err != nil {
			return nil, err
		}
		m := NewDetPrp(opId, epEui)
		return &m, nil
	}
	return nil, errors.New("invalid DetachPropagate command")
}

func (m *DetPrp) GetOpId() int64 {
	return m.OpId
}

func (m *DetPrp) GetCommand() structs.Command {
	return structs.MsgDetPrp
}

// implements ServerMessage
func (m *DetPrp) SetOpId(opId int64) {
	m.OpId = opId
}

// Detach propagate response
//
// Basestation -> Service Center
type DetPrpRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDetPrpRsp(opId int64) DetPrpRsp {
	return DetPrpRsp{
		Command: structs.MsgDetPrpRsp,
		OpId:    opId,
	}
}

func (m *DetPrpRsp) GetOpId() int64 {
	return m.OpId
}

func (m *DetPrpRsp) GetCommand() structs.Command {
	return structs.MsgDetPrpRsp
}

// Detach propagate complete
//
// Service Center -> Basestation
type DetPrpCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDetPrpCmp(opId int64) DetPrpCmp {
	return DetPrpCmp{OpId: opId, Command: structs.MsgDetPrpCmp}
}

func (m *DetPrpCmp) GetOpId() int64 {
	return m.OpId
}

func (m *DetPrpCmp) GetCommand() structs.Command {
	return structs.MsgDetPrpCmp
}
