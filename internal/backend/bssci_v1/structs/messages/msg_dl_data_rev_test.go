package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
	"reflect"
	"testing"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

func TestNewDlDataRev(t *testing.T) {
	type args struct {
		opId  int64
		epEui common.EUI64
		queId uint64
	}
	tests := []struct {
		name string
		args args
		want DlDataRev
	}{
		{
			name: "dlDataRev",
			args: args{
				opId:  10,
				epEui: common.EUI64{1},
				queId: 20,
			},
			want: DlDataRev{
				Command: structs.MsgDlDataRev,
				OpId:    10,
				EpEui:   common.EUI64{1},
				QueId:   20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlDataRev(tt.args.opId, tt.args.epEui, tt.args.queId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataRev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDlDataRevFromProto(t *testing.T) {
	type args struct {
		opId int64
		pb   *bs.RevokeDownlink
	}
	tests := []struct {
		name    string
		args    args
		want    *DlDataRev
		wantErr bool
	}{
		{
			name: "dlDataRev",
			args: args{
				opId: 10,
				pb: &bs.RevokeDownlink{
					EndnodeEui: "0100000000000000",
					DlQueId:    20,
				},
			},
			want: &DlDataRev{
				Command: structs.MsgDlDataRev,
				OpId:    10,
				EpEui:   common.EUI64{1},
				QueId:   20,
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
			got, err := NewDlDataRevFromProto(tt.args.opId, tt.args.pb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDlDataRevFromProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataRevFromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataRev_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		EpEui   common.EUI64
		QueId   uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "dlDataRev",
			fields: fields{
				Command: structs.MsgDlDataRev,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataRev{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				EpEui:   tt.fields.EpEui,
				QueId:   tt.fields.QueId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlDataRev.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataRev_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		EpEui   common.EUI64
		QueId   uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "dlDataRev",
			fields: fields{
				Command: structs.MsgDlDataRev,
				OpId:    10,
			},
			want: structs.MsgDlDataRev,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataRev{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				EpEui:   tt.fields.EpEui,
				QueId:   tt.fields.QueId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlDataRev.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataRev_SetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		EpEui   common.EUI64
		QueId   uint64
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
			name: "detPrp",
			fields: fields{
				Command: structs.MsgDlDataRev,
				OpId:    1,
				EpEui:   common.EUI64{},
				QueId:   0,
			},
			args: args{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataRev{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				EpEui:   tt.fields.EpEui,
				QueId:   tt.fields.QueId,
			}
			m.SetOpId(tt.args.opId)

			if m.OpId != tt.args.opId {
				t.Errorf("DlDataRev.SetOpId() = %v, want %v", m.OpId, tt.args.opId)
			}
		})
	}
}

func TestNewDlDataRevRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DlDataRevRsp
	}{
		{
			name: "dlDataRevRsp",
			args: args{1},
			want: DlDataRevRsp{
				Command: structs.MsgDlDataRevRsp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlDataRevRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataRevRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataRevRsp_GetOpId(t *testing.T) {
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
			name: "dlDataRevRsp",
			fields: fields{
				Command: structs.MsgDlDataRevRsp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataRevRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlDataRevRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataRevRsp_GetCommand(t *testing.T) {
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
			name: "dlDataRevRsp",
			fields: fields{
				Command: structs.MsgDlDataRevRsp,
				OpId:    10,
			},
			want: structs.MsgDlDataRevRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataRevRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlDataRevRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDlDataRevCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DlDataRevCmp
	}{
		{
			name: "dlDataRevCmp",
			args: args{1},
			want: DlDataRevCmp{
				Command: structs.MsgDlDataRevCmp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlDataRevCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataRevCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataRevCmp_GetOpId(t *testing.T) {
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
			name: "dlDataRevCmp",
			fields: fields{
				Command: structs.MsgDlDataRevCmp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataRevCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlDataRevCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataRevCmp_GetCommand(t *testing.T) {
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
			name: "dlDataRevCmp",
			fields: fields{
				Command: structs.MsgDlDataRevCmp,
				OpId:    10,
			},
			want: structs.MsgDlDataRevCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataRevCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlDataRevCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
