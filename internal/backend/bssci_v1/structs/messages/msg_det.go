package messages

import (
	"encoding/binary"
	"errors"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

//go:generate msgp
//msgp:shim common.EUI64 as:uint64 using:common.Eui64toUnsignedInt/common.Eui64FromUnsignedInt

// Detach
//
// The detach operation is initiated by the Base Station after receiving an over the air
// detachment request from an End Point.
//
// Basestation -> Service Center
type Det struct {
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
	// Subpackets object with reception info for every subpacket, optional
	Subpackets *Subpackets `msg:"subpackets,omitempty" json:"subpackets,omitempty"`
	// End Point signature
	Sign [4]byte `msg:"sign" json:"sign"`
}

func NewDet(
	opId int64,
	epEui common.EUI64,
	rxTime uint64,
	rxDuration *uint64,
	packetCnt uint32,
	snr float64,
	rssi float64,
	eqSnr *float64,
	profile *string,
	subpackets *Subpackets,
	sign [4]byte,
) Det {
	return Det{
		Command:    structs.MsgDet,
		OpId:       opId,
		EpEui:      epEui,
		RxTime:     rxTime,
		RxDuration: rxDuration,
		PacketCnt:  packetCnt,
		RSSI:       rssi,
		SNR:        snr,
		Profile:    profile,
		Subpackets: subpackets,
		Sign:       sign,
	}
}

func (m *Det) GetOpId() int64 {
	return m.OpId
}

func (m *Det) GetCommand() structs.Command {
	return structs.MsgDet
}

// implements EndnodeMessage.GetEventType()
func (m *Det) GetEventType() events.EventType {
	return events.EventTypeEpOtaa
}

// implements EndnodeMessage.IntoProto()
func (m *Det) IntoProto(bsEui *common.EUI64) *bs.EndnodeUplink {
	bsEuiB := bsEui.String()
	epEuiB := m.EpEui.String()

	sign := binary.LittleEndian.Uint32(m.Sign[:])

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
		Message: &bs.EndnodeUplink_Det{
			Det: &bs.EndnodeDetMessage{

				OpId: m.OpId,
				EpEui: epEuiB,
				Sign: sign,
			},
		},
		Meta: metadata.IntoProto(),
	}
	return &message
}

// Detach response
//
// Service Center -> Basestation
type DetRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// End Point signature
	Sign [4]byte `msg:"sign" json:"sign"`
}

func NewDetRsp(opId int64, sign [4]byte) DetRsp {
	return DetRsp{
		Command: structs.MsgDetRsp,
		OpId:    opId,
		Sign:    sign,
	}
}

func NewDetRspFromProto(opId int64, pb *bs.EndnodeDetachResponse) (*DetRsp, error) {
	if pb != nil {
		sign := pb.Sign
		signB := [4]byte{
			byte(0xff & sign),
			byte(0xff & (sign >> 8)),
			byte(0xff & (sign >> 16)),
			byte(0xff & (sign >> 24)),
		}

		msg := NewDetRsp(opId, signB)
		return &msg, nil
	}
	return nil, errors.New("invalid EndnodeDetachResponse command")
}

func (m *DetRsp) GetOpId() int64 {
	return m.OpId
}

func (m *DetRsp) GetCommand() structs.Command {
	return structs.MsgDetRsp
}

// Detach complete
//
// Basestation -> Service Center
type DetCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewDetCmp(opId int64) DetCmp {
	return DetCmp{OpId: opId, Command: structs.MsgDetCmp}
}

func (m *DetCmp) GetOpId() int64 {
	return m.OpId
}

func (m *DetCmp) GetCommand() structs.Command {
	return structs.MsgDetCmp
}
