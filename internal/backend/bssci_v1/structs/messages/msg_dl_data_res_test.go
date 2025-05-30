package messages

import (
	"reflect"
	"testing"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewDlDataRes(t *testing.T) {
	type args struct {
		opId      int64
		epEui     common.EUI64
		queId     uint64
		result    dlDataResult
		txTime    *uint64
		packetCnt *uint32
	}
	tests := []struct {
		name string
		args args
		want DlDataRes
	}{
		{
			name: "dlDataRes",
			args: args{
				opId:      10,
				epEui:     common.EUI64{},
				queId:     20,
				result:    dlDataResult_Sent,
				txTime:    new(uint64),
				packetCnt: new(uint32),
			},
			want: DlDataRes{
				Command:   structs.MsgDlDataRes,
				OpId:      10,
				EpEui:     common.EUI64{},
				QueId:     20,
				Result:    dlDataResult_Sent,
				TxTime:    new(uint64),
				PacketCnt: new(uint32),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlDataRes(tt.args.opId, tt.args.epEui, tt.args.queId, tt.args.result, tt.args.txTime, tt.args.packetCnt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataRes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataRes_GetOpId(t *testing.T) {
	type fields struct {
		Command   structs.Command
		OpId      int64
		EpEui     common.EUI64
		QueId     uint64
		Result    dlDataResult
		TxTime    *uint64
		PacketCnt *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "dlDataRes",
			fields: fields{
				Command: structs.MsgDlDataRes,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataRes{
				Command:   tt.fields.Command,
				OpId:      tt.fields.OpId,
				EpEui:     tt.fields.EpEui,
				QueId:     tt.fields.QueId,
				Result:    tt.fields.Result,
				TxTime:    tt.fields.TxTime,
				PacketCnt: tt.fields.PacketCnt,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlDataRes.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataRes_GetCommand(t *testing.T) {
	type fields struct {
		Command   structs.Command
		OpId      int64
		EpEui     common.EUI64
		QueId     uint64
		Result    dlDataResult
		TxTime    *uint64
		PacketCnt *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "dlDataRes",
			fields: fields{
				Command: structs.MsgDlDataRes,
				OpId:    10,
			},
			want: structs.MsgDlDataRes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataRes{
				Command:   tt.fields.Command,
				OpId:      tt.fields.OpId,
				EpEui:     tt.fields.EpEui,
				QueId:     tt.fields.QueId,
				Result:    tt.fields.Result,
				TxTime:    tt.fields.TxTime,
				PacketCnt: tt.fields.PacketCnt,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlDataRes.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataRes_GetEventType(t *testing.T) {
	type fields struct {
		Command   structs.Command
		OpId      int64
		EpEui     common.EUI64
		QueId     uint64
		Result    dlDataResult
		TxTime    *uint64
		PacketCnt *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   events.EventType
	}{
		{
			name: "dlDataRes",
			fields: fields{
				Command: structs.MsgDlDataRes,
				OpId:    10,
				EpEui:   common.EUI64{1, 2, 3, 4, 5, 6, 7, 8},
			},
			want: events.EventTypeEpDl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataRes{
				Command:   tt.fields.Command,
				OpId:      tt.fields.OpId,
				EpEui:     tt.fields.EpEui,
				QueId:     tt.fields.QueId,
				Result:    tt.fields.Result,
				TxTime:    tt.fields.TxTime,
				PacketCnt: tt.fields.PacketCnt,
			}
			if got := m.GetEventType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlDataRes.GetEventType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataRes_IntoProto(t *testing.T) {

	var testTxTime uint64 = 1000000000000005

	testTxTimePb := timestamppb.Timestamp{
		Seconds: int64(1000000),
		Nanos:   int32(5),
	}

	var testPacketCnt uint32 = 10

	type fields struct {
		Command   structs.Command
		OpId      int64
		EpEui     common.EUI64
		QueId     uint64
		Result    dlDataResult
		TxTime    *uint64
		PacketCnt *uint32
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
			name: "dlDataRes_Sent",
			fields: fields{
				Command:   structs.MsgDlDataRes,
				OpId:      10,
				EpEui:     common.EUI64{1},
				QueId:     20,
				Result:    dlDataResult_Sent,
				TxTime:    &testTxTime,
				PacketCnt: &testPacketCnt,
			},
			args: args{
				bsEui: common.EUI64{2},
			},
			want: &bs.BasestationUplink{
				BsEui: "0200000000000000",
				Message: &bs.BasestationUplink_DlRes{
					DlRes: &bs.BasestationDownlinkResult{
						DlQueId:     20,
						EpEui:       "0100000000000000",
						Result:      bs.DownlinkResultEnum_SENT,
						EpPacketCnt: testPacketCnt,
						TxTime:      &testTxTimePb,
					},
				},
			},
		},
		{
			name: "dlDataRes_Expired",
			fields: fields{
				Command: structs.MsgDlDataRes,
				OpId:    10,
				EpEui:   common.EUI64{1},
				QueId:   20,
				Result:  dlDataResult_Expired,
			},
			args: args{
				bsEui: common.EUI64{2},
			},
			want: &bs.BasestationUplink{
				BsEui: "0200000000000000",
				Message: &bs.BasestationUplink_DlRes{
					DlRes: &bs.BasestationDownlinkResult{
						DlQueId: 20,
						EpEui:   "0100000000000000",
						Result:  bs.DownlinkResultEnum_EXPIRED,
					},
				},
			},
		},
		{
			name: "dlDataRes_invalid",
			fields: fields{
				Command: structs.MsgDlDataRes,
				OpId:    10,
				EpEui:   common.EUI64{1},
				QueId:   20,
				Result:  dlDataResult_Invalid,
			},
			args: args{
				bsEui: common.EUI64{2},
			},
			want: &bs.BasestationUplink{
				BsEui: "0200000000000000",
				Message: &bs.BasestationUplink_DlRes{
					DlRes: &bs.BasestationDownlinkResult{
						EpEui:   "0100000000000000",
						DlQueId: 20,
						Result:  bs.DownlinkResultEnum_INVALID,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataRes{
				Command:   tt.fields.Command,
				OpId:      tt.fields.OpId,
				EpEui:     tt.fields.EpEui,
				QueId:     tt.fields.QueId,
				Result:    tt.fields.Result,
				TxTime:    tt.fields.TxTime,
				PacketCnt: tt.fields.PacketCnt,
			}
			if got := m.IntoProto(&tt.args.bsEui); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlDataRes.IntoProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDlDataResRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DlDataResRsp
	}{
		{
			name: "dlDataResRsp",
			args: args{opId: 10},
			want: DlDataResRsp{
				Command: structs.MsgDlDataResRsp,
				OpId:    10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlDataResRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataResRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataResRsp_GetOpId(t *testing.T) {
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
			name: "dlDataResRsp",
			fields: fields{
				Command: structs.MsgDlDataResRsp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataResRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlDataResRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataResRsp_GetCommand(t *testing.T) {
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
			name: "dlDataResRsp",
			fields: fields{
				Command: structs.MsgDlDataResRsp,
				OpId:    10,
			},
			want: structs.MsgDlDataResRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataResRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlDataResRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDlDataResCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DlDataResCmp
	}{
		{
			name: "dlDataResCmp",
			args: args{opId: 10},
			want: DlDataResCmp{
				Command: structs.MsgDlDataResCmp,
				OpId:    10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlDataResCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataResCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataResCmp_GetOpId(t *testing.T) {
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
			name: "dlDataResCmp",
			fields: fields{
				Command: structs.MsgDlDataResCmp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataResCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlDataResCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataResCmp_GetCommand(t *testing.T) {
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
			name: "dlDataResCmp",
			fields: fields{
				Command: structs.MsgDlDataResCmp,
				OpId:    10,
			},
			want: structs.MsgDlDataResCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataResCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlDataResCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
