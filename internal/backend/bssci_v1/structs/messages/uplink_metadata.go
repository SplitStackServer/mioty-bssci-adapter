package messages

import "github.com/SplitStackServer/splitstack/api/go/v5/bs"

type UplinkMetadata struct {
	OpId       int64       `json:"opId"`
	RxTime     uint64      `json:"rxTime"`
	RxDuration *uint64     `json:"rxDuration,omitempty"`
	PacketCnt  uint32      `json:"packetCnt"`
	Profile    *string     `json:"profile,omitempty"`
	SNR        float64     `json:"snr"`
	RSSI       float64     `json:"rssi"`
	EqSnr      *float64    `json:"eqSnr,omitempty"`
	Subpackets *Subpackets `json:"subpackets,omitempty"`
}

func NewUplinkMetadata(
	opId int64,
	rxTime uint64,
	rxDuration *uint64,
	packetCnt uint32,
	snr float64,
	rssi float64,
	eqSnr *float64,
	profile *string,
	subpackets *Subpackets,
) UplinkMetadata {
	return UplinkMetadata{
		OpId:       opId,
		RxTime:     rxTime,
		RxDuration: rxDuration,
		PacketCnt:  packetCnt,
		SNR:        snr,
		RSSI:       rssi,
		EqSnr:      eqSnr,
		Profile:    profile,
		Subpackets: subpackets,
	}
}

func (m *UplinkMetadata) IntoProto() *bs.EndnodeUplinkMetadata {
	var message bs.EndnodeUplinkMetadata
	if m != nil {

		rxTime := TimestampNsToProto(int64(m.RxTime))

		message = bs.EndnodeUplinkMetadata{
			OpId:      m.OpId,
			RxTime:    rxTime,
			PacketCnt: m.PacketCnt,
			Profile:   m.Profile,
			Rssi:      m.RSSI,
			Snr:       m.SNR,
			EqSnr:     m.EqSnr,
		}
		if m.Subpackets != nil {
			message.SubpacketInfo = m.Subpackets.IntoProto()
		}
		if m.RxDuration != nil {
			message.RxDuration = DurationNsToProto(int64(*m.RxDuration))
		}
	}
	return &message
}
