package messages

import (
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

//go:generate msgp
//msgp:shim common.EUI64 as:uint64 using:common.Eui64toUnsignedInt/common.Eui64FromUnsignedInt

// Uplink data
//
// The UL data operation is initiated by the Base Station after receiving uplink data from
// an End Point. Telegrams carrying control data exclusively are considered as empty data.
//
// Basestation -> Service Center
type UlData struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui" json:"epEui"`
	// Unix UTC time of reception, center of last subpacket, 64 bit, ns resolution
	RxTime uint64 `msg:"rxTime" json:"rxTime"`
	// Duration of the reception, center of first subpacket to center of last subpacket in ns, optional
	RxDuration *uint64 `msg:"rxDuration,omitempty" json:"rxDuration,omitempty"`
	// End Point packet counter
	PacketCnt uint32 `msg:"packetCnt" json:"packetCnt"`
	// Reception signal to noise ratio in dB
	SNR float64 `msg:"snr" json:"snr"`
	// Reception signal strength in dBm
	RSSI float64 `msg:"rssi" json:"rssi"`
	// AWGN equivalent reception SNR in dB, optional
	EqSnr *float64 `msg:"eqSnr,omitempty" json:"eqSnr,omitempty"`
	// Name of the Mioty profile used for reception, i.e. eu1, optional
	Profile *string `msg:"profile,omitempty" json:"profile,omitempty"`
	// Mioty mode and variant used for reception, i.e. ulp, ulp-rep, ulp-ll, optional
	Mode *string `msg:"mode,omitempty" json:"mode,omitempty"`
	// Subpackets object with reception info for every subpacket, optional
	Subpackets *Subpackets `msg:"subpackets,omitempty" json:"subpackets,omitempty"`
	// End Point user data, might be empty
	UserData []uint8 `msg:"userData" json:"userData"`
	// User data format identifier, 8 bit, optional, default 0
	Format *byte `msg:"format,omitempty" json:"format,omitempty"`
	// True if End Point downlink window is opened
	DlOpen bool `msg:"dlOpen" json:"dlOpen"`
	// True if End Point expects a response in the DL window, requires dlOpen
	ResponseExp bool `msg:"responseExp" json:"responseExp"`
	// True if End Point acknowledges the reception of a DL transmission in the last DL window (packetCnt - 1)
	DlAck bool `msg:"dlAck" json:"dlAck"`
}

func NewUlData(
	opId int64,
	epEui common.EUI64,
	rxTime uint64,
	rxDuration *uint64,
	packetCnt uint32,
	snr float64,
	rssi float64,
	eqSnr *float64,
	profile *string,
	mode *string,
	subpackets *Subpackets,
	userData []byte,
	format *byte,
	dlOpen bool,
	responseExp bool,
	dlAck bool,
) UlData {
	return UlData{
		Command:     structs.MsgUlData,
		OpId:        opId,
		EpEui:       epEui,
		RxTime:      rxTime,
		RxDuration:  rxDuration,
		PacketCnt:   packetCnt,
		SNR:         snr,
		RSSI:        rssi,
		EqSnr:       eqSnr,
		Profile:     profile,
		Mode:        mode,
		Subpackets:  subpackets,
		UserData:    userData,
		Format:      format,
		DlOpen:      dlOpen,
		ResponseExp: responseExp,
		DlAck:       dlAck,
	}
}

func (m *UlData) GetOpId() int64 {
	return m.OpId
}

func (m *UlData) GetCommand() structs.Command {
	return structs.MsgUlData
}

// implements EndnodeMessage.GetEventType()
func (m *UlData) GetEventType() events.EventType {
	return events.EventTypeEpUl
}

// implements EndnodeMessage.IntoProto()
func (m *UlData) IntoProto(bsEui *common.EUI64) *bs.EndnodeUplink {
	bsEuiB := bsEui.String()
	epEuiB := m.EpEui.String()

	var format uint32
	if m.Format == nil {
		format = 0
	} else {
		format = uint32(*m.Format)
	}

	metadata := UplinkMetadata{
		RxTime:     m.RxTime,
		RxDuration: m.RxDuration,
		PacketCnt:  m.PacketCnt,
		Profile:    m.Profile,
		SNR:        m.SNR,
		RSSI:       m.RSSI,
		EqSnr:      m.EqSnr,
		Subpackets: m.Subpackets,
	}

	message := bs.EndnodeUplink{
		BsEui:      bsEuiB,
		Message: &bs.EndnodeUplink_UlData{
			UlData: &bs.EndnodeUlDataMessage{
				EpEui: epEuiB,
				Data:   m.UserData,
				Format: format,
				Mode:   m.Mode,
				DlAck:  m.DlAck,
				DlOpen: m.DlOpen,
			},
		},
		Meta:   metadata.IntoProto(),
	}
	return &message
}

// Uplink data response
//
// Service Center -> Basestation
type UlDataRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewUlDataRsp(opId int64) UlDataRsp {
	return UlDataRsp{
		Command: structs.MsgUlDataRsp,
		OpId:    opId,
	}
}

func (m *UlDataRsp) GetOpId() int64 {
	return m.OpId
}

func (m *UlDataRsp) GetCommand() structs.Command {
	return structs.MsgUlDataRsp
}

// UlDataach complete
//
// Basestation -> Service Center
type UlDataCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewUlDataCmp(opId int64) UlDataCmp {
	return UlDataCmp{OpId: opId, Command: structs.MsgUlDataCmp}
}

func (m *UlDataCmp) GetOpId() int64 {
	return m.OpId
}

func (m *UlDataCmp) GetCommand() structs.Command {
	return structs.MsgUlDataCmp
}
