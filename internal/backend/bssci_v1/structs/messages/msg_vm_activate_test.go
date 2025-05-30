package messages

import (
	"reflect"
	"testing"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

func TestNewVmActivate(t *testing.T) {
	type args struct {
		opId    int64
		macType uint32
	}
	tests := []struct {
		name string
		args args
		want VmActivate
	}{
		{
			name: "vmActivate",
			args: args{1, 10},
			want: VmActivate{
				Command: structs.MsgVmActivate,
				OpId:    1,
				MacType: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmActivate(tt.args.opId, tt.args.macType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmActivate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVmActivateFromProto(t *testing.T) {
	type args struct {
		opId int64
		pb   *bs.EnableVariableMac
	}
	tests := []struct {
		name    string
		args    args
		want    *VmActivate
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				opId: 1,
				pb: &bs.EnableVariableMac{
					MacType: 10,
				},
			},
			want: &VmActivate{
				Command: structs.MsgVmActivate,
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
			got, err := NewVmActivateFromProto(tt.args.opId, tt.args.pb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVmActivateFromProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmActivateFromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmActivate_GetOpId(t *testing.T) {
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
			name: "vmActivate",
			fields: fields{
				structs.MsgVmActivate,
				1,
				10,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmActivate{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				MacType: tt.fields.MacType,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmActivate.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmActivate_GetCommand(t *testing.T) {
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
			name: "vmActivate",
			fields: fields{
				structs.MsgVmActivate,
				1,
				10,
			},
			want: structs.MsgVmActivate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmActivate{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				MacType: tt.fields.MacType,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmActivate.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmActivate_SetOpId(t *testing.T) {
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
			name:   "vmActivate",
			fields: fields{structs.MsgVmActivate, 1, 10},
			args:   args{opId: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmActivate{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
				MacType: tt.fields.MacType,
			}
			m.SetOpId(tt.args.opId)

			if m.OpId != tt.args.opId {
				t.Errorf("VmActivate.SetOpId() = %v, want %v", m.OpId, tt.args.opId)
			}
		})
	}
}

func TestNewVmActivateRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want VmActivateRsp
	}{
		{
			name: "vmActivateRsp",
			args: args{1},
			want: VmActivateRsp{
				Command: structs.MsgVmActivateRsp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmActivateRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmActivateRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmActivateRsp_GetOpId(t *testing.T) {
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
			name: "vmActivateRsp",
			fields: fields{
				structs.MsgVmActivateRsp,
				1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmActivateRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmActivateRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmActivateRsp_GetCommand(t *testing.T) {
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
			name: "vmActivateRsp",
			fields: fields{
				structs.MsgVmActivateRsp,
				1,
			},
			want: structs.MsgVmActivateRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmActivateRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmActivateRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVmActivateCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want VmActivateCmp
	}{
		{
			name: "vmActivateCmp",
			args: args{1},
			want: VmActivateCmp{
				Command: structs.MsgVmActivateCmp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmActivateCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmActivateCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmActivateCmp_GetOpId(t *testing.T) {
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
			name: "vmActivateCmp",
			fields: fields{
				structs.MsgVmActivateCmp,
				1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmActivateCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmActivateCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmActivateCmp_GetCommand(t *testing.T) {
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
			name: "vmActivateCmp",
			fields: fields{
				structs.MsgVmActivateCmp,
				1,
			},
			want: structs.MsgVmActivateCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmActivateCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmActivateCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
