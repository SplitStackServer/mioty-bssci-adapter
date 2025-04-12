package messages

import (
	"errors"
	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
)

//go:generate msgp
//msgp:shim common.EUI64 as:uint64 using:common.Eui64toUnsignedInt/common.Eui64FromUnsignedInt

// Downlink data queue
//
// The DL data queue operation is initiated by the Service Center to schedule downlink
// data at the Base Station for an End Point. This might be done either within the interval
// between an uplink message and the according downlink window for direct responses or
// a priority for predefined downlink data.
//
// Counter dependent downlink data (i.e. due to application encryption) must be provided
// for one or multiple specific packet counters. It can only be transmitted in a downlink
// window with a matching counter. Only one downlink packet is transmitted for one queue
// operation, using the first available and suitable downlink window. If user data is empty,
// a pure acknowledgement downlink is queued.
//
// Service Center -> Basestation
type DlDataQue struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui" json:"epEui"`
	// Assigned queue ID for reference, 64 bit
	QueId uint64 `msg:"queId" json:"queId"`
	// True if userData is counter dependent
	CntDepend bool `msg:"cntDepend" json:"cntDepend"`
	// End Point packet counter for which the according userData entry is valid, omitted if cntDepend is false
	PacketCnt *[]uint32 `msg:"packetCnt,omitempty" json:"packetCnt,omitempty"`
	// End Point user data for each of the m packet counters, single user data entry if cntDepend is false
	UserData [][]uint8 `msg:"userData" json:"userData"`
	// User data format identifier, 8 bit, optional, default 0
	Format *byte `msg:"format,omitempty" json:"format,omitempty"`
	// Priority, higher values are prioritized, single precision floating point, optional, default 0
	Prio *float32 `msg:"prio,omitempty" json:"prio,omitempty"`
	// True to request End Point response, optional
	ResponseExp *bool `msg:"responseExp,omitempty" json:"responseExp,omitempty"`
	// True to request priority End Point response, optional
	ResponsePrio *bool `msg:"responsePrio,omitempty" json:"responsePrio,omitempty"`
	// True to request further End Point DL window, optional
	DlWindReq *bool `msg:"dlWindReq,omitempty" json:"dlWindReq,omitempty"`
	// True to send downlink only if End Point expects a response, optional
	ExpOnly *bool `msg:"expOnly,omitempty" json:"expOnly,omitempty"`
}

// new downlink with unencrypted data
func NewDlDataQue(
	opId int64,
	epEui common.EUI64,
	queId uint64,
	prio *float32,
	format *byte,
	userData []byte,
	responseExp *bool,
	responsePrio *bool,
	dlWindReq *bool,
	expOnly *bool,

) DlDataQue {
	return DlDataQue{
		Command:      structs.MsgDlDataQue,
		OpId:         opId,
		EpEui:        epEui,
		QueId:        queId,
		CntDepend:    false,
		Format:       format,
		UserData:     [][]byte{userData},
		Prio:         prio,
		ResponseExp:  responseExp,
		ResponsePrio: responsePrio,
		DlWindReq:    dlWindReq,
		ExpOnly:      expOnly,
	}
}

// new downlink with encrypted data
//
// length of packetCnt and userData must match, but not checked here
func NewDlDataQueEnc(
	opId int64,
	epEui common.EUI64,
	queId uint64,
	prio *float32,
	format *byte,
	packetCnt []uint32,
	userData [][]byte,
	responseExp *bool,
	responsePrio *bool,
	dlWindReq *bool,
	expOnly *bool,

) DlDataQue {
	return DlDataQue{
		Command:      structs.MsgDlDataQue,
		OpId:         opId,
		EpEui:        epEui,
		QueId:        queId,
		CntDepend:    true,
		PacketCnt:    &packetCnt,
		UserData:     userData,
		Format:       format,
		Prio:         prio,
		ResponseExp:  responseExp,
		ResponsePrio: responsePrio,
		DlWindReq:    dlWindReq,
		ExpOnly:      expOnly,
	}
}

func NewDlDataQueFromProto(opId int64, pb *cmd.EnqueDownlink) (*DlDataQue, error) {
	if pb != nil {
		var format byte

		if pb.Format != nil {
			format = uint8(0xff & *pb.Format)
		}

		epEui, err := common.Eui64FromHexString(pb.EndnodeEui)
		if err != nil {
			return nil, err
		}

		msg := DlDataQue{
			Command:      structs.MsgDlDataQue,
			OpId:         opId,
			EpEui:        epEui,
			QueId:        pb.DlQueId,
			Format:       &format,
			Prio:         pb.Priority,
			ResponseExp:  pb.ResponseExp,
			ResponsePrio: pb.ResponsePrio,
			DlWindReq:    pb.ReqDlWindow,
			ExpOnly:      pb.OnlyIfExpected,
		}

		switch pb.Payload.(type) {
		case *cmd.EnqueDownlink_Ack:
			return &msg, nil
		case *cmd.EnqueDownlink_Data:
			payload := pb.GetData()
			if len(payload.Data) != 0 {
				msg.UserData = [][]byte{payload.Data}
				return &msg, nil
			} else {
				return nil, errors.New("EnqueDownlink_Data has no payload")
			}
		case *cmd.EnqueDownlink_DataEnc:
			payload := pb.GetDataEnc()
			if len(payload.Data) != 0 && len(payload.PacketCnt) != 0 {
				msg.CntDepend = true
				msg.UserData = payload.Data
				msg.PacketCnt = &payload.PacketCnt
				return &msg, nil
			} else {
				return nil, errors.New("EnqueDownlink_DataEnc has no payload or counter")
			}
		default:
			return nil, errors.New("EnqueDownlink has no payload")
		}
	}
	return nil, errors.New("invalid EnqueDownlink command")
}

func (m *DlDataQue) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataQue) GetCommand() structs.Command {
	return structs.MsgDlDataQue
}

// implements ServerMessage
func (m *DlDataQue) SetOpId(opId int64) {
	m.OpId = opId
}

// Downlink data queue response
//
// Basestation -> Service Center
type DlDataQueRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDlDataQueRsp(opId int64) DlDataQueRsp {
	return DlDataQueRsp{
		Command: structs.MsgDlDataQueRsp,
		OpId:    opId,
	}
}

func (m *DlDataQueRsp) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataQueRsp) GetCommand() structs.Command {
	return structs.MsgDlDataQueRsp
}

// Downlink data queue complete
//
// Service Center -> Basestation
type DlDataQueCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDlDataQueCmp(opId int64) DlDataQueCmp {
	return DlDataQueCmp{OpId: opId, Command: structs.MsgDlDataQueCmp}
}

func (m *DlDataQueCmp) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataQueCmp) GetCommand() structs.Command {
	return structs.MsgDlDataQueCmp
}
