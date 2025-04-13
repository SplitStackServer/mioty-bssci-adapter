package messages

import "github.com/SplitStackServer/splitstack/api/go/v4/bs"

//go:generate msgp

// Subpackets
//
// reception info for every subpacket
type Subpackets struct {
	// Subpacket signal to noise ratio in dB
	SNR []int32 `msg:"snr" json:"snr"`
	// Subpacket signal strength in dBm
	RSSI []int32 `msg:"rssi" json:"rssi"`
	// Subpacket frequencies in Hz
	Frequency []int32 `msg:"frequency" json:"frequency"`
	// Subpacket phases in degree +-180, optional
	Phase *[]int32 `msg:"phase,omitempty" json:"phase,omitempty"`
}

func (subpackets *Subpackets) IntoProto() []*bs.EndnodeUplinkSubpacket {
	var pb []*bs.EndnodeUplinkSubpacket
	if subpackets != nil {
		pb = make([]*bs.EndnodeUplinkSubpacket, 0, len(subpackets.RSSI))

		if subpackets.Phase == nil {
			for i, v := range subpackets.RSSI {
				proto := bs.EndnodeUplinkSubpacket{
					Snr:       subpackets.SNR[i],
					Rssi:      v,
					Frequency: subpackets.Frequency[i],
				}
				pb = append(pb, &proto)

			}

		} else {
			phase := *subpackets.Phase
			for i, v := range subpackets.RSSI {
				proto := bs.EndnodeUplinkSubpacket{
					Snr:       subpackets.SNR[i],
					Rssi:      v,
					Frequency: subpackets.Frequency[i],
					Phase:     &phase[i],
				}
				pb = append(pb, &proto)

			}
		}
	}
	return pb
}
