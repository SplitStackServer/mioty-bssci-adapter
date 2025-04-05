package messages

import (
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/api/rsp"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/backend/events"
	"mioty-bssci-adapter/internal/common"
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewAtt(t *testing.T) {
	type args struct {
		opId        int64
		epEui       common.EUI64
		rxTime      uint64
		rxDuration  *uint64
		attachCnt   uint32
		snr         float64
		rssi        float64
		eqSnr       *float64
		profile     *string
		subpackets  *Subpackets
		nonce       [4]byte
		sign        [4]byte
		shAddr      *uint16
		dualChan    bool
		repetition  bool
		wideCarrOff bool
		longBlkDist bool
	}
	tests := []struct {
		name string
		args args
		want Att
	}{
		{
			name: "att",
			args: args{
				opId:        0,
				epEui:       common.EUI64{},
				rxTime:      1,
				rxDuration:  nil,
				attachCnt:   2,
				snr:         3.0,
				rssi:        -100.0,
				eqSnr:       nil,
				profile:     new(string),
				subpackets:  &Subpackets{},
				sign:        [4]byte{},
				nonce:       [4]byte{},
				shAddr:      nil,
				dualChan:    false,
				repetition:  false,
				wideCarrOff: false,
				longBlkDist: false,
			},
			want: Att{
				Command:     structs.MsgAtt,
				OpId:        0,
				EpEui:       common.EUI64{},
				RxTime:      1,
				RxDuration:  nil,
				AttachCnt:   2,
				SNR:         3.0,
				RSSI:        -100.0,
				EqSnr:       nil,
				Profile:     new(string),
				Subpackets:  &Subpackets{},
				Sign:        [4]byte{},
				Nonce:       [4]byte{},
				ShAddr:      nil,
				DualChan:    false,
				Repetition:  false,
				WideCarrOff: false,
				LongBlkDist: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAtt(tt.args.opId, tt.args.epEui, tt.args.rxTime, tt.args.rxDuration, tt.args.attachCnt, tt.args.snr, tt.args.rssi, tt.args.eqSnr, tt.args.profile, tt.args.subpackets, tt.args.nonce, tt.args.sign, tt.args.shAddr, tt.args.dualChan, tt.args.repetition, tt.args.wideCarrOff, tt.args.longBlkDist); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAtt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtt_GetOpId(t *testing.T) {
	type fields struct {
		Command     structs.Command
		OpId        int64
		EpEui       common.EUI64
		RxTime      uint64
		RxDuration  *uint64
		AttachCnt   uint32
		SNR         float64
		RSSI        float64
		EqSnr       *float64
		Profile     *string
		Subpackets  *Subpackets
		Nonce       [4]byte
		Sign        [4]byte
		ShAddr      *uint16
		DualChan    bool
		Repetition  bool
		WideCarrOff bool
		LongBlkDist bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "att",
			fields: fields{
				Command: structs.MsgAtt,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Att{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				EpEui:       tt.fields.EpEui,
				RxTime:      tt.fields.RxTime,
				RxDuration:  tt.fields.RxDuration,
				AttachCnt:   tt.fields.AttachCnt,
				SNR:         tt.fields.SNR,
				RSSI:        tt.fields.RSSI,
				EqSnr:       tt.fields.EqSnr,
				Profile:     tt.fields.Profile,
				Subpackets:  tt.fields.Subpackets,
				Nonce:       tt.fields.Nonce,
				Sign:        tt.fields.Sign,
				ShAddr:      tt.fields.ShAddr,
				DualChan:    tt.fields.DualChan,
				Repetition:  tt.fields.Repetition,
				WideCarrOff: tt.fields.WideCarrOff,
				LongBlkDist: tt.fields.LongBlkDist,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("Att.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtt_GetCommand(t *testing.T) {
	type fields struct {
		Command     structs.Command
		OpId        int64
		EpEui       common.EUI64
		RxTime      uint64
		RxDuration  *uint64
		AttachCnt   uint32
		SNR         float64
		RSSI        float64
		EqSnr       *float64
		Profile     *string
		Subpackets  *Subpackets
		Nonce       [4]byte
		Sign        [4]byte
		ShAddr      *uint16
		DualChan    bool
		Repetition  bool
		WideCarrOff bool
		LongBlkDist bool
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "att",
			fields: fields{
				Command: structs.MsgAtt,
				OpId:    10,
			},
			want: structs.MsgAtt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Att{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				EpEui:       tt.fields.EpEui,
				RxTime:      tt.fields.RxTime,
				RxDuration:  tt.fields.RxDuration,
				AttachCnt:   tt.fields.AttachCnt,
				SNR:         tt.fields.SNR,
				RSSI:        tt.fields.RSSI,
				EqSnr:       tt.fields.EqSnr,
				Profile:     tt.fields.Profile,
				Subpackets:  tt.fields.Subpackets,
				Nonce:       tt.fields.Nonce,
				Sign:        tt.fields.Sign,
				ShAddr:      tt.fields.ShAddr,
				DualChan:    tt.fields.DualChan,
				Repetition:  tt.fields.Repetition,
				WideCarrOff: tt.fields.WideCarrOff,
				LongBlkDist: tt.fields.LongBlkDist,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Att.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtt_GetEventType(t *testing.T) {
	type fields struct {
		Command     structs.Command
		OpId        int64
		EpEui       common.EUI64
		RxTime      uint64
		RxDuration  *uint64
		AttachCnt   uint32
		SNR         float64
		RSSI        float64
		EqSnr       *float64
		Profile     *string
		Subpackets  *Subpackets
		Nonce       [4]byte
		Sign        [4]byte
		ShAddr      *uint16
		DualChan    bool
		Repetition  bool
		WideCarrOff bool
		LongBlkDist bool
	}
	tests := []struct {
		name   string
		fields fields
		want   events.EventType
	}{
		{
			name: "att",
			fields: fields{
				Command: structs.MsgAtt,
				OpId:    10,
				EpEui:   common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
			},
			want: events.EventTypeEpOtaa,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Att{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				EpEui:       tt.fields.EpEui,
				RxTime:      tt.fields.RxTime,
				RxDuration:  tt.fields.RxDuration,
				AttachCnt:   tt.fields.AttachCnt,
				SNR:         tt.fields.SNR,
				RSSI:        tt.fields.RSSI,
				EqSnr:       tt.fields.EqSnr,
				Profile:     tt.fields.Profile,
				Subpackets:  tt.fields.Subpackets,
				Nonce:       tt.fields.Nonce,
				Sign:        tt.fields.Sign,
				ShAddr:      tt.fields.ShAddr,
				DualChan:    tt.fields.DualChan,
				Repetition:  tt.fields.Repetition,
				WideCarrOff: tt.fields.WideCarrOff,
				LongBlkDist: tt.fields.LongBlkDist,
			}
			if got := m.GetEventType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Att.GetEventType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtt_IntoProto(t *testing.T) {
	var testRxTime uint64 = 1000000000000005

	testRxTimePb := timestamppb.Timestamp{
		Seconds: int64(1000000),
		Nanos:   int32(5),
	}

	var testRxDuration uint64 = 1000001005

	testRxDurationPb := durationpb.Duration{
		Seconds: int64(1),
		Nanos:   int32(1005),
	}

	var testShAddr uint16 = 0x1010
	var testShAddr32 uint32 = 0x1010

	type fields struct {
		Command     structs.Command
		OpId        int64
		EpEui       common.EUI64
		RxTime      uint64
		RxDuration  *uint64
		AttachCnt   uint32
		SNR         float64
		RSSI        float64
		EqSnr       *float64
		Profile     *string
		Subpackets  *Subpackets
		Nonce       [4]byte
		Sign        [4]byte
		ShAddr      *uint16
		DualChan    bool
		Repetition  bool
		WideCarrOff bool
		LongBlkDist bool
	}
	type args struct {
		bsEui common.EUI64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *msg.ProtoEndnodeMessage
	}{
		{
			name: "att1",
			fields: fields{
				Command:    structs.MsgDet,
				OpId:       10,
				EpEui:      common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
				RxTime:     testRxTime,
				RxDuration: &testRxDuration,
				ShAddr:     &testShAddr,
				AttachCnt:  2,
				SNR:        3.0,
				RSSI:       -100.0,
				EqSnr:      nil,
				Profile:    new(string),
				Subpackets: &Subpackets{},
				Sign:       [4]byte{3, 2, 1, 0},
				Nonce:      [4]byte{7, 6, 5, 4},
			},
			want: &msg.ProtoEndnodeMessage{
				BsEui:      0,
				EndnodeEui: 0x0706050403020100,
				V1: &msg.ProtoEndnodeMessage_Att{
					Att: &msg.EndnodeAttMessage{
						OpId:          10,
						Sign:          0x00010203,
						Nonce:         0x04050607,
						AttachmentCnt: 2,
						ShAddr:        &testShAddr32,
						Meta: &msg.EndnodeUplinkMetadata{
							RxTime:        &testRxTimePb,
							RxDuration:    &testRxDurationPb,
							PacketCnt:     0,
							Profile:       new(string),
							Rssi:          -100.0,
							Snr:           3.0,
							EqSnr:         nil,
							SubpacketInfo: []*msg.EndnodeUplinkSubpacket{},
						},
					},
				},
			},
		},
		{
			name: "att2",
			fields: fields{
				Command:    structs.MsgDet,
				OpId:       10,
				EpEui:      common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
				RxTime:     testRxTime,
				RxDuration: &testRxDuration,
				AttachCnt:  2,
				SNR:        3.0,
				RSSI:       -100.0,
				EqSnr:      nil,
				Profile:    new(string),
				Subpackets: &Subpackets{},
				Sign:       [4]byte{3, 2, 1, 0},
				Nonce:      [4]byte{7, 6, 5, 4},
			},
			want: &msg.ProtoEndnodeMessage{
				BsEui:      0,
				EndnodeEui: 0x0706050403020100,
				V1: &msg.ProtoEndnodeMessage_Att{
					Att: &msg.EndnodeAttMessage{
						OpId:          10,
						Sign:          0x00010203,
						Nonce:         0x04050607,
						AttachmentCnt: 2,
						ShAddr:        nil,
						Meta: &msg.EndnodeUplinkMetadata{
							RxTime:        &testRxTimePb,
							RxDuration:    &testRxDurationPb,
							PacketCnt:     0,
							Profile:       new(string),
							Rssi:          -100.0,
							Snr:           3.0,
							EqSnr:         nil,
							SubpacketInfo: []*msg.EndnodeUplinkSubpacket{},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Att{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				EpEui:       tt.fields.EpEui,
				RxTime:      tt.fields.RxTime,
				RxDuration:  tt.fields.RxDuration,
				AttachCnt:   tt.fields.AttachCnt,
				SNR:         tt.fields.SNR,
				RSSI:        tt.fields.RSSI,
				EqSnr:       tt.fields.EqSnr,
				Profile:     tt.fields.Profile,
				Subpackets:  tt.fields.Subpackets,
				Nonce:       tt.fields.Nonce,
				Sign:        tt.fields.Sign,
				ShAddr:      tt.fields.ShAddr,
				DualChan:    tt.fields.DualChan,
				Repetition:  tt.fields.Repetition,
				WideCarrOff: tt.fields.WideCarrOff,
				LongBlkDist: tt.fields.LongBlkDist,
			}
			if got := m.IntoProto(tt.args.bsEui); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Att.IntoProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAttRsp(t *testing.T) {
	type args struct {
		opId          int64
		nwkSessionKey [16]byte
		shAddr        *uint16
	}
	tests := []struct {
		name string
		args args
		want AttRsp
	}{
		{
			name: "attRsp",
			args: args{
				opId: 10,
			},
			want: AttRsp{
				Command: structs.MsgAttRsp,
				OpId:    10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAttRsp(tt.args.opId, tt.args.nwkSessionKey, tt.args.shAddr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAttRspFromProto(t *testing.T) {

	var testShAddr uint16 = 0x1010
	var testShAddr32 uint32 = 0x1010

	type args struct {
		opId int64
		pb   *rsp.EndnodeAttachResponse
	}
	tests := []struct {
		name    string
		args    args
		want    *AttRsp
		wantErr bool
	}{
		{
			name: "attRsp",
			args: args{
				opId: 10,
				pb: &rsp.EndnodeAttachResponse{
					EndnodeEui:    0x0001020304050607,
					ShAddr:        nil,
					NwkSessionKey: []byte{3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0},
				},
			},
			want: &AttRsp{
				Command:       structs.MsgAttRsp,
				OpId:          10,
				ShAddr:        nil,
				NwkSessionKey: [16]byte{3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0},
			},
			wantErr: false,
		},
		{
			name: "attRsp_shAddr",
			args: args{
				opId: 10,
				pb: &rsp.EndnodeAttachResponse{
					EndnodeEui:    0x0001020304050607,
					ShAddr:        &testShAddr32,
					NwkSessionKey: []byte{3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0},
				},
			},
			want: &AttRsp{
				Command:       structs.MsgAttRsp,
				OpId:          10,
				ShAddr:        &testShAddr,
				NwkSessionKey: [16]byte{3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0},
			},
			wantErr: false,
		},
		{
			name: "attRsp_invalid_NwkSessionKey",
			args: args{
				opId: 10,
				pb: &rsp.EndnodeAttachResponse{
					EndnodeEui:    0x0001020304050607,
					ShAddr:        &testShAddr32,
					NwkSessionKey: []byte{},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "attRspNil",
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
			got, err := NewAttRspFromProto(tt.args.opId, tt.args.pb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAttRspFromProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttRspFromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttRsp_GetOpId(t *testing.T) {
	type fields struct {
		Command       structs.Command
		OpId          int64
		NwkSessionKey [16]byte
		ShAddr        *uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "attRsp",
			fields: fields{
				Command: structs.MsgAttRsp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AttRsp{
				Command:       tt.fields.Command,
				OpId:          tt.fields.OpId,
				NwkSessionKey: tt.fields.NwkSessionKey,
				ShAddr:        tt.fields.ShAddr,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("AttRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttRsp_GetCommand(t *testing.T) {
	type fields struct {
		Command       structs.Command
		OpId          int64
		NwkSessionKey [16]byte
		ShAddr        *uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "attRsp",
			fields: fields{
				Command: structs.MsgAttRsp,
				OpId:    10,
			},
			want: structs.MsgAttRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AttRsp{
				Command:       tt.fields.Command,
				OpId:          tt.fields.OpId,
				NwkSessionKey: tt.fields.NwkSessionKey,
				ShAddr:        tt.fields.ShAddr,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAttCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want AttCmp
	}{
		{
			name: "attCmp",
			args: args{
				opId: 10,
			},
			want: AttCmp{
				Command: structs.MsgAttCmp,
				OpId:    10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAttCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttCmp_GetOpId(t *testing.T) {
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
			name: "attCmp",
			fields: fields{
				Command: structs.MsgAtt,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AttCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("AttCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttCmp_GetCommand(t *testing.T) {
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
			name: "attCmp",
			fields: fields{
				Command: structs.MsgAttCmp,
				OpId:    10,
			},
			want: structs.MsgAttCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AttCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
