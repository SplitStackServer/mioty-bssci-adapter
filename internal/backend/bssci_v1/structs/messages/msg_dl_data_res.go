package messages

import (
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
)

type dlDataResult int

const (
	dlDataResult_Invalid dlDataResult = iota
	dlDataResult_Sent
	dlDataResult_Expired
)

var dlDataResultName = map[dlDataResult]string{
	dlDataResult_Sent:    "sent",
	dlDataResult_Expired: "expired",
	dlDataResult_Invalid: "invalid",
}

var dlDataResultValue = map[string]dlDataResult{
	"sent":    dlDataResult_Sent,
	"expired": dlDataResult_Expired,
	"invalid": dlDataResult_Invalid,
}

func (e dlDataResult) String() string {
	return dlDataResultName[e]
}

func ParseDlDataResult(s string) dlDataResult {
	return dlDataResultValue[s]
}

//go:generate msgp
//msgp:shim common.EUI64 as:int64 using:common.Eui64toInt/common.Eui64FromInt
//msgp:shim dlDataResult as:string using:(dlDataResult).String/ParseDlDataResult

// Downlink data result
//
// The DL data result operation is initiated by the Base Station after queued DL data has
// either been sent or discarded.
//
// Basestation -> Service Center
type DlDataRes struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui" json:"epEui"`
	// Assigned queue ID for reference, 64 bit
	QueId uint64 `msg:"queId" json:"queId"`
	// sent, expired, invalid
	Result dlDataResult `msg:"result" json:"result"`
	// Unix UTC time of transmission, center of first subpacket, 64 bit, ns resolution, only if result is sent
	TxTime *uint64 `msg:"txTime" json:"txTime"`
	// End Point packet counter, only if result is “sent”
	PacketCnt *uint32 `msg:"packetCnt" json:"packetCnt"`
}

func NewDlDataRes(
	opId int64,
	epEui common.EUI64,
	queId uint64,
	result dlDataResult,
	txTime *uint64,
	packetCnt *uint32,
) DlDataRes {
	return DlDataRes{
		Command:   structs.MsgDlDataRes,
		OpId:      opId,
		EpEui:     epEui,
		QueId:     queId,
		Result:    result,
		TxTime:    txTime,
		PacketCnt: packetCnt,
	}
}

func (m *DlDataRes) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataRes) GetCommand() structs.Command {
	return structs.MsgDlDataRes
}

// implements EndnodeMessage.GetEndpointEui()
func (m *DlDataRes) GetEndpointEui() common.EUI64 {
	return m.EpEui
}

// implements EndnodeMessage.IntoProto()
func (m *DlDataRes) IntoProto(bsEui common.EUI64) *msg.ProtoEndnodeMessage {
	bsEuiB := bsEui.ToUnsignedInt()
	epEuiB := m.EpEui.ToUnsignedInt()

	result := msg.EndnodeDownlinkResult{
		DlQueId: m.QueId,
	}

	switch m.Result {
	case dlDataResult_Sent:
		result.Result = msg.DownlinkResultEnum_SENT
		result.TxTime = TimestampNsToProto(int64(*m.TxTime))
		result.EpPacketCnt = *m.PacketCnt
	case dlDataResult_Expired:
		result.Result = msg.DownlinkResultEnum_EXPIRED
	case dlDataResult_Invalid:
		result.Result = msg.DownlinkResultEnum_INVALID
	}

	message := msg.ProtoEndnodeMessage{
		BsEui:      bsEuiB,
		EndnodeEui: epEuiB,
		Message: &msg.ProtoEndnodeMessage_DlRes{
			DlRes: &result,
		},
	}
	return &message
}

// Downlink data result response
//
// Service Center -> Basestation
type DlDataResRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDlDataResRsp(opId int64) DlDataResRsp {
	return DlDataResRsp{
		Command: structs.MsgDlDataResRsp,
		OpId:    opId,
	}
}

func (m *DlDataResRsp) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataResRsp) GetCommand() structs.Command {
	return structs.MsgDlDataResRsp
}

// Downlink data result complete
//
// Basestation -> Service Center
type DlDataResCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDlDataResCmp(opId int64) DlDataResCmp {
	return DlDataResCmp{OpId: opId, Command: structs.MsgDlDataResCmp}
}

func (m *DlDataResCmp) GetOpId() int64 {
	return m.OpId
}

func (m *DlDataResCmp) GetCommand() structs.Command {
	return structs.MsgDlDataResCmp
}
