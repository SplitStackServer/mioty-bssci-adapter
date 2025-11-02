package messages

import (
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v5/bs"
)

//go:generate msgp
//msgp:shim common.EUI64 as:uint64 using:common.Eui64toUnsignedInt/common.Eui64FromUnsignedInt

// Downlink RX Status
//
// The DL RX Status operation is initiated by the Base Station after receiving a DL RX status
// response control segment from an End Point.
//
// Basestation -> Service Center
type DlRxStat struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui" json:"epEui"`
	// Unix UTC time of reception, center of last subpacket, 64 bit, ns resolution
	RxTime uint64 `msg:"rxTime" json:"rxTime"`
	// End Point packet counter
	PacketCnt uint32 `msg:"packetCnt" json:"packetCnt"`
	// End Point DL reception signal to noise ratio in dB
	DlRxSnr float64 `msg:"dlRxSnr" json:"dlRxSnr"`
	// End Point DL reception signal strength in dBm
	DlRxRssi float64 `msg:"dlRxRssi" json:"dlRxRssi"`
}

func NewDlRxStat(
	opId int64,
	epEui common.EUI64,
	result string,
	rxTime uint64,
	packetCnt uint32,
	dlRxSnr float64,
	dlRxRssi float64,
) DlRxStat {
	return DlRxStat{
		Command:   structs.MsgDlRxStat,
		OpId:      opId,
		EpEui:     epEui,
		RxTime:    rxTime,
		PacketCnt: packetCnt,
		DlRxSnr:   dlRxSnr,
		DlRxRssi:  dlRxRssi,
	}
}

func (m *DlRxStat) GetOpId() int64 {
	return m.OpId
}

func (m *DlRxStat) GetCommand() structs.Command {
	return structs.MsgDlRxStat
}

// implements BasestationMessage.GetEventType()
func (m *DlRxStat) GetEventType() events.EventType {
	return events.EventTypeEpRx
}

// implements BasestationMessage.IntoProto()
func (m *DlRxStat) IntoProto(bsEui *common.EUI64) *bs.BasestationUplink {
	var message bs.BasestationUplink
	if m != nil {
		bsEuiB := bsEui.String()
		epEuiB := m.EpEui.String()

		now := getNow().UnixNano()
		ts := TimestampNsToProto(now)

		message = bs.BasestationUplink{
			Ts:    ts,
			BsEui: bsEuiB,
			OpId:  m.OpId,
			Message: &bs.BasestationUplink_DlRxStat{
				DlRxStat: &bs.BasestationDownlinkRxStatus{
					EpEui:     epEuiB,
					RxTime:    TimestampNsToProto(int64(m.RxTime)),
					PacketCnt: m.PacketCnt,
					DlRxRssi:  m.DlRxRssi,
					DlRxSnr:   m.DlRxSnr,
				},
			},
		}
	}

	return &message
}

// Downlink RX Status response
//
// Service Center -> Basestation
type DlRxStatRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDlRxStatRsp(opId int64) DlRxStatRsp {
	return DlRxStatRsp{
		Command: structs.MsgDlRxStatRsp,
		OpId:    opId,
	}
}

func (m *DlRxStatRsp) GetOpId() int64 {
	return m.OpId
}

func (m *DlRxStatRsp) GetCommand() structs.Command {
	return structs.MsgDlRxStatRsp
}

// Downlink RX Status complete
//
// Basestation -> Service Center
type DlRxStatCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDlRxStatCmp(opId int64) DlRxStatCmp {
	return DlRxStatCmp{OpId: opId, Command: structs.MsgDlRxStatCmp}
}

func (m *DlRxStatCmp) GetOpId() int64 {
	return m.OpId
}

func (m *DlRxStatCmp) GetCommand() structs.Command {
	return structs.MsgDlRxStatCmp
}
