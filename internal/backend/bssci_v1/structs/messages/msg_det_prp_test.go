package messages

import (
	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
	"reflect"
	"testing"
)

func TestNewDetPrp(t *testing.T) {
	type args struct {
		opId  int64
		epEui common.EUI64
	}
	tests := []struct {
		name string
		args args
		want DetPrp
	}{
		{
			name: "detPrp",
			args: args{1, common.EUI64{1}},
			want: DetPrp{
				Command: structs.MsgDetPrp,
				OpId:    1,
				EpEui:   common.EUI64{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDetPrp(tt.args.opId, tt.args.epEui); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDetPrp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDetPrpFromProto(t *testing.T) {
	type args struct {
		opId int64
		pb   *cmd.DetachPropagate
	}
	tests := []struct {
		name    string
		args    args
		want    *DetPrp
		wantErr bool
	}{
		{
			name: "attRsp",
			args: args{
				opId: 10,
				pb: &cmd.DetachPropagate{
					EndnodeEui: 0x0001020304050607,
				},
			},
			want: &DetPrp{
				Command: structs.MsgDetPrp,
				OpId:    10,
				EpEui:   common.EUI64{7, 6, 5, 4, 3, 2, 1, 0},
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
			got, err := NewDetPrpFromProto(tt.args.opId, tt.args.pb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDetPrpFromProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDetPrpFromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrp_GetOpId(t *testing.T) {
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
		{name: "detPrp", fields: fields{structs.MsgDetPrp, 1, common.EUI64{1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				EpEui:   tt.fields.EpEui,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DetPrp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrp_SetOpId(t *testing.T) {
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
			name:   "detPrp",
			fields: fields{structs.MsgDetPrp, 1, common.EUI64{1}},
			args:   args{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				EpEui:   tt.fields.EpEui,
			}
			m.SetOpId(tt.args.opId)

			if m.OpId != tt.args.opId {
				t.Errorf("DetPrp.SetOpId() = %v, want %v", m.OpId, tt.args.opId)
			}
		})
	}
}

func TestDetPrp_GetCommand(t *testing.T) {
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
		{name: "detPrp", fields: fields{structs.MsgDetPrp, 1, common.EUI64{1}}, want: structs.MsgDetPrp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				EpEui:   tt.fields.EpEui,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetPrp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDetPrpRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DetPrpRsp
	}{
		{
			name: "detPrpRsp", args: args{1},
			want: DetPrpRsp{
				Command: structs.MsgDetPrpRsp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDetPrpRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDetPrpRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrpRsp_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "detPrpRsp", fields: fields{structs.MsgDetPrpRsp, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrpRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DetPrpRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrpRsp_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{name: "detPrpRsp", fields: fields{structs.MsgDetPrpRsp, 1}, want: structs.MsgDetPrpRsp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrpRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetPrpRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDetPrpCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DetPrpCmp
	}{
		{
			name: "detPrpCmp", args: args{1},
			want: DetPrpCmp{
				Command: structs.MsgDetPrpCmp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDetPrpCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDetPrpCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrpCmp_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "detPrpCmp", fields: fields{structs.MsgDetPrpCmp, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrpCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DetPrpCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetPrpCmp_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{name: "detPrpCmp", fields: fields{structs.MsgDetPrpCmp, 1}, want: structs.MsgDetPrpCmp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DetPrpCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetPrpCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
