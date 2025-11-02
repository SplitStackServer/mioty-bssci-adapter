package messages

import (
	"errors"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"

	"github.com/SplitStackServer/splitstack/api/go/v5/bs"
)

//go:generate msgp

// Activate variable mac support
//
// Service Center -> Basestation
type VmActivate struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// MAC-Type of the intended Variable MAC
	MacType uint32 `msg:"macType" json:"macType"`
}

func NewVmActivate(opId int64, macType uint32) VmActivate {
	return VmActivate{Command: structs.MsgVmActivate, OpId: opId, MacType: macType}
}

func NewVmActivateFromProto(opId int64, pb *bs.EnableVariableMac) (*VmActivate, error) {
	if pb != nil {
		m := NewVmActivate(opId, pb.MacType)
		return &m, nil
	}
	return nil, errors.New("invalid EnableVariableMac command")
}

func (m *VmActivate) GetOpId() int64 {
	return m.OpId
}

func (m *VmActivate) GetCommand() structs.Command {
	return structs.MsgVmActivate
}

// implements ServerMessage
func (m *VmActivate) SetOpId(opId int64) {
	m.OpId = opId
}

// VmActivate response
//
// Basestation -> Service Center
type VmActivateRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewVmActivateRsp(opId int64) VmActivateRsp {
	return VmActivateRsp{Command: structs.MsgVmActivateRsp, OpId: opId}
}

func (m *VmActivateRsp) GetOpId() int64 {
	return m.OpId
}

func (m *VmActivateRsp) GetCommand() structs.Command {
	return structs.MsgVmActivateRsp
}

// VmActivate complete
//
// Service Center -> Basestation
type VmActivateCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewVmActivateCmp(opId int64) VmActivateCmp {
	return VmActivateCmp{Command: structs.MsgVmActivateCmp, OpId: opId}
}

func (m *VmActivateCmp) GetOpId() int64 {
	return m.OpId
}

func (m *VmActivateCmp) GetCommand() structs.Command {
	return structs.MsgVmActivateCmp
}
