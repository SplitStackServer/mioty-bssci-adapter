package messages

import (
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
	"reflect"
	"testing"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

func TestNewDlDataQue(t *testing.T) {
	type args struct {
		opId         int64
		epEui        common.EUI64
		queId        uint64
		prio         *float32
		format       *byte
		userData     []byte
		responseExp  *bool
		responsePrio *bool
		dlWindReq    *bool
		expOnly      *bool
	}
	tests := []struct {
		name string
		args args
		want DlDataQue
	}{
		{
			name: "dlDataQue",
			args: args{
				opId:         0,
				epEui:        common.EUI64{},
				queId:        0,
				prio:         new(float32),
				format:       new(byte),
				userData:     []byte{0, 1, 2, 3},
				responseExp:  new(bool),
				responsePrio: new(bool),
				dlWindReq:    new(bool),
				expOnly:      new(bool),
			},
			want: DlDataQue{
				Command:      structs.MsgDlDataQue,
				OpId:         0,
				EpEui:        common.EUI64{},
				QueId:        0,
				CntDepend:    false,
				PacketCnt:    nil,
				UserData:     [][]byte{{0, 1, 2, 3}},
				Format:       new(byte),
				Prio:         new(float32),
				ResponseExp:  new(bool),
				ResponsePrio: new(bool),
				DlWindReq:    new(bool),
				ExpOnly:      new(bool),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlDataQue(tt.args.opId, tt.args.epEui, tt.args.queId, tt.args.prio, tt.args.format, tt.args.userData, tt.args.responseExp, tt.args.responsePrio, tt.args.dlWindReq, tt.args.expOnly); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataQue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDlDataQueEnc(t *testing.T) {
	type args struct {
		opId         int64
		epEui        common.EUI64
		queId        uint64
		prio         *float32
		format       *byte
		packetCnt    []uint32
		userData     [][]byte
		responseExp  *bool
		responsePrio *bool
		dlWindReq    *bool
		expOnly      *bool
	}
	tests := []struct {
		name string
		args args
		want DlDataQue
	}{
		{
			name: "dlDataQue",
			args: args{
				opId:         0,
				epEui:        common.EUI64{},
				queId:        0,
				prio:         new(float32),
				format:       new(byte),
				userData:     [][]byte{{0, 1, 2, 3}, {0, 1, 2, 3}},
				responseExp:  new(bool),
				responsePrio: new(bool),
				dlWindReq:    new(bool),
				expOnly:      new(bool),
				packetCnt:    []uint32{1, 2},
			},
			want: DlDataQue{
				Command:      structs.MsgDlDataQue,
				OpId:         0,
				EpEui:        common.EUI64{},
				QueId:        0,
				CntDepend:    true,
				PacketCnt:    &[]uint32{1, 2},
				UserData:     [][]byte{{0, 1, 2, 3}, {0, 1, 2, 3}},
				Format:       new(byte),
				Prio:         new(float32),
				ResponseExp:  new(bool),
				ResponsePrio: new(bool),
				DlWindReq:    new(bool),
				ExpOnly:      new(bool),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlDataQueEnc(tt.args.opId, tt.args.epEui, tt.args.queId, tt.args.prio, tt.args.format, tt.args.packetCnt, tt.args.userData, tt.args.responseExp, tt.args.responsePrio, tt.args.dlWindReq, tt.args.expOnly); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataQueEnc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDlDataQueFromProto(t *testing.T) {
	type args struct {
		opId int64
		pb   *bs.EnqueDownlink
	}
	tests := []struct {
		name    string
		args    args
		want    *DlDataQue
		wantErr bool
	}{
		{
			name: "dlDataQue_Ack",
			args: args{
				opId: 10,
				pb: &bs.EnqueDownlink{
					EndnodeEui:     "0706050403020100",
					DlQueId:        20,
					Priority:       new(float32),
					Format:         new(uint32),
					Payload:        &bs.EnqueDownlink_Ack{Ack: &bs.Acknowledgement{}},
					ResponseExp:    new(bool),
					ResponsePrio:   new(bool),
					ReqDlWindow:    new(bool),
					OnlyIfExpected: new(bool),
				},
			},
			want: &DlDataQue{
				Command:      structs.MsgDlDataQue,
				OpId:         10,
				EpEui:        common.EUI64{7, 6, 5, 4, 3, 2, 1, 0},
				QueId:        20,
				CntDepend:    false,
				PacketCnt:    nil,
				UserData:     nil,
				Format:       new(byte),
				Prio:         new(float32),
				ResponseExp:  new(bool),
				ResponsePrio: new(bool),
				DlWindReq:    new(bool),
				ExpOnly:      new(bool),
			},
			wantErr: false,
		},
		{
			name: "dlDataQue_Data",
			args: args{
				opId: 10,
				pb: &bs.EnqueDownlink{
					EndnodeEui:     "0706050403020100",
					DlQueId:        20,
					Priority:       new(float32),
					Format:         new(uint32),
					Payload:        &bs.EnqueDownlink_Data{Data: &bs.DownlinkData{Data: []byte{0, 1, 2, 3}}},
					ResponseExp:    new(bool),
					ResponsePrio:   new(bool),
					ReqDlWindow:    new(bool),
					OnlyIfExpected: new(bool),
				},
			},
			want: &DlDataQue{
				Command:      structs.MsgDlDataQue,
				OpId:         10,
				EpEui:        common.EUI64{7, 6, 5, 4, 3, 2, 1, 0},
				QueId:        20,
				CntDepend:    false,
				PacketCnt:    nil,
				UserData:     [][]byte{{0, 1, 2, 3}},
				Format:       new(byte),
				Prio:         new(float32),
				ResponseExp:  new(bool),
				ResponsePrio: new(bool),
				DlWindReq:    new(bool),
				ExpOnly:      new(bool),
			},
			wantErr: false,
		},
		{
			name: "dlDataQue_Data_Err",
			args: args{
				opId: 10,
				pb: &bs.EnqueDownlink{
					EndnodeEui:     "0706050403020100",
					DlQueId:        20,
					Priority:       new(float32),
					Format:         new(uint32),
					Payload:        &bs.EnqueDownlink_Data{Data: &bs.DownlinkData{Data: []byte{}}},
					ResponseExp:    new(bool),
					ResponsePrio:   new(bool),
					ReqDlWindow:    new(bool),
					OnlyIfExpected: new(bool),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "dlDataQue_DataEnc",
			args: args{
				opId: 10,
				pb: &bs.EnqueDownlink{
					EndnodeEui: "0706050403020100",
					DlQueId:    20,
					Priority:   new(float32),
					Format:     new(uint32),

					Payload: &bs.EnqueDownlink_DataEnc{DataEnc: &bs.DownlinkDataEncrypted{
						Data:      [][]byte{{0, 1, 2, 3}, {0, 1, 2, 3}},
						PacketCnt: []uint32{1, 2},
					},
					},
					ResponseExp:    new(bool),
					ResponsePrio:   new(bool),
					ReqDlWindow:    new(bool),
					OnlyIfExpected: new(bool),
				},
			},
			want: &DlDataQue{
				Command:      structs.MsgDlDataQue,
				OpId:         10,
				EpEui:        common.EUI64{7, 6, 5, 4, 3, 2, 1, 0},
				QueId:        20,
				CntDepend:    true,
				PacketCnt:    &[]uint32{1, 2},
				UserData:     [][]byte{{0, 1, 2, 3}, {0, 1, 2, 3}},
				Format:       new(byte),
				Prio:         new(float32),
				ResponseExp:  new(bool),
				ResponsePrio: new(bool),
				DlWindReq:    new(bool),
				ExpOnly:      new(bool),
			},
			wantErr: false,
		},
		{
			name: "dlDataQue_DataEnc_Err",
			args: args{
				opId: 10,
				pb: &bs.EnqueDownlink{
					EndnodeEui:     "0706050403020100",
					DlQueId:        20,
					Priority:       new(float32),
					Format:         new(uint32),
					Payload:        &bs.EnqueDownlink_DataEnc{DataEnc: &bs.DownlinkDataEncrypted{}},
					ResponseExp:    new(bool),
					ResponsePrio:   new(bool),
					ReqDlWindow:    new(bool),
					OnlyIfExpected: new(bool),
				},
			},
			want:    nil,
			wantErr: true,
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
			name: "invalid",
			args: args{
				opId: 10,
				pb: &bs.EnqueDownlink{
					EndnodeEui:     "0706050403020100",
					DlQueId:        20,
					Priority:       new(float32),
					Format:         new(uint32),
					Payload:        nil,
					ResponseExp:    new(bool),
					ResponsePrio:   new(bool),
					ReqDlWindow:    new(bool),
					OnlyIfExpected: new(bool),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDlDataQueFromProto(tt.args.opId, tt.args.pb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDlDataQueFromProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataQueFromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataQue_GetOpId(t *testing.T) {
	type fields struct {
		Command      structs.Command
		OpId         int64
		EpEui        common.EUI64
		QueId        uint64
		CntDepend    bool
		PacketCnt    *[]uint32
		UserData     [][]byte
		Format       *byte
		Prio         *float32
		ResponseExp  *bool
		ResponsePrio *bool
		DlWindReq    *bool
		ExpOnly      *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "dlDataQue",
			fields: fields{
				Command: structs.MsgDlDataQue,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataQue{
				Command:      tt.fields.Command,
				OpId:         tt.fields.OpId,
				EpEui:        tt.fields.EpEui,
				QueId:        tt.fields.QueId,
				CntDepend:    tt.fields.CntDepend,
				PacketCnt:    tt.fields.PacketCnt,
				UserData:     tt.fields.UserData,
				Format:       tt.fields.Format,
				Prio:         tt.fields.Prio,
				ResponseExp:  tt.fields.ResponseExp,
				ResponsePrio: tt.fields.ResponsePrio,
				DlWindReq:    tt.fields.DlWindReq,
				ExpOnly:      tt.fields.ExpOnly,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlDataQue.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataQue_GetCommand(t *testing.T) {
	type fields struct {
		Command      structs.Command
		OpId         int64
		EpEui        common.EUI64
		QueId        uint64
		CntDepend    bool
		PacketCnt    *[]uint32
		UserData     [][]byte
		Format       *byte
		Prio         *float32
		ResponseExp  *bool
		ResponsePrio *bool
		DlWindReq    *bool
		ExpOnly      *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "dlDataQue",
			fields: fields{
				Command: structs.MsgDlDataQue,
				OpId:    10,
			},
			want: structs.MsgDlDataQue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataQue{
				Command:      tt.fields.Command,
				OpId:         tt.fields.OpId,
				EpEui:        tt.fields.EpEui,
				QueId:        tt.fields.QueId,
				CntDepend:    tt.fields.CntDepend,
				PacketCnt:    tt.fields.PacketCnt,
				UserData:     tt.fields.UserData,
				Format:       tt.fields.Format,
				Prio:         tt.fields.Prio,
				ResponseExp:  tt.fields.ResponseExp,
				ResponsePrio: tt.fields.ResponsePrio,
				DlWindReq:    tt.fields.DlWindReq,
				ExpOnly:      tt.fields.ExpOnly,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlDataQue.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataQue_SetOpId(t *testing.T) {
	type fields struct {
		Command      structs.Command
		OpId         int64
		EpEui        common.EUI64
		QueId        uint64
		CntDepend    bool
		PacketCnt    *[]uint32
		UserData     [][]byte
		Format       *byte
		Prio         *float32
		ResponseExp  *bool
		ResponsePrio *bool
		DlWindReq    *bool
		ExpOnly      *bool
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
			name: "dlDataQue",
			fields: fields{
				Command: structs.MsgDlDataQue,
				OpId:    10,
			},
			args: args{opId: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataQue{
				Command:      tt.fields.Command,
				OpId:         tt.fields.OpId,
				EpEui:        tt.fields.EpEui,
				QueId:        tt.fields.QueId,
				CntDepend:    tt.fields.CntDepend,
				PacketCnt:    tt.fields.PacketCnt,
				UserData:     tt.fields.UserData,
				Format:       tt.fields.Format,
				Prio:         tt.fields.Prio,
				ResponseExp:  tt.fields.ResponseExp,
				ResponsePrio: tt.fields.ResponsePrio,
				DlWindReq:    tt.fields.DlWindReq,
				ExpOnly:      tt.fields.ExpOnly,
			}
			m.SetOpId(tt.args.opId)

			if m.OpId != tt.args.opId {
				t.Errorf("DlDataQue.SetOpId() = %v, want %v", m.OpId, tt.args.opId)
			}
		})
	}
}

func TestNewDlDataQueRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DlDataQueRsp
	}{
		{
			name: "dlDataQueRsp",
			args: args{
				opId: 0,
			},
			want: DlDataQueRsp{
				Command: structs.MsgDlDataQueRsp,
				OpId:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlDataQueRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataQueRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataQueRsp_GetOpId(t *testing.T) {
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
			name: "dlDataQueRsp",
			fields: fields{
				Command: structs.MsgDlDataQueRsp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataQueRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlDataQueRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataQueRsp_GetCommand(t *testing.T) {
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
			name: "dlDataQueRsp",
			fields: fields{
				Command: structs.MsgDlDataQueRsp,
				OpId:    10,
			},
			want: structs.MsgDlDataQueRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataQueRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlDataQueRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDlDataQueCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want DlDataQueCmp
	}{
		{
			name: "dlDataQueCmp",
			args: args{
				opId: 0,
			},
			want: DlDataQueCmp{
				Command: structs.MsgDlDataQueCmp,
				OpId:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDlDataQueCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDlDataQueCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataQueCmp_GetOpId(t *testing.T) {
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
			name: "dlDataQueCmp",
			fields: fields{
				Command: structs.MsgDlDataQueCmp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataQueCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("DlDataQueCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDlDataQueCmp_GetCommand(t *testing.T) {
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
			name: "dlDataQueCmp",
			fields: fields{
				Command: structs.MsgDlDataQueCmp,
				OpId:    10,
			},
			want: structs.MsgDlDataQueCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DlDataQueCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DlDataQueCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
