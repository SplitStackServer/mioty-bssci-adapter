package messages

import (
	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
	"reflect"
	"testing"
)

func TestNewAttPrp(t *testing.T) {
	type args struct {
		opId            int64
		epEui           common.EUI64
		bidi            bool
		nwkSessionKey   [16]byte
		shAddr          uint16
		lastPacketCount uint32
		dualChan        bool
		repetition      bool
		wideCarrOff     bool
		longBlkDist     bool
	}
	tests := []struct {
		name string
		args args
		want AttPrp
	}{
		{
			name: "attPrp",
			args: args{
				opId:            1,
				epEui:           common.EUI64{},
				bidi:            false,
				nwkSessionKey:   [16]byte{},
				shAddr:          0,
				lastPacketCount: 0,
				dualChan:        false,
				repetition:      false,
				wideCarrOff:     false,
				longBlkDist:     false,
			},
			want: AttPrp{
				Command:         structs.MsgAttPrp,
				OpId:            1,
				EpEui:           common.EUI64{},
				Bidi:            false,
				NwkSessionKey:   [16]byte{},
				ShAddr:          0,
				LastPacketCount: 0,
				DualChan:        false,
				Repetition:      false,
				WideCarrOff:     false,
				LongBlkDist:     false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAttPrp(tt.args.opId, tt.args.epEui, tt.args.bidi, tt.args.nwkSessionKey, tt.args.shAddr, tt.args.lastPacketCount, tt.args.dualChan, tt.args.repetition, tt.args.wideCarrOff, tt.args.longBlkDist); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttPrp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAttPrpFromProto(t *testing.T) {
	type args struct {
		opId int64
		pb   *cmd.AttachPropagate
	}
	tests := []struct {
		name    string
		args    args
		want    *AttPrp
		wantErr bool
	}{
		{
			name: "attPrp",
			args: args{
				opId: 10,
				pb: &cmd.AttachPropagate{
					EndnodeEui:    0x0001020304050607,
					ShAddr:        0,
					NwkSessionKey: []byte{3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0},
					LastPacketCnt: 0,
					Bidi:          false,
					DualChannel:   false,
					Repetition:    false,
					WideCarrOff:   false,
					LongBlkDist:   false,
				},
			},
			want: &AttPrp{
				Command:         structs.MsgAttPrp,
				OpId:            10,
				EpEui:           common.EUI64{7, 6, 5, 4, 3, 2, 1, 0},
				Bidi:            false,
				NwkSessionKey:   [16]byte{3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0},
				ShAddr:          0,
				LastPacketCount: 0,
				DualChan:        false,
				Repetition:      false,
				WideCarrOff:     false,
				LongBlkDist:     false,
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
		{
			name: "invalid_NwkSessionKey",
			args: args{
				opId: 10,
				pb: &cmd.AttachPropagate{
					EndnodeEui:    0x0001020304050607,
					ShAddr:        0,
					NwkSessionKey: []byte{},
					LastPacketCnt: 0,
					Bidi:          false,
					DualChannel:   false,
					Repetition:    false,
					WideCarrOff:   false,
					LongBlkDist:   false,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAttPrpFromProto(tt.args.opId, tt.args.pb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttPrpFromProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttPrpFromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttPrp_GetOpId(t *testing.T) {
	type fields struct {
		Command         structs.Command
		OpId            int64
		EpEui           common.EUI64
		Bidi            bool
		NwkSessionKey   [16]byte
		ShAddr          uint16
		LastPacketCount uint32
		DualChan        bool
		Repetition      bool
		WideCarrOff     bool
		LongBlkDist     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "attPrp",
			fields: fields{
				Command:         structs.MsgAttPrp,
				OpId:            1,
				EpEui:           common.EUI64{},
				Bidi:            false,
				NwkSessionKey:   [16]byte{},
				ShAddr:          0,
				LastPacketCount: 0,
				DualChan:        false,
				Repetition:      false,
				WideCarrOff:     false,
				LongBlkDist:     false,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AttPrp{
				Command:         tt.fields.Command,
				OpId:            tt.fields.OpId,
				EpEui:           tt.fields.EpEui,
				Bidi:            tt.fields.Bidi,
				NwkSessionKey:   tt.fields.NwkSessionKey,
				ShAddr:          tt.fields.ShAddr,
				LastPacketCount: tt.fields.LastPacketCount,
				DualChan:        tt.fields.DualChan,
				Repetition:      tt.fields.Repetition,
				WideCarrOff:     tt.fields.WideCarrOff,
				LongBlkDist:     tt.fields.LongBlkDist,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("AttPrp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttPrp_GetCommand(t *testing.T) {
	type fields struct {
		Command         structs.Command
		OpId            int64
		EpEui           common.EUI64
		Bidi            bool
		NwkSessionKey   [16]byte
		ShAddr          uint16
		LastPacketCount uint32
		DualChan        bool
		Repetition      bool
		WideCarrOff     bool
		LongBlkDist     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "attPrp",
			fields: fields{
				Command:         structs.MsgAttPrp,
				OpId:            1,
				EpEui:           common.EUI64{},
				Bidi:            false,
				NwkSessionKey:   [16]byte{},
				ShAddr:          0,
				LastPacketCount: 0,
				DualChan:        false,
				Repetition:      false,
				WideCarrOff:     false,
				LongBlkDist:     false,
			},
			want: structs.MsgAttPrp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AttPrp{
				Command:         tt.fields.Command,
				OpId:            tt.fields.OpId,
				EpEui:           tt.fields.EpEui,
				Bidi:            tt.fields.Bidi,
				NwkSessionKey:   tt.fields.NwkSessionKey,
				ShAddr:          tt.fields.ShAddr,
				LastPacketCount: tt.fields.LastPacketCount,
				DualChan:        tt.fields.DualChan,
				Repetition:      tt.fields.Repetition,
				WideCarrOff:     tt.fields.WideCarrOff,
				LongBlkDist:     tt.fields.LongBlkDist,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttPrp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttPrp_SetOpId(t *testing.T) {
	type fields struct {
		Command         structs.Command
		OpId            int64
		EpEui           common.EUI64
		Bidi            bool
		NwkSessionKey   [16]byte
		ShAddr          uint16
		LastPacketCount uint32
		DualChan        bool
		Repetition      bool
		WideCarrOff     bool
		LongBlkDist     bool
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
				Command:         structs.MsgAttPrp,
				OpId:            1,
				EpEui:           common.EUI64{},
				Bidi:            false,
				NwkSessionKey:   [16]byte{},
				ShAddr:          0,
				LastPacketCount: 0,
				DualChan:        false,
				Repetition:      false,
				WideCarrOff:     false,
				LongBlkDist:     false,
			},
			args: args{opId: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AttPrp{
				Command:         tt.fields.Command,
				OpId:            tt.fields.OpId,
				EpEui:           tt.fields.EpEui,
				Bidi:            tt.fields.Bidi,
				NwkSessionKey:   tt.fields.NwkSessionKey,
				ShAddr:          tt.fields.ShAddr,
				LastPacketCount: tt.fields.LastPacketCount,
				DualChan:        tt.fields.DualChan,
				Repetition:      tt.fields.Repetition,
				WideCarrOff:     tt.fields.WideCarrOff,
				LongBlkDist:     tt.fields.LongBlkDist,
			}
			m.SetOpId(tt.args.opId)

			if m.OpId != tt.args.opId {
				t.Errorf("AttPrp.SetOpId() = %v, want %v", m.OpId, tt.args.opId)
			}
		})
	}
}

func TestNewAttPrpRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want AttPrpRsp
	}{
		{
			name: "attPrpRsp",
			args: args{1},
			want: AttPrpRsp{
				Command: structs.MsgAttPrpRsp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAttPrpRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttPrpRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttPrpRsp_GetOpId(t *testing.T) {
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
			name: "attPrpRsp",
			fields: fields{
				structs.MsgAttPrpRsp,
				1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AttPrpRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("AttPrpRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttPrpRsp_GetCommand(t *testing.T) {
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
			name: "attPrpRsp",
			fields: fields{
				structs.MsgAttPrpRsp,
				1,
			},
			want: structs.MsgAttPrpRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AttPrpRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttPrpRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAttPrpCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want AttPrpCmp
	}{
		{
			name: "attPrpCmp",
			args: args{1},
			want: AttPrpCmp{
				Command: structs.MsgAttPrpCmp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAttPrpCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttPrpCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttPrpCmp_GetOpId(t *testing.T) {
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
			name: "attPrpCmp",
			fields: fields{
				structs.MsgAttPrpCmp,
				1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AttPrpCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("AttPrpCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttPrpCmp_GetCommand(t *testing.T) {
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
			name: "attPrpCmp",
			fields: fields{
				structs.MsgAttPrpCmp,
				1,
			},
			want: structs.MsgAttPrpCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AttPrpCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttPrpCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
