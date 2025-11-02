package messages

import (
	"errors"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"

	"github.com/SplitStackServer/splitstack/api/go/v5/bs"
)

//go:generate msgp

// Deactivate variable mac support
//
// Service Center -> Basestation
type VmDeactivate struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// MAC-Type of the intended Variable MAC
	MacType uint32 `msg:"macType" json:"macType"`
}

func NewVmDeactivate(opId int64, macType uint32) VmDeactivate {
	return VmDeactivate{Command: structs.MsgVmDeactivate, OpId: opId, MacType: macType}
}

func NewVmDeactivateFromProto(opId int64, pb *bs.DisableVariableMac) (*VmDeactivate, error) {
	if pb != nil {
		m := NewVmDeactivate(opId, pb.MacType)
		return &m, nil
	}
	return nil, errors.New("invalid DisableVariableMac command")
}

func (m *VmDeactivate) GetOpId() int64 {
	return m.OpId
}

func (m *VmDeactivate) GetCommand() structs.Command {
	return structs.MsgVmDeactivate
}

// implements ServerMessage
func (m *VmDeactivate) SetOpId(opId int64) {
	m.OpId = opId
}

// VmDeactivate response
//
// Basestation -> Service Center
type VmDeactivateRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewVmDeactivateRsp(opId int64) VmDeactivateRsp {
	return VmDeactivateRsp{Command: structs.MsgVmDeactivateRsp, OpId: opId}
}

func (m *VmDeactivateRsp) GetOpId() int64 {
	return m.OpId
}

func (m *VmDeactivateRsp) GetCommand() structs.Command {
	return structs.MsgVmDeactivateRsp
}

// VmDeactivate complete
//
// Service Center -> Basestation
type VmDeactivateCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewVmDeactivateCmp(opId int64) VmDeactivateCmp {
	return VmDeactivateCmp{Command: structs.MsgVmDeactivateCmp, OpId: opId}
}

func (m *VmDeactivateCmp) GetOpId() int64 {
	return m.OpId
}

func (m *VmDeactivateCmp) GetCommand() structs.Command {
	return structs.MsgVmDeactivateCmp
}
