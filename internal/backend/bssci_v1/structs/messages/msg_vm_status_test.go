package messages

import (
	"reflect"
	"testing"
	"time"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

func TestNewVmStatus(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want VmStatus
	}{
		{
			name: "vmStatus",
			args: args{1},
			want: VmStatus{
				Command: structs.MsgVmStatus,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmStatus(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVmStatusFromProto(t *testing.T) {
	type args struct {
		opId int64
		pb   *bs.RequestVariableMacStatus
	}
	tests := []struct {
		name    string
		args    args
		want    *VmStatus
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				opId: 1,
				pb:   &bs.RequestVariableMacStatus{},
			},
			want: &VmStatus{
				Command: structs.MsgVmStatus,
				OpId:    1,
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
			got, err := NewVmStatusFromProto(tt.args.opId, tt.args.pb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVmStatusFromProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmStatusFromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmStatus_GetOpId(t *testing.T) {
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
			name: "vmStatus",
			fields: fields{
				structs.MsgVmStatus,
				1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmStatus{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmStatus.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmStatus_GetCommand(t *testing.T) {
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
			name: "vmStatus",
			fields: fields{
				structs.MsgVmStatus,
				1,
			},
			want: structs.MsgVmStatus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmStatus{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmStatus.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmStatus_SetOpId(t *testing.T) {
	type fields struct {
		Command structs.Command
		OpId    int64
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
			name:   "vmStatus",
			fields: fields{structs.MsgVmStatus, 1},
			args:   args{opId: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmStatus{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			m.SetOpId(tt.args.opId)

			if m.OpId != tt.args.opId {
				t.Errorf("VmStatus.SetOpId() = %v, want %v", m.OpId, tt.args.opId)
			}
		})
	}
}

func TestNewVmStatusRsp(t *testing.T) {
	type args struct {
		opId     int64
		macTypes []int64
	}
	tests := []struct {
		name string
		args args
		want VmStatusRsp
	}{
		{
			name: "vmStatusRsp",
			args: args{1, []int64{10, 20}},
			want: VmStatusRsp{
				Command:  structs.MsgVmStatusRsp,
				OpId:     1,
				MacTypes: []int64{10, 20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmStatusRsp(tt.args.opId, tt.args.macTypes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmStatusRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmStatusRsp_GetOpId(t *testing.T) {
	type fields struct {
		Command  structs.Command
		OpId     int64
		MacTypes []int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "vmStatusRsp",
			fields: fields{
				structs.MsgVmStatusRsp,
				1,
				[]int64{10, 20},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmStatusRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmStatusRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmStatusRsp_GetCommand(t *testing.T) {
	type fields struct {
		Command  structs.Command
		OpId     int64
		MacTypes []int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "vmStatusRsp",
			fields: fields{
				structs.MsgVmStatusRsp,
				1,
				[]int64{10, 20},
			},
			want: structs.MsgVmStatusRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmStatusRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmStatusRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmStatusRsp_IntoProto(t *testing.T) {

	//monkey patch time.now()

	var seconds int64 = 1000000
	var nanos int64 = 123

	fakeNow := time.Unix(seconds, nanos)

	getNow = func() time.Time { return fakeNow }

	testTs := timestamppb.Timestamp{
		Seconds: int64(seconds),
		Nanos:   int32(nanos),
	}

	type fields struct {
		Command  structs.Command
		OpId     int64
		MacTypes []int64
	}
	type args struct {
		bsEui *common.EUI64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *bs.BasestationUplink
	}{
		{
			name: "vmStatusRsp",
			fields: fields{
				structs.MsgVmStatusRsp,
				1,
				[]int64{10, 20},
			},
			args: args{bsEui: &common.EUI64{1}},
			want: &bs.BasestationUplink{
				BsEui: "0100000000000000",
				Ts:    &testTs,
				OpId:  1,
				Message: &bs.BasestationUplink_VmStatus{
					VmStatus: &bs.BasestationVariableMacStatus{
						MacTypes: []uint32{10, 20},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmStatusRsp{
				Command:  tt.fields.Command,
				OpId:     tt.fields.OpId,
				MacTypes: tt.fields.MacTypes,
			}
			if got := m.IntoProto(tt.args.bsEui); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmStatusRsp.IntoProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVmStatusCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want VmStatusCmp
	}{
		{
			name: "vmStatusCmp",
			args: args{1},
			want: VmStatusCmp{
				Command: structs.MsgVmStatusCmp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmStatusCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmStatusCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmStatusCmp_GetOpId(t *testing.T) {
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
			name: "vmStatusCmp",
			fields: fields{
				structs.MsgVmStatusCmp,
				1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmStatusCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmStatusCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmStatusCmp_GetCommand(t *testing.T) {
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
			name: "vmStatusCmp",
			fields: fields{
				structs.MsgVmStatusCmp,
				1,
			},
			want: structs.MsgVmStatusCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmStatusCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmStatusCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
