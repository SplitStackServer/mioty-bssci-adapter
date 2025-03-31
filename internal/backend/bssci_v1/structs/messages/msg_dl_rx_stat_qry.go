package messages

import (
	"errors"
	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
)

//go:generate msgp
//msgp:shim common.EUI64 as:int64 using:common.Eui64toInt/common.Eui64FromInt

// Downlink rx status query
//
// The DL RX status query operation is initiated by the Service Center to schedule a DL RX
// status query control segment for the next downlink transmission of the Base Station to
// an End Point.
//
// Service Center -> Basestation
type DlRxStatQry struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui" json:"epEui"`
}

func NewDlRxStatQry(opId int64, epEui common.EUI64) DlRxStatQry {
	return DlRxStatQry{Command: structs.MsgDlRxStatQry, OpId: opId, EpEui: epEui}
}

func NewDlRxStatQryFromProto(opId int64, pb *cmd.DownlinkRxStatusQuery) (*DlRxStatQry, error) {
	if pb != nil {
		epEui := common.Eui64FromUnsignedInt(pb.EndnodeEui)
		m := NewDlRxStatQry(opId, epEui)
		return &m, nil
	}
	return nil, errors.New("invalid DownlinkRxStatusQuery command")
}

func (m *DlRxStatQry) GetOpId() int64 {
	return m.OpId
}

func (m *DlRxStatQry) GetCommand() structs.Command {
	return structs.MsgDlRxStatQry
}

// implements ServerMessage
func (m *DlRxStatQry) SetOpId(opId int64) {
	m.OpId = opId
}

// Downlink rx status query response
//
// Basestation <-> Service Center
type DlRxStatQryRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDlRxStatQryRsp(opId int64) DlRxStatQryRsp {
	return DlRxStatQryRsp{OpId: opId, Command: structs.MsgDlRxStatQryRsp}
}

func (m *DlRxStatQryRsp) GetOpId() int64 {
	return m.OpId
}

func (m *DlRxStatQryRsp) GetCommand() structs.Command {
	return structs.MsgDlRxStatQryRsp
}

// Downlink rx status query complete
//
// Service Center <-> Basestation
type DlRxStatQryCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDlRxStatQryCmp(opId int64) DlRxStatQryCmp {
	return DlRxStatQryCmp{OpId: opId, Command: structs.MsgDlRxStatQryCmp}
}

func (m *DlRxStatQryCmp) GetOpId() int64 {
	return m.OpId
}

func (m *DlRxStatQryCmp) GetCommand() structs.Command {
	return structs.MsgDlRxStatQryCmp
}
