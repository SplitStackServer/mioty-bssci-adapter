package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/backend/events"
	"mioty-bssci-adapter/internal/common"
	"reflect"
	"testing"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewDet(t *testing.T) {
	type args struct {
		opId       int64
		epEui      common.EUI64
		rxTime     uint64
		rxDuration *uint64
		packetCnt  uint32
		snr        float64
		rssi       float64
		eqSnr      *float64
		profile    *string
		subpackets *Subpackets
		sign       [4]byte
	}
	tests := []struct {
		name string
		args args
		want Det
	}{
		{
			name: "det",
			args: args{
				opId:       0,
				epEui:      common.EUI64{},
				rxTime:     1,
				rxDuration: new(uint64),
				packetCnt:  2,
				snr:        3.0,
				rssi:       -100.0,
				eqSnr:      nil,
				profile:    new(string),
				subpackets: &Subpackets{},
				sign:       [4]byte{},
			},
			want: Det{
				Command:    structs.MsgDet,
				OpId:       0,
				EpEui:      common.EUI64{},
				RxTime:     1,
				RxDuration: new(uint64),
				PacketCnt:  2,
				SNR:        3.0,
				RSSI:       -100.0,
				EqSnr:      nil,
				Profile:    new(string),
				Subpackets: &Subpackets{},
				Sign:       [4]byte{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDet(tt.args.opId, tt.args.epEui, tt.args.rxTime, tt.args.rxDuration, tt.args.packetCnt, tt.args.snr, tt.args.rssi, tt.args.eqSnr, tt.args.profile, tt.args.subpackets, tt.args.sign); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDet_GetOpId(t *testing.T) {
	type fields struct {
		Command    structs.Command
		OpId       int64
		EpEui      common.EUI64
		RxTime     uint64
		RxDuration *uint64
		PacketCnt  uint32
		SNR        float64
		RSSI       float64
		EqSnr      *float64
		Profile    *string
		Subpackets *Subpackets
		Sign       [4]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "det",
			fields: fields{
				Command:    structs.MsgDet,
				OpId:       10,
				EpEui:      common.EUI64{},
				RxTime:     1,
				RxDuration: new(uint64),
				PacketCnt:  2,
				SNR:        3.0,
				RSSI:       -100.0,
				EqSnr:      nil,
				Profile:    new(string),
				Subpackets: &Subpackets{},
				Sign:       [4]byte{},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Det{
				Command:    tt.fields.Command,
				OpId:       tt.fields.OpId,
				EpEui:      tt.fields.EpEui,
				RxTime:     tt.fields.RxTime,
				RxDuration: tt.fields.RxDuration,
				PacketCnt:  tt.fields.PacketCnt,
				SNR:        tt.fields.SNR,
				RSSI:       tt.fields.RSSI,
				EqSnr:      tt.fields.EqSnr,
				Profile:    tt.fields.Profile,
				Subpackets: tt.fields.Subpackets,
				Sign:       tt.fields.Sign,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("Det.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDet_GetCommand(t *testing.T) {
	type fields struct {
		Command    structs.Command
		OpId       int64
		EpEui      common.EUI64
		RxTime     uint64
		RxDuration *uint64
		PacketCnt  uint32
		SNR        float64
		RSSI       float64
		EqSnr      *float64
		Profile    *string
		Subpackets *Subpackets
		Sign       [4]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "det",
			fields: fields{
				Command:    structs.MsgDet,
				OpId:       10,
				EpEui:      common.EUI64{},
				RxTime:     1,
				RxDuration: new(uint64),
				PacketCnt:  2,
				SNR:        3.0,
				RSSI:       -100.0,
				EqSnr:      nil,
				Profile:    new(string),
				Subpackets: &Subpackets{},
				Sign:       [4]byte{},
			},
			want: structs.MsgDet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Det{
				Command:    tt.fields.Command,
				OpId:       tt.fields.OpId,
				EpEui:      tt.fields.EpEui,
				RxTime:     tt.fields.RxTime,
				RxDuration: tt.fields.RxDuration,
				PacketCnt:  tt.fields.PacketCnt,
				SNR:        tt.fields.SNR,
				RSSI:       tt.fields.RSSI,
				EqSnr:      tt.fields.EqSnr,
				Profile:    tt.fields.Profile,
				Subpackets: tt.fields.Subpackets,
				Sign:       tt.fields.Sign,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Det.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDet_GetEventType(t *testing.T) {
	type fields struct {
		Command    structs.Command
		OpId       int64
		EpEui      common.EUI64
		RxTime     uint64
		RxDuration *uint64
		PacketCnt  uint32
		SNR        float64
		RSSI       float64
		EqSnr      *float64
		Profile    *string
		Subpackets *Subpackets
		Sign       [4]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   events.EventType
	}{
		{
			name: "det",
			fields: fields{
				Command:    structs.MsgDet,
				OpId:       10,
				EpEui:      common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
				RxTime:     1,
				RxDuration: new(uint64),
				PacketCnt:  2,
				SNR:        3.0,
				RSSI:       -100.0,
				EqSnr:      nil,
				Profile:    new(string),
				Subpackets: &Subpackets{},
				Sign:       [4]byte{},
			},
			want: events.EventTypeEpOtaa,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Det{
				Command:    tt.fields.Command,
				OpId:       tt.fields.OpId,
				EpEui:      tt.fields.EpEui,
				RxTime:     tt.fields.RxTime,
				RxDuration: tt.fields.RxDuration,
				PacketCnt:  tt.fields.PacketCnt,
				SNR:        tt.fields.SNR,
				RSSI:       tt.fields.RSSI,
				EqSnr:      tt.fields.EqSnr,
				Profile:    tt.fields.Profile,
				Subpackets: tt.fields.Subpackets,
				Sign:       tt.fields.Sign,
			}
			if got := m.GetEventType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Det.GetEventType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDet_IntoProto(t *testing.T) {
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

	type fields struct {
		Command    structs.Command
		OpId       int64
		EpEui      common.EUI64
		RxTime     uint64
		RxDuration *uint64
		PacketCnt  uint32
		SNR        float64
		RSSI       float64
		EqSnr      *float64
		Profile    *string
		Subpackets *Subpackets
		Sign       [4]byte
	}
	type args struct {
		bsEui common.EUI64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *bs.ProtoEndnodeMessage
	}{
		{
			name: "det",
			fields: fields{
				Command:    structs.MsgDet,
				OpId:       10,
				EpEui:      common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
				RxTime:     testRxTime,
				RxDuration: &testRxDuration,
				PacketCnt:  2,
				SNR:        3.0,
				RSSI:       -100.0,
				EqSnr:      nil,
				Profile:    new(string),
				Subpackets: &Subpackets{},
				Sign:       [4]byte{3, 2, 1, 0},
			},
			args: args{bsEui: common.EUI64{1}},
			want: &bs.ProtoEndnodeMessage{
				BsEui:      "0100000000000000",
				EndnodeEui: "0001020304050607",
				V1: &bs.ProtoEndnodeMessage_Det{
					Det: &bs.EndnodeDetMessage{
						OpId: 10,
						Sign: 0x00010203,
						Meta: &bs.EndnodeUplinkMetadata{
							RxTime:        &testRxTimePb,
							RxDuration:    &testRxDurationPb,
							PacketCnt:     2,
							Profile:       new(string),
							Rssi:          -100.0,
							Snr:           3.0,
							EqSnr:         nil,
							SubpacketInfo: []*bs.EndnodeUplinkSubpacket{},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Det{
				Command:    tt.fields.Command,
				OpId:       tt.fields.OpId,
				EpEui:      tt.fields.EpEui,
				RxTime:     tt.fields.RxTime,
				RxDuration: tt.fields.RxDuration,
				PacketCnt:  tt.fields.PacketCnt,
				SNR:        tt.fields.SNR,
				RSSI:       tt.fields.RSSI,
				EqSnr:      tt.fields.EqSnr,
				Profile:    tt.fields.Profile,
				Subpackets: tt.fields.Subpackets,
				Sign:       tt.fields.Sign,
			}
			if got := m.IntoProto(tt.args.bsEui); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Det.IntoProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDetRsp(t *testing.T) {
	type args struct {
		opId int64
		sign [4]byte
	}
	tests := []struct {
		name string
		args args
		want DetRsp
	}{
		{
			name: "detRsp",
			args: args{
				opId: 10,
				sign: [4]byte{},
			},
			want: DetRsp{
				Command: structs.MsgDetRsp,
				OpId:    10,
				Sign:    [4]byte{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDetRsp(tt.args.opId, tt.args.sign); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDetRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDetRspFromProto(t *testing.T) {
	type args struct {
		opId int64
		pb   *bs.EndnodeDetachResponse
	}
	tests := []struct {
		name    string
		args    args
		want    *DetRsp
		wantErr bool
	}{
		{
			name: "detRsp",
			args: args{
				opId: 10,
				pb: &bs.EndnodeDetachResponse{
					Sign: 0x00010203,
				},
			},
			want: &DetRsp{
				Command: structs.MsgDetRsp,
				OpId:    10,
				Sign:    [4]byte{3, 2, 1, 0},
			},
			wantErr: false,
		},
		{
			name: "detRspNil",
			args: args{
				opId: 10,
				pb:   nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDetRspFromProto(tt.args.opId, tt.args.pb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDetRspFromProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDetRspFromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetRsp_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		Sign    [4]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "detRsp",
			fields: fields{
				Command: structs.MsgDetRsp,
				OpId:    10,
				Sign:    [4]byte{},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				Sign:    tt.fields.Sign,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DetRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetRsp_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		Sign    [4]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "detRsp",
			fields: fields{
				Command: structs.MsgDetRsp,
				OpId:    10,
				Sign:    [4]byte{},
			},
			want: structs.MsgDetRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				Sign:    tt.fields.Sign,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDetCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DetCmp
	}{
		{
			name: "detCmp",
			args: args{
				opId: 10,
			},
			want: DetCmp{
				Command: structs.MsgDetCmp,
				OpId:    10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDetCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDetCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetCmp_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "detCmp",
			fields: fields{
				Command: structs.MsgDetCmp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DetCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetCmp_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "detCmp",
			fields: fields{
				Command: structs.MsgDetCmp,
				OpId:    10,
			},
			want: structs.MsgDetCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
