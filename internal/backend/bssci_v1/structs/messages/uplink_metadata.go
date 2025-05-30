package messages

import "github.com/SplitStackServer/splitstack/api/go/v4/bs"

type UplinkMetadata struct {
	RxTime     uint64      `json:"rxTime"`
	RxDuration *uint64     `json:"rxDuration,omitempty"`
	PacketCnt  uint32      `json:"packetCnt"`
	Profile    *string     `json:"profile,omitempty"`
	SNR        float64     `json:"snr"`
	RSSI       float64     `json:"rssi"`
	EqSnr      *float64    `json:"eqSnr,omitempty"`
	Subpackets *Subpackets `json:"subpackets,omitempty"`
}

func (m *UplinkMetadata) IntoProto() *bs.EndnodeUplinkMetadata {
	var message bs.EndnodeUplinkMetadata
	if m != nil {

		rxTime := TimestampNsToProto(int64(m.RxTime))

		message = bs.EndnodeUplinkMetadata{
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
