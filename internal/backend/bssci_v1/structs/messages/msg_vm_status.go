package messages

import (
	"errors"
	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
)

//go:generate msgp

// The VM Status operation delivers a list of the activated MAC-Types.
//
// Service Center -> Basestation
type VmStatus struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewVmStatus(opId int64) VmStatus {
	return VmStatus{Command: structs.MsgVmStatus, OpId: opId}
}

func NewVmStatusFromProto(opId int64, pb *cmd.RequestVariableMacStatus) (*VmStatus, error) {
	if pb != nil {
		m := NewVmStatus(opId)
		return &m, nil
	}
	return nil, errors.New("invalid RequestVariableMacStatus command")
}

func (m *VmStatus) GetOpId() int64 {
	return m.OpId
}

func (m *VmStatus) GetCommand() structs.Command {
	return structs.MsgVmStatus
}

// implements ServerMessage
func (m *VmStatus) SetOpId(opId int64) {
	m.OpId = opId
}

// VmStatus response
//
// Basestation -> Service Center
type VmStatusRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// // List of activated macTypes
	MacTypes []int64 `msg:"macTypes" json:"macTypes"`
}

func NewVmStatusRsp(opId int64, macTypes []int64) VmStatusRsp {
	return VmStatusRsp{Command: structs.MsgVmStatusRsp, OpId: opId, MacTypes: macTypes}
}

func (m *VmStatusRsp) GetOpId() int64 {
	return m.OpId
}

func (m *VmStatusRsp) GetCommand() structs.Command {
	return structs.MsgVmStatusRsp
}

// implements BasestationMessage.IntoProto()
func (m *VmStatusRsp) IntoProto(bsEui *common.EUI64) *msg.ProtoBasestationMessage {

	var message msg.ProtoBasestationMessage

	if m != nil && bsEui != nil {
		bsEuiB := bsEui.ToUnsignedInt()

		message = msg.ProtoBasestationMessage{
			BsEui: bsEuiB,
			V1: &msg.ProtoBasestationMessage_VmStatus{
				VmStatus: &msg.BasestationVariableMacStatus{
					MacTypes: m.MacTypes,
				},
			},
		}
	}

	return &message
}

// VmStatus complete
//
// Service Center -> Basestation
type VmStatusCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewVmStatusCmp(opId int64) VmStatusCmp {
	return VmStatusCmp{Command: structs.MsgVmStatusCmp, OpId: opId}
}

func (m *VmStatusCmp) GetOpId() int64 {
	return m.OpId
}

func (m *VmStatusCmp) GetCommand() structs.Command {
	return structs.MsgVmStatusCmp
}
