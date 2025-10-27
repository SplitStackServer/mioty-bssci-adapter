package messages

import (
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

//go:generate msgp

// The VM UL data operation is initiated by the Base Station after receiving uplink data from
// an End Point using a variable MAC (VM)
//
// Basestation -> Service Center
type VmUlData struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// MAC-Type of the Variable MAC
	MacType uint8 `msg:"macType" json:"macType"`
	// n Byte End Point user data U-MPDU; starting with first byte after MAC-Type
	UserData []uint8 `msg:"userData" json:"userData"`
	// Transceiver time of reception, center of last subpacket, 64 bit, ns resolution
	TrxTime uint64 `msg:"trxTime" json:"trxTime"`
	// Unix UTC time of reception, center of last subpacket, 64 bit, ns resolution
	SysTime uint64 `msg:"sysTime" json:"sysTime"`
	// Frequency offset from center between primary and secondary channel in Hz
	FreqOff float64 `msg:"freqOff" json:"freqOff"`
	// Reception signal to noise ratio in dB
	SNR float64 `msg:"snr" json:"snr"`
	// Reception signal strength in dBm
	RSSI float64 `msg:"rssi" json:"rssi"`
	// AWGN equivalent reception SNR in dB, optional
	EqSnr *float64 `msg:"eqSnr,omitempty" json:"eqSnr,omitempty"`
	// Subpackets object with reception info for every subpacket, optional
	Subpackets *Subpackets `msg:"subpackets,omitempty" json:"subpackets,omitempty"`
	// Carrier spacing step size Bc, 0 = narrow, 1 = standard, 2 = wide
	CarrSpace byte `msg:"carrSpace" json:"carrSpace"`
	// Uplink TSMA Pattern group, 0 = normal, 1 = repetition, 2 = low delay
	PattGrp byte `msg:"pattGrp" json:"pattGrp"`
	// Uplink TSMA Pattern number p
	PattNum byte `msg:"pattNum" json:"pattNum"`
	// Header and payload CRC, crc[0] = header CRC, crc[1] = payload CRC
	CRC [2]uint8 `msg:"crc" json:"crc"`
}

func NewVmUlData(
	opId int64,
	macType uint8,
	userData []byte,
	trxTime uint64,
	freqOff float64,
	snr float64,
	rssi float64,
	eqSnr *float64,
	subpackets *Subpackets,
	carrSpace byte,
	pattGrp byte,
	pattNum byte,
	crc [2]byte,

) VmUlData {
	return VmUlData{
		Command:    structs.MsgVmUlData,
		OpId:       opId,
		MacType:    macType,
		UserData:   userData,
		TrxTime:    trxTime,
		SysTime:    trxTime,
		FreqOff:    freqOff,
		SNR:        snr,
		RSSI:       rssi,
		EqSnr:      eqSnr,
		Subpackets: subpackets,
		CarrSpace:  carrSpace,
		PattGrp:    pattGrp,
		PattNum:    pattNum,
		CRC:        crc,
	}
}

func (m *VmUlData) GetOpId() int64 {
	return m.OpId
}

func (m *VmUlData) GetCommand() structs.Command {
	return structs.MsgVmUlData
}

// implements EndnodeMessage.GetEventType()
func (m *VmUlData) GetEventType() events.EventType {
	return events.EventTypeEpUl
}

// implements EndnodeMessage.IntoProto()
func (m *VmUlData) IntoProto(bsEui *common.EUI64) *bs.EndnodeUplink {
	bsEuiB := bsEui.String()

	crc := uint64(m.CRC[0]) | uint64(m.CRC[0])<<32

	metadata := UplinkMetadata{
		RxTime:     m.SysTime,
		RxDuration: nil,
		PacketCnt:  0,
		Profile:    nil,
		SNR:        m.SNR,
		RSSI:       m.RSSI,
		EqSnr:      m.EqSnr,
		Subpackets: m.Subpackets,
	}

	now := getNow().UnixNano()
	ts := TimestampNsToProto(now)

	message := bs.EndnodeUplink{
		BsEui: bsEuiB,
		Ts:    ts,
		Message: &bs.EndnodeUplink_VmUlData{
			VmUlData: &bs.EndnodeVariableMacUlDataMessage{
				Data:      m.UserData,
				MacType:   uint32(m.MacType),
				FreqOff:   m.FreqOff,
				CarrSpace: bs.CarrierSpacingEnum(m.CarrSpace),
				PattGrp:   bs.TsmaPatternGroupEnum(m.PattGrp),
				PattNum:   uint32(m.PattNum),
				Crc:       crc,
			},
		},
		Meta: metadata.IntoProto(),
	}
	return &message
}

// VmUlData response
//
// Service Center -> Basestation
type VmUlDataRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewVmUlDataRsp(opId int64) VmUlDataRsp {
	return VmUlDataRsp{Command: structs.MsgVmUlDataRsp, OpId: opId}
}

func (m *VmUlDataRsp) GetOpId() int64 {
	return m.OpId
}

func (m *VmUlDataRsp) GetCommand() structs.Command {
	return structs.MsgVmUlDataRsp
}

// VmUlData complete
//
// Basestation -> Service Center
type VmUlDataCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewVmUlDataCmp(opId int64) VmUlDataCmp {
	return VmUlDataCmp{Command: structs.MsgVmUlDataCmp, OpId: opId}
}

func (m *VmUlDataCmp) GetOpId() int64 {
	return m.OpId
}

func (m *VmUlDataCmp) GetCommand() structs.Command {
	return structs.MsgVmUlDataCmp
}
