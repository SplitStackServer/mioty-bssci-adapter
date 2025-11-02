package messages

import (
	"reflect"
	"testing"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v5/bs"
)

func TestNewDlRxStatQry(t *testing.T) {
	type args struct {
		opId  int64
		epEui common.EUI64
	}
	tests := []struct {
		name string
		args args
		want DlRxStatQry
	}{
		{
			name: "dlRxStatQry",
			args: args{opId: 10, epEui: common.EUI64{1}},
			want: DlRxStatQry{
				Command: structs.MsgDlRxStatQry,
				OpId:    10,
				EpEui:   common.EUI64{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlRxStatQry(tt.args.opId, tt.args.epEui); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlRxStatQry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDlRxStatQryFromProto(t *testing.T) {
	type args struct {
		opId int64
		pb   *bs.DownlinkRxStatusQuery
	}
	tests := []struct {
		name    string
		args    args
		want    *DlRxStatQry
		wantErr bool
	}{

		{
			name: "dlRxStatQry",
			args: args{
				opId: 10,
				pb: &bs.DownlinkRxStatusQuery{
					EndnodeEui: "0100000000000000",
				},
			},
			want: &DlRxStatQry{
				Command: structs.MsgDlRxStatQry,
				OpId:    10,
				EpEui:   common.EUI64{1},
			},
			wantErr: false,
		},
		{
			name: "nil",
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
			got, err := NewDlRxStatQryFromProto(tt.args.opId, tt.args.pb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDlRxStatQryFromProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlRxStatQryFromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStatQry_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		EpEui   common.EUI64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "dlRxStatQry",
			fields: fields{
				Command: structs.MsgDlRxStatQry,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStatQry{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				EpEui:   tt.fields.EpEui,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlRxStatQry.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStatQry_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		EpEui   common.EUI64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "dlRxStatQry",
			fields: fields{
				Command: structs.MsgDlRxStatQry,
				OpId:    10,
			},
			want: structs.MsgDlRxStatQry,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStatQry{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				EpEui:   tt.fields.EpEui,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlRxStatQry.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStatQry_SetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		EpEui   common.EUI64
	}
	type args struct {
		opId int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "attPrp",
			fields: fields{
				Command: structs.MsgAttPrp,
				OpId:    1,
				EpEui:   common.EUI64{1},
			},
			args: args{opId: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStatQry{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				EpEui:   tt.fields.EpEui,
			}
			m.SetOpId(tt.args.opId)

			if m.OpId != tt.args.opId {
				t.Errorf("DlRxStatQry.SetOpId() = %v, want %v", m.OpId, tt.args.opId)
			}
		})
	}
}

func TestNewDlRxStatQryRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DlRxStatQryRsp
	}{
		{
			name: "dlRxStatQryRsp",
			args: args{opId: 10},
			want: DlRxStatQryRsp{
				Command: structs.MsgDlRxStatQryRsp,
				OpId:    10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlRxStatQryRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlRxStatQryRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStatQryRsp_GetOpId(t *testing.T) {
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
			name: "dlRxStatQryRsp",
			fields: fields{
				Command: structs.MsgDlRxStatQryRsp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStatQryRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlRxStatQryRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStatQryRsp_GetCommand(t *testing.T) {
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
			name: "dlRxStatQryRsp",
			fields: fields{
				Command: structs.MsgDlRxStatQryRsp,
				OpId:    10,
			},
			want: structs.MsgDlRxStatQryRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStatQryRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlRxStatQryRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDlRxStatQryCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DlRxStatQryCmp
	}{
		{
			name: "dlRxStatQryCmp",
			args: args{opId: 10},
			want: DlRxStatQryCmp{
				Command: structs.MsgDlRxStatQryCmp,
				OpId:    10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlRxStatQryCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlRxStatQryCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStatQryCmp_GetOpId(t *testing.T) {
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
			name: "dlRxStatQryCmp",
			fields: fields{
				Command: structs.MsgDlRxStatQryCmp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStatQryCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlRxStatQryCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlRxStatQryCmp_GetCommand(t *testing.T) {
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
			name: "dlRxStatQryCmp",
			fields: fields{
				Command: structs.MsgDlRxStatQryCmp,
				OpId:    10,
			},
			want: structs.MsgDlRxStatQryCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlRxStatQryCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlRxStatQryCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
