package messages

import (
	"reflect"
	"testing"
	"time"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v5/bs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewDlRxStat(t *testing.T) {
	type args struct {
		opId      int64
		epEui     common.EUI64
		result    string
		rxTime    uint64
		packetCnt uint32
		dlRxSnr   float64
		dlRxRssi  float64
	}
	tests := []struct {
		name string
		args args
		want DlRxStat
	}{
		{
			name: "dlRxStat",
			args: args{
				opId:      10,
				epEui:     common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
				result:    "",
				rxTime:    1,
				packetCnt: 10,
				dlRxSnr:   20,
				dlRxRssi:  30,
			},
			want: DlRxStat{
				Command:   structs.MsgDlRxStat,
				OpId:      10,
				EpEui:     common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
				RxTime:    1,
				PacketCnt: 10,
				DlRxSnr:   20,
				DlRxRssi:  30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlRxStat(tt.args.opId, tt.args.epEui, tt.args.result, tt.args.rxTime, tt.args.packetCnt, tt.args.dlRxSnr, tt.args.dlRxRssi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlRxStat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStat_GetOpId(t *testing.T) {
	type fields struct {
		Command   structs.Command
		OpId      int64
		EpEui     common.EUI64
		RxTime    uint64
		PacketCnt uint32
		DlRxSnr   float64
		DlRxRssi  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "dlRxStat",
			fields: fields{
				Command: structs.MsgDlRxStat,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStat{
				Command:   tt.fields.Command,
				OpId:      tt.fields.OpId,
				EpEui:     tt.fields.EpEui,
				RxTime:    tt.fields.RxTime,
				PacketCnt: tt.fields.PacketCnt,
				DlRxSnr:   tt.fields.DlRxSnr,
				DlRxRssi:  tt.fields.DlRxRssi,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlRxStat.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStat_GetCommand(t *testing.T) {
	type fields struct {
		Command   structs.Command
		OpId      int64
		EpEui     common.EUI64
		RxTime    uint64
		PacketCnt uint32
		DlRxSnr   float64
		DlRxRssi  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "dlRxStat",
			fields: fields{
				Command: structs.MsgDlRxStat,
				OpId:    10,
			},
			want: structs.MsgDlRxStat,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStat{
				Command:   tt.fields.Command,
				OpId:      tt.fields.OpId,
				EpEui:     tt.fields.EpEui,
				RxTime:    tt.fields.RxTime,
				PacketCnt: tt.fields.PacketCnt,
				DlRxSnr:   tt.fields.DlRxSnr,
				DlRxRssi:  tt.fields.DlRxRssi,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlRxStat.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStat_GetEventType(t *testing.T) {
	type fields struct {
		Command   structs.Command
		OpId      int64
		EpEui     common.EUI64
		RxTime    uint64
		PacketCnt uint32
		DlRxSnr   float64
		DlRxRssi  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   events.EventType
	}{
		{
			name: "dlRxStat",
			fields: fields{
				Command:   structs.MsgDlRxStat,
				OpId:      10,
				EpEui:     common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
				RxTime:    1,
				PacketCnt: 10,
				DlRxSnr:   0,
				DlRxRssi:  0,
			},
			want: events.EventTypeEpRx,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStat{
				Command:   tt.fields.Command,
				OpId:      tt.fields.OpId,
				EpEui:     tt.fields.EpEui,
				RxTime:    tt.fields.RxTime,
				PacketCnt: tt.fields.PacketCnt,
				DlRxSnr:   tt.fields.DlRxSnr,
				DlRxRssi:  tt.fields.DlRxRssi,
			}
			if got := m.GetEventType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlRxStat.GetEventType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStat_IntoProto(t *testing.T) {

	var testRxTime uint64 = 1000000000000005

	testRxTimePb := timestamppb.Timestamp{
		Seconds: int64(1000000),
		Nanos:   int32(5),
	}

	var seconds int64 = 1000000
	var nanos int64 = 123

	fakeNow := time.Unix(seconds, nanos)

	getNow = func() time.Time { return fakeNow }

	testTs := timestamppb.Timestamp{
		Seconds: int64(seconds),
		Nanos:   int32(nanos),
	}

	type fields struct {
		Command   structs.Command
		OpId      int64
		EpEui     common.EUI64
		RxTime    uint64
		PacketCnt uint32
		DlRxSnr   float64
		DlRxRssi  float64
	}
	type args struct {
		bsEui common.EUI64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *bs.BasestationUplink
	}{
		{
			name: "dlRxStat",
			fields: fields{
				Command:   structs.MsgDlRxStat,
				OpId:      10,
				EpEui:     common.EUI64{1},
				RxTime:    testRxTime,
				PacketCnt: 10,
				DlRxSnr:   2.5,
				DlRxRssi:  -100.5,
			},
			args: args{bsEui: common.EUI64{2}},
			want: &bs.BasestationUplink{
				BsEui: "0200000000000000",
				Ts:    &testTs,
				OpId:  10,
				Message: &bs.BasestationUplink_DlRxStat{
					DlRxStat: &bs.BasestationDownlinkRxStatus{
						EpEui:     "0100000000000000",
						RxTime:    &testRxTimePb,
						PacketCnt: 10,
						DlRxRssi:  -100.5,
						DlRxSnr:   2.5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStat{
				Command:   tt.fields.Command,
				OpId:      tt.fields.OpId,
				EpEui:     tt.fields.EpEui,
				RxTime:    tt.fields.RxTime,
				PacketCnt: tt.fields.PacketCnt,
				DlRxSnr:   tt.fields.DlRxSnr,
				DlRxRssi:  tt.fields.DlRxRssi,
			}
			if got := m.IntoProto(&tt.args.bsEui); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlRxStat.IntoProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDlRxStatRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DlRxStatRsp
	}{
		{
			name: "dlRxStatCmp",
			args: args{opId: 10},
			want: DlRxStatRsp{
				Command: structs.MsgDlRxStatRsp,
				OpId:    10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlRxStatRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlRxStatRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStatRsp_GetOpId(t *testing.T) {
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
			name: "dlRxStatRsp",
			fields: fields{
				Command: structs.MsgDlRxStatRsp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStatRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlRxStatRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStatRsp_GetCommand(t *testing.T) {
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
			name: "dlRxStatRsp",
			fields: fields{
				Command: structs.MsgDlRxStatRsp,
				OpId:    10,
			},
			want: structs.MsgDlRxStatRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStatRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlRxStatRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDlRxStatCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DlRxStatCmp
	}{
		{
			name: "dlRxStatCmp",
			args: args{opId: 10},
			want: DlRxStatCmp{
				Command: structs.MsgDlRxStatCmp,
				OpId:    10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlRxStatCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlRxStatCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStatCmp_GetOpId(t *testing.T) {
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
			name: "dlRxStatCmp",
			fields: fields{
				Command: structs.MsgDlRxStatCmp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStatCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlRxStatCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStatCmp_GetCommand(t *testing.T) {
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
			name: "dlRxStatCmp",
			fields: fields{
				Command: structs.MsgDlRxStatCmp,
				OpId:    10,
			},
			want: structs.MsgDlRxStatCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStatCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlRxStatCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
