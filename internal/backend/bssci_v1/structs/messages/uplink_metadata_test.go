package messages

import (
	"reflect"
	"testing"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestUplinkMetadata_IntoProto(t *testing.T) {

	var testRxTime uint64 = 1000000000000005

	testRxTimePb := timestamppb.Timestamp{
		Seconds: int64(1000000),
		Nanos:   int32(5),
	}

	var testRxDuration uint64 = 1000001005

	testRxDurationPb := durationpb.Duration{
		Seconds: int64(1),
		Nanos:   int32(1005),
	}

	testProfile := "testProfile"

	testEqSnr := 3.0

	type fields struct {
		OpId       int64
		RxTime     uint64
		RxDuration *uint64
		PacketCnt  uint32
		Profile    *string
		SNR        float64
		RSSI       float64
		EqSnr      *float64
		Subpackets *Subpackets
	}
	tests := []struct {
		name   string
		fields fields
		want   *bs.EndnodeUplinkMetadata
	}{
		{name: "metadata1", fields: fields{
			OpId:       10,
			RxTime:     testRxTime,
			RxDuration: &testRxDuration,
			PacketCnt:  10,
			Profile:    &testProfile,
			SNR:        1.0,
			RSSI:       2.0,
			EqSnr:      &testEqSnr,
			Subpackets: nil,
		}, want: &bs.EndnodeUplinkMetadata{
			OpId:          10,
			RxTime:        &testRxTimePb,
			RxDuration:    &testRxDurationPb,
			PacketCnt:     10,
			Profile:       &testProfile,
			Rssi:          2.0,
			Snr:           1.0,
			EqSnr:         &testEqSnr,
			SubpacketInfo: nil,
		}},

		{name: "metadata2", fields: fields{
			OpId:       10,
			RxTime:     testRxTime,
			RxDuration: &testRxDuration,
			PacketCnt:  10,
			Profile:    &testProfile,
			SNR:        1.0,
			RSSI:       2.0,
			EqSnr:      &testEqSnr,
			Subpackets: &Subpackets{},
		}, want: &bs.EndnodeUplinkMetadata{
			OpId:          10,
			RxTime:        &testRxTimePb,
			RxDuration:    &testRxDurationPb,
			PacketCnt:     10,
			Profile:       &testProfile,
			Rssi:          2.0,
			Snr:           1.0,
			EqSnr:         &testEqSnr,
			SubpacketInfo: []*bs.EndnodeUplinkSubpacket{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UplinkMetadata{
				OpId:       tt.fields.OpId,
				RxTime:     tt.fields.RxTime,
				RxDuration: tt.fields.RxDuration,
				PacketCnt:  tt.fields.PacketCnt,
				Profile:    tt.fields.Profile,
				SNR:        tt.fields.SNR,
				RSSI:       tt.fields.RSSI,
				EqSnr:      tt.fields.EqSnr,
				Subpackets: tt.fields.Subpackets,
			}
			if got := m.IntoProto(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UplinkMetadata.IntoProto() = %v,\n want %v", got, tt.want)
			}
		})
	}
}
