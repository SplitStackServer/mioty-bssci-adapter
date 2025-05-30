package messages

import (
	"errors"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

//go:generate msgp
//msgp:shim common.EUI64 as:uint64 using:common.Eui64toUnsignedInt/common.Eui64FromUnsignedInt

// Attach Propagate
//
// The attach propagate operation is initiated by the Service Center to propagate an End
// Point attachment to the Base Station. The attachment information can either be
// acquired via an over the air attachment at another Base Station or in the form of an
// offline preattachment of an End Point (as required for unidirectional End Points).
//
// Service Center -> Basestation
type AttPrp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// End Point EUI64
	EpEui common.EUI64 `msg:"epEui" json:"epEui"`
	// True if End Point is bidirectional
	Bidi bool `msg:"bidi" json:"bidi"`
	// 16 Byte End Point network session key
	NwkSessionKey [16]byte `msg:"nwkSessionKey" json:"nwkSessionKey"`
	// End Point short address
	ShAddr uint16 `msg:"shAddr" json:"shAddr"`
	// Last known End Point packet counter
	LastPacketCount uint32 `msg:"lastPacketCount" json:"lastPacketCount"`
	// True if End Point uses dual channel mode
	DualChan bool `msg:"dualChan" json:"dualChan"`
	// True if End Point uses DL repetition
	Repetition bool `msg:"repetition" json:"repetition"`
	// True if End Point uses wide carrier offset
	WideCarrOff bool `msg:"wideCarrOff" json:"wideCarrOff"`
	// True if End Point uses long DL interblock distance
	LongBlkDist bool `msg:"longBlkDist" json:"longBlkDist"`
}

func NewAttPrp(opId int64, epEui common.EUI64, bidi bool, nwkSessionKey [16]byte, shAddr uint16, lastPacketCount uint32, dualChan bool, repetition bool, wideCarrOff bool, longBlkDist bool,
) AttPrp {
	return AttPrp{
		Command:         structs.MsgAttPrp,
		OpId:            opId,
		EpEui:           epEui,
		Bidi:            bidi,
		NwkSessionKey:   nwkSessionKey,
		ShAddr:          shAddr,
		LastPacketCount: lastPacketCount,
		DualChan:        dualChan,
		Repetition:      repetition,
		WideCarrOff:     wideCarrOff,
		LongBlkDist:     longBlkDist,
	}
}

func NewAttPrpFromProto(opId int64, pb *bs.AttachPropagate) (*AttPrp, error) {
	if pb != nil {
		epEui, err := common.Eui64FromHexString(pb.EndnodeEui)
		if err != nil {
			return nil, err
		}

		if len(pb.NwkSessionKey) != 16 {
			return nil, errors.New("invalid NwkSessionKey")
		}

		m := NewAttPrp(
			opId,
			epEui,
			pb.Bidi,
			[16]byte(pb.NwkSessionKey),
			uint16(pb.ShAddr),
			uint32(pb.LastPacketCnt),
			pb.DualChannel,
			pb.Repetition,
			pb.WideCarrOff,
			pb.LongBlkDist,
		)
		return &m, nil
	}
	return nil, errors.New("invalid AttachPropagate command")
}

func (m *AttPrp) GetOpId() int64 {
	return m.OpId
}

func (m *AttPrp) GetCommand() structs.Command {
	return structs.MsgAttPrp
}

// implements ServerMessage
func (m *AttPrp) SetOpId(opId int64) {
	m.OpId = opId
}

// Attach propagate response
//
// Basestation -> Service Center
type AttPrpRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewAttPrpRsp(opId int64) AttPrpRsp {
	return AttPrpRsp{OpId: opId, Command: structs.MsgAttPrpRsp}
}

func (m *AttPrpRsp) GetOpId() int64 {
	return m.OpId
}

func (m *AttPrpRsp) GetCommand() structs.Command {
	return structs.MsgAttPrpRsp
}

// Attach propagate complete
//
// Service Center -> Basestation
type AttPrpCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewAttPrpCmp(opId int64) AttPrpCmp {
	return AttPrpCmp{OpId: opId, Command: structs.MsgAttPrpCmp}
}

func (m *AttPrpCmp) GetOpId() int64 {
	return m.OpId
}

func (m *AttPrpCmp) GetCommand() structs.Command {
	return structs.MsgAttPrpCmp
}
