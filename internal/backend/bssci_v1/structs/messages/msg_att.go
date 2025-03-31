package messages

import (
	"encoding/binary"
	"errors"
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/api/rsp"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
)

//go:generate msgp
//msgp:shim common.EUI64 as:int64 using:common.Eui64toInt/common.Eui64FromInt

// Attach
//
// The attach operation is initiated by the Base Station after receiving an over the air
// attachment request from an End Point.
//
// Basestation -> Service Center
type Att struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui" json:"epEui"`
	// Unix UTC time of reception, center of last subpacket, 64 bit, ns resolution
	RxTime uint64 `msg:"rxTime" json:"rxTime"`
	// Duration of the reception, center of first subpacket to center of last subpacket in ns, optional
	RxDuration *uint64 `msg:"rxDuration,omitempty" json:"rxDuration,omitempty"`
	// End Point attachment counter
	AttachCnt uint32 `msg:"attachCnt" json:"attachCnt"`
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
	// End Point nonce
	Nonce [4]byte `msg:"nonce" json:"nonce"`
	// End Point signature
	Sign [4]byte `msg:"sign" json:"sign"`
	// End Point short address, only if assigned by the Base Station
	ShAddr *uint16 `msg:"shAddr,omitempty" json:"shAddr,omitempty"`
	// True if End Point uses dual channel mode
	DualChan bool `msg:"dualChan" json:"dualChan"`
	// True if End Point uses DL repetition
	Repetition bool `msg:"repetition" json:"repetition"`
	// True if End Point uses wide carrier offset
	WideCarrOff bool `msg:"wideCarrOff" json:"wideCarrOff"`
	// True if End Point uses long DL interblock distance
	LongBlkDist bool `msg:"longBlkDist" json:"longBlkDist"`
}

func NewAtt(
	opId int64,
	epEui common.EUI64,
	rxTime uint64,
	rxDuration *uint64,
	attachCnt uint32,
	snr float64,
	rssi float64,
	eqSnr *float64,
	profile *string,
	subpackets *Subpackets,
	nonce [4]byte,
	sign [4]byte,
	shAddr *uint16,
	dualChan bool,
	repetition bool,
	wideCarrOff bool,
	longBlkDist bool,
) Att {
	return Att{
		Command:     structs.MsgAtt,
		OpId:        opId,
		EpEui:       epEui,
		RxTime:      rxTime,
		RxDuration:  rxDuration,
		AttachCnt:   attachCnt,
		RSSI:        rssi,
		SNR:         snr,
		Profile:     profile,
		Subpackets:  subpackets,
		Nonce:       nonce,
		Sign:        sign,
		ShAddr:      shAddr,
		DualChan:    dualChan,
		Repetition:  repetition,
		WideCarrOff: wideCarrOff,
		LongBlkDist: longBlkDist,
	}
}

func (m *Att) GetOpId() int64 {
	return m.OpId
}

func (m *Att) GetCommand() structs.Command {
	return structs.MsgAtt
}

// implements EndnodeMessage.GetEndpointEui()
func (m *Att) GetEndpointEui() common.EUI64 {
	return m.EpEui
}

// implements EndnodeMessage.IntoProto()
func (m *Att) IntoProto(bsEui common.EUI64) *msg.ProtoEndnodeMessage {
	bsEuiB := bsEui.ToUnsignedInt()
	epEuiB := m.EpEui.ToUnsignedInt()

	nonce := binary.LittleEndian.Uint32(m.Nonce[:])

	sign := binary.LittleEndian.Uint32(m.Sign[:])

	var shAddr *uint32

	if m.ShAddr != nil {
		shAddrT := uint32(*m.ShAddr)
		shAddr = &shAddrT
	}

	metadata := UplinkMetadata{
		RxTime:     m.RxTime,
		RxDuration: m.RxDuration,
		PacketCnt:  0,
		Profile:    m.Profile,
		SNR:        m.SNR,
		RSSI:       m.RSSI,
		EqSnr:      m.EqSnr,
		Subpackets: m.Subpackets,
	}

	message := msg.ProtoEndnodeMessage{
		BsEui:      bsEuiB,
		EndnodeEui: epEuiB,
		Message: &msg.ProtoEndnodeMessage_Att{
			Att: &msg.EndnodeAttMessage{
				OpId:          m.OpId,
				AttachmentCnt: m.AttachCnt,
				Nonce:         nonce,
				Sign:          sign,
				ShAddr:        shAddr,
				Meta:          metadata.IntoProto(),
				DualChannel:   m.DualChan,
				Repetition:    m.Repetition,
				WideCarrOff:   m.WideCarrOff,
				LongBlkDist:   m.LongBlkDist,
			},
		},
	}
	return &message
}

// Attach response
//
// Service Center -> Basestation
type AttRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// End Point network session key
	NwkSessionKey [16]byte `msg:"nwkSessionKey" json:"nwkSessionKey"`
	// End Point short address, only if not assigned by the Base Station
	ShAddr *uint16 `msg:"shAddr,omitempty" json:"shAddr,omitempty"`
}

func NewAttRsp(opId int64, nwkSessionKey [16]byte, shAddr *uint16) AttRsp {
	return AttRsp{
		Command:       structs.MsgAttRsp,
		OpId:          opId,
		NwkSessionKey: nwkSessionKey,
		ShAddr:        shAddr,
	}
}

func NewAttRspFromProto(opId int64, pb *rsp.EndnodeAttachResponse) (*AttRsp, error) {
	if pb != nil {
		var shAddr *uint16
		if shAddrPb := pb.ShAddr; shAddrPb != nil {
			shAddrA := uint16(0xffff & *shAddrPb)
			shAddr = &shAddrA
		}

		if len(pb.NwkSessionKey) != 16 {
			return nil, errors.New("invalid NwkSessionKey")
		}

		msg := NewAttRsp(opId, [16]byte(pb.NwkSessionKey), shAddr)
		return &msg, nil
	}
	return nil, errors.New("invalid EndnodeAttachResponse command")
}

func (m *AttRsp) GetOpId() int64 {
	return m.OpId
}

func (m *AttRsp) GetCommand() structs.Command {
	return structs.MsgAttRsp
}

// Attach complete
//
// Basestation -> Service Center
type AttCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operationF
	OpId int64 `msg:"opId" json:"opId"`
}

func NewAttCmp(opId int64) AttCmp {
	return AttCmp{OpId: opId, Command: structs.MsgAttCmp}
}

func (m *AttCmp) GetOpId() int64 {
	return m.OpId
}

func (m *AttCmp) GetCommand() structs.Command {
	return structs.MsgAttCmp
}
