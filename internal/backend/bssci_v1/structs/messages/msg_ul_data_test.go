package messages

import (
	"reflect"
	"testing"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewUlData(t *testing.T) {
	type args struct {
		opId        int64
		epEui       common.EUI64
		rxTime      uint64
		rxDuration  *uint64
		packetCnt   uint32
		snr         float64
		rssi        float64
		eqSnr       *float64
		profile     *string
		mode        *string
		subpackets  *Subpackets
		userData    []byte
		format      *byte
		dlOpen      bool
		responseExp bool
		DlAck       bool
	}
	tests := []struct {
		name string
		args args
		want UlData
	}{
		{
			name: "ulData",
			args: args{
				opId:        10,
				epEui:       common.EUI64{},
				rxTime:      1,
				rxDuration:  nil,
				packetCnt:   2,
				snr:         3.0,
				rssi:        -100.0,
				eqSnr:       nil,
				profile:     nil,
				subpackets:  nil,
				mode:        nil,
				userData:    []byte{0, 1, 2, 3},
				format:      nil,
				dlOpen:      false,
				responseExp: false,
				DlAck:       false,
			},
			want: UlData{
				Command:     structs.MsgUlData,
				OpId:        10,
				EpEui:       common.EUI64{},
				RxTime:      1,
				PacketCnt:   2,
				SNR:         3.0,
				RSSI:        -100.0,
				UserData:    []byte{0, 1, 2, 3},
				DlOpen:      false,
				ResponseExp: false,
				DlAck:       false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUlData(tt.args.opId, tt.args.epEui, tt.args.rxTime, tt.args.rxDuration, tt.args.packetCnt, tt.args.snr, tt.args.rssi, tt.args.eqSnr, tt.args.profile, tt.args.mode, tt.args.subpackets, tt.args.userData, tt.args.format, tt.args.dlOpen, tt.args.responseExp, tt.args.DlAck); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUlData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlData_GetOpId(t *testing.T) {
	type fields struct {
		Command     structs.Command
		OpId        int64
		EpEui       common.EUI64
		RxTime      uint64
		RxDuration  *uint64
		PacketCnt   uint32
		SNR         float64
		RSSI        float64
		EqSnr       *float64
		Profile     *string
		Mode        *string
		Subpackets  *Subpackets
		UserData    []byte
		Format      *byte
		DlOpen      bool
		ResponseExp bool
		DlAck       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "ulData",
			fields: fields{
				Command:     structs.MsgUlData,
				OpId:        10,
				EpEui:       common.EUI64{},
				RxTime:      1,
				PacketCnt:   2,
				SNR:         3.0,
				RSSI:        -100.0,
				UserData:    []byte{0, 1, 2, 3},
				DlOpen:      false,
				ResponseExp: false,
				DlAck:       false,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UlData{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				EpEui:       tt.fields.EpEui,
				RxTime:      tt.fields.RxTime,
				RxDuration:  tt.fields.RxDuration,
				PacketCnt:   tt.fields.PacketCnt,
				SNR:         tt.fields.SNR,
				RSSI:        tt.fields.RSSI,
				EqSnr:       tt.fields.EqSnr,
				Profile:     tt.fields.Profile,
				Mode:        tt.fields.Mode,
				Subpackets:  tt.fields.Subpackets,
				UserData:    tt.fields.UserData,
				Format:      tt.fields.Format,
				DlOpen:      tt.fields.DlOpen,
				ResponseExp: tt.fields.ResponseExp,
				DlAck:       tt.fields.DlAck,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("UlData.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlData_GetCommand(t *testing.T) {
	type fields struct {
		Command     structs.Command
		OpId        int64
		EpEui       common.EUI64
		RxTime      uint64
		RxDuration  *uint64
		PacketCnt   uint32
		SNR         float64
		RSSI        float64
		EqSnr       *float64
		Profile     *string
		Mode        *string
		Subpackets  *Subpackets
		UserData    []byte
		Format      *byte
		DlOpen      bool
		ResponseExp bool
		DlAck       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "ulData",
			fields: fields{
				Command:     structs.MsgUlData,
				OpId:        10,
				EpEui:       common.EUI64{},
				RxTime:      1,
				PacketCnt:   2,
				SNR:         3.0,
				RSSI:        -100.0,
				UserData:    []byte{0, 1, 2, 3},
				DlOpen:      false,
				ResponseExp: false,
				DlAck:       false,
			},
			want: structs.MsgUlData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UlData{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				EpEui:       tt.fields.EpEui,
				RxTime:      tt.fields.RxTime,
				RxDuration:  tt.fields.RxDuration,
				PacketCnt:   tt.fields.PacketCnt,
				SNR:         tt.fields.SNR,
				RSSI:        tt.fields.RSSI,
				EqSnr:       tt.fields.EqSnr,
				Profile:     tt.fields.Profile,
				Mode:        tt.fields.Mode,
				Subpackets:  tt.fields.Subpackets,
				UserData:    tt.fields.UserData,
				Format:      tt.fields.Format,
				DlOpen:      tt.fields.DlOpen,
				ResponseExp: tt.fields.ResponseExp,
				DlAck:       tt.fields.DlAck,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UlData.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlData_GetEventType(t *testing.T) {
	type fields struct {
		Command     structs.Command
		OpId        int64
		EpEui       common.EUI64
		RxTime      uint64
		RxDuration  *uint64
		PacketCnt   uint32
		SNR         float64
		RSSI        float64
		EqSnr       *float64
		Profile     *string
		Mode        *string
		Subpackets  *Subpackets
		UserData    []byte
		Format      *byte
		DlOpen      bool
		ResponseExp bool
		DlAck       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   events.EventType
	}{
		{
			name: "ulData",
			fields: fields{
				Command:     structs.MsgUlData,
				OpId:        10,
				EpEui:       common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
				RxTime:      1,
				PacketCnt:   2,
				SNR:         3.0,
				RSSI:        -100.0,
				UserData:    []byte{0, 1, 2, 3},
				DlOpen:      false,
				ResponseExp: false,
				DlAck:       false,
			},
			want: events.EventTypeEpUl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UlData{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				EpEui:       tt.fields.EpEui,
				RxTime:      tt.fields.RxTime,
				RxDuration:  tt.fields.RxDuration,
				PacketCnt:   tt.fields.PacketCnt,
				SNR:         tt.fields.SNR,
				RSSI:        tt.fields.RSSI,
				EqSnr:       tt.fields.EqSnr,
				Profile:     tt.fields.Profile,
				Mode:        tt.fields.Mode,
				Subpackets:  tt.fields.Subpackets,
				UserData:    tt.fields.UserData,
				Format:      tt.fields.Format,
				DlOpen:      tt.fields.DlOpen,
				ResponseExp: tt.fields.ResponseExp,
				DlAck:       tt.fields.DlAck,
			}
			if got := m.GetEventType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UlData.GetEventType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlData_IntoProto(t *testing.T) {

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

	var testMode string = "test"
	var testFormat byte = 0x83

	type fields struct {
		Command     structs.Command
		OpId        int64
		EpEui       common.EUI64
		RxTime      uint64
		RxDuration  *uint64
		PacketCnt   uint32
		SNR         float64
		RSSI        float64
		EqSnr       *float64
		Profile     *string
		Mode        *string
		Subpackets  *Subpackets
		UserData    []byte
		Format      *byte
		DlOpen      bool
		ResponseExp bool
		DlAck       bool
	}
	type args struct {
		bsEui common.EUI64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *bs.EndnodeUplink
	}{
		{
			name: "ulData1",
			fields: fields{
				Command:    structs.MsgUlData,
				OpId:       10,
				EpEui:      common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
				UserData:   []byte{0, 1, 2, 3},
				RxTime:     testRxTime,
				RxDuration: &testRxDuration,
				PacketCnt:  2,
				SNR:        3.0,
				RSSI:       -100.0,
				EqSnr:      nil,
				Profile:    new(string),
				Subpackets: &Subpackets{},
				Mode:       &testMode,
				Format:     &testFormat,
				DlOpen:     true,
				DlAck:      true,
			},
			want: &bs.EndnodeUplink{
				BsEui: "0000000000000000",

				Message: &bs.EndnodeUplink_UlData{
					UlData: &bs.EndnodeUlDataMessage{
						Data:   []byte{0, 1, 2, 3},
						EpEui:  "0001020304050607",
						Format: 0x83,
						Mode:   &testMode,
						DlOpen: true,
						DlAck:  true,
					},
				},
				Meta: &bs.EndnodeUplinkMetadata{
					RxTime:        &testRxTimePb,
					RxDuration:    &testRxDurationPb,
					PacketCnt:     2,
					Profile:       new(string),
					Rssi:          -100.0,
					Snr:           3.0,
					EqSnr:         nil,
					SubpacketInfo: []*bs.EndnodeUplinkSubpacket{},
				},
			},
		},
		{
			name: "ulData2",
			fields: fields{
				Command:    structs.MsgUlData,
				OpId:       10,
				EpEui:      common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
				UserData:   []byte{0, 1, 2, 3},
				RxTime:     testRxTime,
				RxDuration: &testRxDuration,
				PacketCnt:  2,
				SNR:        3.0,
				RSSI:       -100.0,
				EqSnr:      nil,
				Profile:    new(string),
				Subpackets: &Subpackets{},
				Mode:       &testMode,
			},
			want: &bs.EndnodeUplink{
				BsEui: "0000000000000000",

				Message: &bs.EndnodeUplink_UlData{
					UlData: &bs.EndnodeUlDataMessage{
						EpEui:  "0001020304050607",
						Data:   []byte{0, 1, 2, 3},
						Format: 0,
						Mode:   &testMode,
					},
				},
				Meta: &bs.EndnodeUplinkMetadata{
					RxTime:        &testRxTimePb,
					RxDuration:    &testRxDurationPb,
					PacketCnt:     2,
					Profile:       new(string),
					Rssi:          -100.0,
					Snr:           3.0,
					EqSnr:         nil,
					SubpacketInfo: []*bs.EndnodeUplinkSubpacket{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UlData{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				EpEui:       tt.fields.EpEui,
				RxTime:      tt.fields.RxTime,
				RxDuration:  tt.fields.RxDuration,
				PacketCnt:   tt.fields.PacketCnt,
				SNR:         tt.fields.SNR,
				RSSI:        tt.fields.RSSI,
				EqSnr:       tt.fields.EqSnr,
				Profile:     tt.fields.Profile,
				Mode:        tt.fields.Mode,
				Subpackets:  tt.fields.Subpackets,
				UserData:    tt.fields.UserData,
				Format:      tt.fields.Format,
				DlOpen:      tt.fields.DlOpen,
				ResponseExp: tt.fields.ResponseExp,
				DlAck:       tt.fields.DlAck,
			}
			if got := m.IntoProto(&tt.args.bsEui); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UlData.IntoProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUlDataRsp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want UlDataRsp
	}{
		{
			name: "ulDataRsp",
			args: args{
				opId: 10,
			},
			want: UlDataRsp{
				Command: structs.MsgUlDataRsp,
				OpId:    10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUlDataRsp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUlDataRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlDataRsp_GetOpId(t *testing.T) {
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
			name: "ulDataRsp",
			fields: fields{
				Command: structs.MsgUlDataRsp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UlDataRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("UlDataRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlDataRsp_GetCommand(t *testing.T) {
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
			name: "ulDataRsp",
			fields: fields{
				Command: structs.MsgUlDataRsp,
				OpId:    10,
			},
			want: structs.MsgUlDataRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UlDataRsp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UlDataRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUlDataCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want UlDataCmp
	}{
		{
			name: "ulDataCmp",
			args: args{
				opId: 10,
			},
			want: UlDataCmp{
				Command: structs.MsgUlDataCmp,
				OpId:    10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUlDataCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUlDataCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlDataCmp_GetOpId(t *testing.T) {
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
			name: "ulDataCmp",
			fields: fields{
				Command: structs.MsgUlDataCmp,
				OpId:    10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UlDataCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("UlDataCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlDataCmp_GetCommand(t *testing.T) {
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
			name: "ulDataCmp",
			fields: fields{
				Command: structs.MsgUlDataCmp,
				OpId:    10,
			},
			want: structs.MsgUlDataCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UlDataCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UlDataCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
