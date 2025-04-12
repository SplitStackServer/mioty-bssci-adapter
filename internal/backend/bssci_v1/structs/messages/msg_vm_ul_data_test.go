package messages

import (
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/backend/events"
	"mioty-bssci-adapter/internal/common"
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewVmUlData(t *testing.T) {
	type args struct {
		opId       int64
		macType    int64
		userData   []byte
		trxTime    uint64
		freqOff    float64
		snr        float64
		rssi       float64
		eqSnr      *float64
		subpackets *Subpackets
		carrSpace  byte
		pattGrp    byte
		pattNum    byte
		crc        [2]byte
	}
	tests := []struct {
		name string
		args args
		want VmUlData
	}{
		{
			name: "vmUlData",
			args: args{
				opId:       1,
				macType:    0,
				userData:   []byte{},
				trxTime:    0,
				freqOff:    0,
				snr:        0,
				rssi:       0,
				eqSnr:      new(float64),
				subpackets: &Subpackets{},
				carrSpace:  0,
				pattGrp:    0,
				pattNum:    0,
				crc:        [2]byte{},
			},
			want: VmUlData{
				Command:    structs.MsgVmUlData,
				OpId:       1,
				MacType:    0,
				UserData:   []byte{},
				TrxTime:    0,
				SysTime:    0,
				FreqOff:    0,
				SNR:        0,
				RSSI:       0,
				EqSnr:      new(float64),
				Subpackets: &Subpackets{},
				CarrSpace:  0,
				PattGrp:    0,
				PattNum:    0,
				CRC:        [2]byte{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmUlData(tt.args.opId, tt.args.macType, tt.args.userData, tt.args.trxTime, tt.args.freqOff, tt.args.snr, tt.args.rssi, tt.args.eqSnr, tt.args.subpackets, tt.args.carrSpace, tt.args.pattGrp, tt.args.pattNum, tt.args.crc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmUlData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmUlData_GetOpId(t *testing.T) {
	type fields struct {
		Command    structs.Command
		OpId       int64
		MacType    int64
		UserData   []byte
		TrxTime    uint64
		SysTime    uint64
		FreqOff    float64
		SNR        float64
		RSSI       float64
		EqSnr      *float64
		Subpackets *Subpackets
		CarrSpace  byte
		PattGrp    byte
		PattNum    byte
		CRC        [2]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "vmUlDataRsp",
			fields: fields{
				Command:    structs.MsgVmUlData,
				OpId:       1,
				MacType:    0,
				UserData:   []byte{},
				TrxTime:    0,
				SysTime:    0,
				FreqOff:    0,
				SNR:        0,
				RSSI:       0,
				EqSnr:      new(float64),
				Subpackets: &Subpackets{},
				CarrSpace:  0,
				PattGrp:    0,
				PattNum:    0,
				CRC:        [2]byte{},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmUlData{
				Command:    tt.fields.Command,
				OpId:       tt.fields.OpId,
				MacType:    tt.fields.MacType,
				UserData:   tt.fields.UserData,
				TrxTime:    tt.fields.TrxTime,
				SysTime:    tt.fields.SysTime,
				FreqOff:    tt.fields.FreqOff,
				SNR:        tt.fields.SNR,
				RSSI:       tt.fields.RSSI,
				EqSnr:      tt.fields.EqSnr,
				Subpackets: tt.fields.Subpackets,
				CarrSpace:  tt.fields.CarrSpace,
				PattGrp:    tt.fields.PattGrp,
				PattNum:    tt.fields.PattNum,
				CRC:        tt.fields.CRC,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmUlData.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmUlData_GetCommand(t *testing.T) {
	type fields struct {
		Command    structs.Command
		OpId       int64
		MacType    int64
		UserData   []byte
		TrxTime    uint64
		SysTime    uint64
		FreqOff    float64
		SNR        float64
		RSSI       float64
		EqSnr      *float64
		Subpackets *Subpackets
		CarrSpace  byte
		PattGrp    byte
		PattNum    byte
		CRC        [2]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "vmUlDataRsp",
			fields: fields{
				Command:    structs.MsgVmUlData,
				OpId:       1,
				MacType:    0,
				UserData:   []byte{},
				TrxTime:    0,
				SysTime:    0,
				FreqOff:    0,
				SNR:        0,
				RSSI:       0,
				EqSnr:      new(float64),
				Subpackets: &Subpackets{},
				CarrSpace:  0,
				PattGrp:    0,
				PattNum:    0,
				CRC:        [2]byte{},
			},
			want: structs.MsgVmUlData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmUlData{
				Command:    tt.fields.Command,
				OpId:       tt.fields.OpId,
				MacType:    tt.fields.MacType,
				UserData:   tt.fields.UserData,
				TrxTime:    tt.fields.TrxTime,
				SysTime:    tt.fields.SysTime,
				FreqOff:    tt.fields.FreqOff,
				SNR:        tt.fields.SNR,
				RSSI:       tt.fields.RSSI,
				EqSnr:      tt.fields.EqSnr,
				Subpackets: tt.fields.Subpackets,
				CarrSpace:  tt.fields.CarrSpace,
				PattGrp:    tt.fields.PattGrp,
				PattNum:    tt.fields.PattNum,
				CRC:        tt.fields.CRC,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmUlData.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmUlData_GetEventType(t *testing.T) {
	type fields struct {
		Command    structs.Command
		OpId       int64
		MacType    int64
		UserData   []byte
		TrxTime    uint64
		SysTime    uint64
		FreqOff    float64
		SNR        float64
		RSSI       float64
		EqSnr      *float64
		Subpackets *Subpackets
		CarrSpace  byte
		PattGrp    byte
		PattNum    byte
		CRC        [2]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   events.EventType
	}{
		{
			name: "vmUlData",
			fields: fields{
				Command:    structs.MsgVmUlData,
				OpId:       1,
				MacType:    0,
				UserData:   []byte{},
				TrxTime:    0,
				SysTime:    0,
				FreqOff:    0,
				SNR:        0,
				RSSI:       0,
				EqSnr:      new(float64),
				Subpackets: &Subpackets{},
				CarrSpace:  0,
				PattGrp:    0,
				PattNum:    0,
				CRC:        [2]byte{},
			},
			want: events.EventTypeEpUl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmUlData{
				Command:    tt.fields.Command,
				OpId:       tt.fields.OpId,
				MacType:    tt.fields.MacType,
				UserData:   tt.fields.UserData,
				TrxTime:    tt.fields.TrxTime,
				SysTime:    tt.fields.SysTime,
				FreqOff:    tt.fields.FreqOff,
				SNR:        tt.fields.SNR,
				RSSI:       tt.fields.RSSI,
				EqSnr:      tt.fields.EqSnr,
				Subpackets: tt.fields.Subpackets,
				CarrSpace:  tt.fields.CarrSpace,
				PattGrp:    tt.fields.PattGrp,
				PattNum:    tt.fields.PattNum,
				CRC:        tt.fields.CRC,
			}
			if got := m.GetEventType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmUlData.GetEventType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmUlData_IntoProto(t *testing.T) {

	var testTRxTime uint64 = 1000000000000005

	testTRxTimePb := timestamppb.Timestamp{
		Seconds: int64(1000000),
		Nanos:   int32(5),
	}

	type fields struct {
		Command    structs.Command
		OpId       int64
		MacType    int64
		UserData   []byte
		TrxTime    uint64
		SysTime    uint64
		FreqOff    float64
		SNR        float64
		RSSI       float64
		EqSnr      *float64
		Subpackets *Subpackets
		CarrSpace  byte
		PattGrp    byte
		PattNum    byte
		CRC        [2]byte
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
			name: "vmUlDataRsp",
			fields: fields{
				Command:    structs.MsgVmUlData,
				OpId:       1,
				MacType:    0,
				UserData:   []byte{},
				TrxTime:    testTRxTime,
				SysTime:    0,
				FreqOff:    0,
				SNR:        0,
				RSSI:       0,
				EqSnr:      nil,
				Subpackets: &Subpackets{},
				CarrSpace:  0,
				PattGrp:    0,
				PattNum:    0,
				CRC:        [2]byte{},
			},
			args: args{bsEui: common.EUI64{1}},
			want: &msg.ProtoEndnodeMessage{
				BsEui:      "1",
				EndnodeEui: "0",
				V1: &msg.ProtoEndnodeMessage_VmUlData{
					VmUlData: &msg.EndnodeVariableMacUlDataMessage{
						Data:    []byte{},
						MacType: 0,
						Meta: &msg.EndnodeUplinkMetadata{
							RxTime:        &testTRxTimePb,
							RxDuration:    nil,
							PacketCnt:     0,
							Profile:       nil,
							Rssi:          0,
							Snr:           0,
							EqSnr:         nil,
							SubpacketInfo: []*msg.EndnodeUplinkSubpacket{},
						},
						FreqOff:   0,
						CarrSpace: 0,
						PattGrp:   0,
						PattNum:   0,
						Crc:       0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmUlData{
				Command:    tt.fields.Command,
				OpId:       tt.fields.OpId,
				MacType:    tt.fields.MacType,
				UserData:   tt.fields.UserData,
				TrxTime:    tt.fields.TrxTime,
				SysTime:    tt.fields.SysTime,
				FreqOff:    tt.fields.FreqOff,
				SNR:        tt.fields.SNR,
				RSSI:       tt.fields.RSSI,
				EqSnr:      tt.fields.EqSnr,
				Subpackets: tt.fields.Subpackets,
				CarrSpace:  tt.fields.CarrSpace,
				PattGrp:    tt.fields.PattGrp,
				PattNum:    tt.fields.PattNum,
				CRC:        tt.fields.CRC,
			}
			if got := m.IntoProto(tt.args.bsEui); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmUlData.IntoProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVmUlDataRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want VmUlDataRsp
	}{
		{
			name: "vmUlDataRsp",
			args: args{1},
			want: VmUlDataRsp{
				Command: structs.MsgVmUlDataRsp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmUlDataRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmUlDataRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmUlDataRsp_GetOpId(t *testing.T) {
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
			name: "vmUlDataRsp",
			fields: fields{
				structs.MsgVmUlDataRsp,
				1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmUlDataRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmUlDataRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmUlDataRsp_GetCommand(t *testing.T) {
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
			name: "vmUlDataRsp",
			fields: fields{
				structs.MsgVmUlDataRsp,
				1,
			},
			want: structs.MsgVmUlDataRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmUlDataRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmUlDataRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVmUlDataCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want VmUlDataCmp
	}{
		{
			name: "vmUlDataCmp",
			args: args{1},
			want: VmUlDataCmp{
				Command: structs.MsgVmUlDataCmp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVmUlDataCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVmUlDataCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmUlDataCmp_GetOpId(t *testing.T) {
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
			name: "vmUlDataCmp",
			fields: fields{
				structs.MsgVmUlDataCmp,
				1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmUlDataCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("VmUlDataCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVmUlDataCmp_GetCommand(t *testing.T) {
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
			name: "vmUlDataCmp",
			fields: fields{
				structs.MsgVmUlDataCmp,
				1,
			},
			want: structs.MsgVmUlDataCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &VmUlDataCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VmUlDataCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
