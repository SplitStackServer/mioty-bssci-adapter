package messages

import (
	"reflect"
	"testing"

	"github.com/SplitStackServer/splitstack/api/go/v5/bs"
)

func TestSubpackets_IntoProto(t *testing.T) {

	testPhase := int32(4)

	type fields struct {
		SNR       []int32
		RSSI      []int32
		Frequency []int32
		Phase     *[]int32
	}
	tests := []struct {
		name   string
		fields fields
		want   []*bs.EndnodeUplinkSubpacket
	}{
		{name: "subpackets1", fields: fields{
			SNR:       []int32{1, 4},
			RSSI:      []int32{2, 5},
			Frequency: []int32{3, 6},
		}, want: []*bs.EndnodeUplinkSubpacket{
			{
				Snr:       1,
				Rssi:      2,
				Frequency: 3,
			},
			{
				Snr:       4,
				Rssi:      5,
				Frequency: 6,
			}}},
		{name: "subpackets2", fields: fields{
			SNR:       []int32{1},
			RSSI:      []int32{2},
			Frequency: []int32{3},
			Phase:     &[]int32{testPhase},
		}, want: []*bs.EndnodeUplinkSubpacket{
			{
				Snr:       1,
				Rssi:      2,
				Frequency: 3,
				Phase:     &testPhase,
			}}},
		{name: "subpackets3", fields: fields{}, want: []*bs.EndnodeUplinkSubpacket{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subpackets := &Subpackets{
				SNR:       tt.fields.SNR,
				RSSI:      tt.fields.RSSI,
				Frequency: tt.fields.Frequency,
				Phase:     tt.fields.Phase,
			}
			if got := subpackets.IntoProto(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subpackets.IntoProto() = %v, want %v", got, tt.want)
			}
		})
	}
}
