package messages

import (
	"reflect"
	"testing"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

func TestNewVmDeactivate(t *testing.T) {
	type args struct {
		opId    int64
		macType uint32
	}
	tests := []struct {
		name string
		args args
		want VmDeactivate
	}{
		{
			name: "vmDeactivate",
			args: args{1, 10},
			want: VmDeactivate{
				Command: structs.MsgVmDeactivate,
				OpId:    1,
				MacType: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmDeactivate(tt.args.opId, tt.args.macType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmDeactivate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVmDeactivateFromProto(t *testing.T) {
	type args struct {
		opId int64
		pb   *bs.DisableVariableMac
	}
	tests := []struct {
		name    string
		args    args
		want    *VmDeactivate
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				opId: 1,
				pb: &bs.DisableVariableMac{
					MacType: 10,
				},
			},
			want: &VmDeactivate{
				Command: structs.MsgVmDeactivate,
				OpId:    1,
				MacType: 10,
			},
		},
		{
			name: "nil",
			args: args{
				opId: 1,
				pb:   nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVmDeactivateFromProto(tt.args.opId, tt.args.pb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVmDeactivateFromProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmDeactivateFromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmDeactivate_GetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		MacType uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "vmDeactivate",
			fields: fields{
				structs.MsgVmDeactivate,
				1,
				10,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmDeactivate{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				MacType: tt.fields.MacType,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmDeactivate.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmDeactivate_GetCommand(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		MacType uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "vmDeactivate",
			fields: fields{
				structs.MsgVmDeactivate,
				1,
				10,
			},
			want: structs.MsgVmDeactivate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmDeactivate{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				MacType: tt.fields.MacType,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmDeactivate.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmDeactivate_SetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
		MacType uint32
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
			name:   "vmDeactivate",
			fields: fields{structs.MsgVmDeactivate, 1, 10},
			args:   args{opId: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmDeactivate{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				MacType: tt.fields.MacType,
			}
			m.SetOpId(tt.args.opId)

			if m.OpId != tt.args.opId {
				t.Errorf("VmDeactivate.SetOpId() = %v, want %v", m.OpId, tt.args.opId)
			}
		})
	}
}

func TestNewVmDeactivateRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want VmDeactivateRsp
	}{
		{
			name: "vmDeactivateRsp",
			args: args{1},
			want: VmDeactivateRsp{
				Command: structs.MsgVmDeactivateRsp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmDeactivateRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmDeactivateRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmDeactivateRsp_GetOpId(t *testing.T) {
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
			name: "vmDeactivateRsp",
			fields: fields{
				structs.MsgVmDeactivateRsp,
				1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmDeactivateRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmDeactivateRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmDeactivateRsp_GetCommand(t *testing.T) {
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
			name: "vmDeactivateRsp",
			fields: fields{
				structs.MsgVmDeactivateRsp,
				1,
			},
			want: structs.MsgVmDeactivateRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmDeactivateRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmDeactivateRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVmDeactivateCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want VmDeactivateCmp
	}{
		{
			name: "vmDeactivateCmp",
			args: args{1},
			want: VmDeactivateCmp{
				Command: structs.MsgVmDeactivateCmp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmDeactivateCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmDeactivateCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmDeactivateCmp_GetOpId(t *testing.T) {
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
			name: "vmDeactivateCmp",
			fields: fields{
				structs.MsgVmDeactivateCmp,
				1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmDeactivateCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmDeactivateCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmDeactivateCmp_GetCommand(t *testing.T) {
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
			name: "vmDeactivateCmp",
			fields: fields{
				structs.MsgVmDeactivateCmp,
				1,
			},
			want: structs.MsgVmDeactivateCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmDeactivateCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmDeactivateCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
