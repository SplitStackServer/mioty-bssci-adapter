package messages

import (
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/common"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCon_GetOpId(t *testing.T) {
	type fields struct {
		Command     structs.Command
		OpId        int64
		Version     string
		BsEui       common.EUI64
		Vendor      *string
		Model       *string
		Name        *string
		SwVersion   *string
		Info        map[string]interface{}
		Bidi        bool
		GeoLocation *GeoLocation
		SnBsUuid    structs.SessionUuid
		SnBsOpId    *int64
		SnScOpId    *int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "con",
			fields: fields{
				Command:     structs.MsgCon,
				OpId:        0,
				Version:     "",
				BsEui:       [8]byte{},
				Vendor:      new(string),
				Model:       new(string),
				Name:        new(string),
				SwVersion:   new(string),
				Info:        map[string]any{},
				Bidi:        false,
				GeoLocation: &GeoLocation{},
				SnBsUuid:    [16]int8{},
				SnBsOpId:    new(int64),
				SnScOpId:    new(int64),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Con{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				Version:     tt.fields.Version,
				BsEui:       tt.fields.BsEui,
				Vendor:      tt.fields.Vendor,
				Model:       tt.fields.Model,
				Name:        tt.fields.Name,
				SwVersion:   tt.fields.SwVersion,
				Info:        tt.fields.Info,
				Bidi:        tt.fields.Bidi,
				GeoLocation: tt.fields.GeoLocation,
				SnBsUuid:    tt.fields.SnBsUuid,
				SnBsOpId:    tt.fields.SnBsOpId,
				SnScOpId:    tt.fields.SnScOpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("Con.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCon_GetCommand(t *testing.T) {
	type fields struct {
		Command     structs.Command
		OpId        int64
		Version     string
		BsEui       common.EUI64
		Vendor      *string
		Model       *string
		Name        *string
		SwVersion   *string
		Info        map[string]interface{}
		Bidi        bool
		GeoLocation *GeoLocation
		SnBsUuid    structs.SessionUuid
		SnBsOpId    *int64
		SnScOpId    *int64
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "con",
			fields: fields{
				Command:     structs.MsgCon,
				OpId:        0,
				Version:     "",
				BsEui:       [8]byte{},
				Vendor:      new(string),
				Model:       new(string),
				Name:        new(string),
				SwVersion:   new(string),
				Info:        map[string]interface{}{},
				Bidi:        false,
				GeoLocation: &GeoLocation{},
				SnBsUuid:    [16]int8{},
				SnBsOpId:    new(int64),
				SnScOpId:    new(int64),
			},
			want: structs.MsgCon,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Con{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				Version:     tt.fields.Version,
				BsEui:       tt.fields.BsEui,
				Vendor:      tt.fields.Vendor,
				Model:       tt.fields.Model,
				Name:        tt.fields.Name,
				SwVersion:   tt.fields.SwVersion,
				Info:        tt.fields.Info,
				Bidi:        tt.fields.Bidi,
				GeoLocation: tt.fields.GeoLocation,
				SnBsUuid:    tt.fields.SnBsUuid,
				SnBsOpId:    tt.fields.SnBsOpId,
				SnScOpId:    tt.fields.SnScOpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Con.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCon_GetEui(t *testing.T) {
	type fields struct {
		Command     structs.Command
		OpId        int64
		Version     string
		BsEui       common.EUI64
		Vendor      *string
		Model       *string
		Name        *string
		SwVersion   *string
		Info        map[string]any
		Bidi        bool
		GeoLocation *GeoLocation
		SnBsUuid    structs.SessionUuid
		SnBsOpId    *int64
		SnScOpId    *int64
	}
	tests := []struct {
		name   string
		fields fields
		want   common.EUI64
	}{
		{
			name: "con",
			fields: fields{
				Command:     structs.MsgCon,
				OpId:        0,
				Version:     "",
				BsEui:       common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
				Vendor:      new(string),
				Model:       new(string),
				Name:        new(string),
				SwVersion:   new(string),
				Info:        map[string]any{},
				Bidi:        false,
				GeoLocation: &GeoLocation{},
				SnBsUuid:    [16]int8{},
				SnBsOpId:    new(int64),
				SnScOpId:    new(int64),
			},
			want: common.EUI64{0, 1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Con{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				Version:     tt.fields.Version,
				BsEui:       tt.fields.BsEui,
				Vendor:      tt.fields.Vendor,
				Model:       tt.fields.Model,
				Name:        tt.fields.Name,
				SwVersion:   tt.fields.SwVersion,
				Info:        tt.fields.Info,
				Bidi:        tt.fields.Bidi,
				GeoLocation: tt.fields.GeoLocation,
				SnBsUuid:    tt.fields.SnBsUuid,
				SnBsOpId:    tt.fields.SnBsOpId,
				SnScOpId:    tt.fields.SnScOpId,
			}
			if got := m.GetEui(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Con.GetEui() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCon_IntoProto(t *testing.T) {

	testVendor := "Test Vendor"
	testModel := "Test Model"
	testVersion := "1.0.0"
	testBsName := "M0007327767F3"
	testSwVersion := "1.2.3"

	testBsEui := common.EUI64{0x00, 0x07, 0x32, 0x00, 0x00, 0x77, 0x67, 0xF3}
	testBsSessionUuid := structs.SessionUuid{-8, -42, -98, -118, -87, -35, 70, -44, -71, 117, 17, -42, 84, 17, 74, 31}

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
		Command     structs.Command
		OpId        int64
		Version     string
		BsEui       common.EUI64
		Vendor      *string
		Model       *string
		Name        *string
		SwVersion   *string
		Info        map[string]any
		Bidi        bool
		GeoLocation *GeoLocation
		SnBsUuid    structs.SessionUuid
		SnBsOpId    *int64
		SnScOpId    *int64
	}
	tests := []struct {
		name   string
		fields fields
		want   *msg.ProtoBasestationMessage
	}{
		{
			name: "con1",
			fields: fields{
				Command:   structs.MsgCon,
				OpId:      0,
				Version:   testVersion,
				BsEui:     testBsEui,
				Vendor:    &testVendor,
				Model:     &testModel,
				Name:      &testBsName,
				SwVersion: &testSwVersion,
				SnBsUuid:  testBsSessionUuid,
				Bidi:      true,
			},
			want: &msg.ProtoBasestationMessage{
				BsEui: testBsEui.String(),
				V1: &msg.ProtoBasestationMessage_Con{
					Con: &msg.BasestationConnection{
						Ts:          &testTs,
						Version:     testVersion,
						Bidi:        true,
						Vendor:      &testVendor,
						Model:       &testModel,
						Name:        &testBsName,
						SwVersion:   &testSwVersion,
					},
				},
			},
		},
		{
			name: "con2",
			fields: fields{
				Command:     structs.MsgCon,
				OpId:        0,
				Version:     testVersion,
				BsEui:       testBsEui,
				Vendor:      &testVendor,
				Model:       &testModel,
				Name:        &testBsName,
				SwVersion:   &testSwVersion,
				SnBsUuid:    testBsSessionUuid,
				Bidi:        true,
				GeoLocation: &GeoLocation{1, 2, 3},
			},
			want: &msg.ProtoBasestationMessage{
				BsEui: testBsEui.String(),
				V1: &msg.ProtoBasestationMessage_Con{
					Con: &msg.BasestationConnection{
						Ts:        &testTs,
						Version:   testVersion,
						Bidi:      true,
						Vendor:    &testVendor,
						Model:     &testModel,
						Name:      &testBsName,
						SwVersion: &testSwVersion,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Con{
				Command:     tt.fields.Command,
				OpId:        tt.fields.OpId,
				Version:     tt.fields.Version,
				BsEui:       tt.fields.BsEui,
				Vendor:      tt.fields.Vendor,
				Model:       tt.fields.Model,
				Name:        tt.fields.Name,
				SwVersion:   tt.fields.SwVersion,
				Info:        tt.fields.Info,
				Bidi:        tt.fields.Bidi,
				GeoLocation: tt.fields.GeoLocation,
				SnBsUuid:    tt.fields.SnBsUuid,
				SnBsOpId:    tt.fields.SnBsOpId,
				SnScOpId:    tt.fields.SnScOpId,
			}
			if got := m.IntoProto(&common.EUI64{}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Con.IntoProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConRsp(t *testing.T) {
	vendor := "SplitStack"
	name := "SplitStack"
	model := "mioty BSSCI Adapter"
	swVersion := "1.0"

	type args struct {
		opId     int64
		version  string
		snScUuid uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want ConRsp
	}{
		{
			name: "conRsp",
			args: args{
				0,
				"1.0.0",
				uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7},
			},
			want: ConRsp{
				Command:   structs.MsgConRsp,
				OpId:      0,
				ScEui:     common.EUI64{1, 1, 1, 1, 1, 1, 1, 1},
				Version:   "1.0.0",
				SnScUuid:  structs.NewSessionUuid(uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7}),
				SnResume:  false,
				Vendor:    &vendor,
				Model:     &model,
				Name:      &name,
				SwVersion: &swVersion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConRsp(tt.args.opId, tt.args.version, tt.args.snScUuid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConRsp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConRsp_ResumeConnection(t *testing.T) {

	vendor := "SplitStack"
	name := "SplitStack"
	model := "mioty BSSCI Adapter"
	swVersion := "1.0"

	type fields struct {
		Command   structs.Command
		OpId      int64
		Version   string
		ScEui     common.EUI64
		Vendor    *string
		Model     *string
		Name      *string
		SwVersion *string
		Info      map[string]interface{}
		SnResume  bool
		SnScUuid  structs.SessionUuid
	}
	type args struct {
		snScUuid uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ConRsp
	}{
		{
			name: "conRsp",
			fields: fields{
				Command:   structs.MsgConRsp,
				OpId:      0,
				Version:   "1.0.0",
				ScEui:     common.EUI64{1, 1, 1, 1, 1, 1, 1, 1},
				Vendor:    &vendor,
				Model:     &model,
				Name:      &name,
				SwVersion: &swVersion,
				SnResume:  false,
				SnScUuid:  structs.NewSessionUuid(uuid.UUID{0, 1, 2, 3, 4, 5, 0, 0, 0, 0, 2, 3, 4, 5, 6, 7}),
			},
			args: args{
				uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7},
			},
			want: ConRsp{
				Command:   structs.MsgConRsp,
				OpId:      0,
				ScEui:     common.EUI64{1, 1, 1, 1, 1, 1, 1, 1},
				Version:   "1.0.0",
				SnScUuid:  structs.NewSessionUuid(uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7}),
				SnResume:  true,
				Vendor:    &vendor,
				Model:     &model,
				Name:      &name,
				SwVersion: &swVersion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ConRsp{
				Command:   tt.fields.Command,
				OpId:      tt.fields.OpId,
				Version:   tt.fields.Version,
				ScEui:     tt.fields.ScEui,
				Vendor:    tt.fields.Vendor,
				Model:     tt.fields.Model,
				Name:      tt.fields.Name,
				SwVersion: tt.fields.SwVersion,
				Info:      tt.fields.Info,
				SnResume:  tt.fields.SnResume,
				SnScUuid:  tt.fields.SnScUuid,
			}
			m.ResumeConnection(tt.args.snScUuid)

			if !reflect.DeepEqual(*m, tt.want) {
				t.Errorf("NewConRsp() = %v, want %v", m, tt.want)
			}

		})
	}
}

func TestConRsp_GetOpId(t *testing.T) {
	type fields struct {
		Command   structs.Command
		OpId      int64
		Version   string
		ScEui     common.EUI64
		Vendor    *string
		Model     *string
		Name      *string
		SwVersion *string
		Info      map[string]interface{}
		SnResume  bool
		SnScUuid  structs.SessionUuid
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "conRsp",
			fields: fields{
				Command: structs.MsgConRsp,
				OpId:    0,
				Version: "1.0.0",
				ScEui:   common.EUI64{1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ConRsp{
				Command:   tt.fields.Command,
				OpId:      tt.fields.OpId,
				Version:   tt.fields.Version,
				ScEui:     tt.fields.ScEui,
				Vendor:    tt.fields.Vendor,
				Model:     tt.fields.Model,
				Name:      tt.fields.Name,
				SwVersion: tt.fields.SwVersion,
				Info:      tt.fields.Info,
				SnResume:  tt.fields.SnResume,
				SnScUuid:  tt.fields.SnScUuid,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("ConRsp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConRsp_GetCommand(t *testing.T) {
	type fields struct {
		Command   structs.Command
		OpId      int64
		Version   string
		ScEui     common.EUI64
		Vendor    *string
		Model     *string
		Name      *string
		SwVersion *string
		Info      map[string]interface{}
		SnResume  bool
		SnScUuid  structs.SessionUuid
	}
	tests := []struct {
		name   string
		fields fields
		want   structs.Command
	}{
		{
			name: "conRsp",
			fields: fields{
				Command: structs.MsgConRsp,
				OpId:    0,
				Version: "1.0.0",
				ScEui:   common.EUI64{1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: structs.MsgConRsp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ConRsp{
				Command:   tt.fields.Command,
				OpId:      tt.fields.OpId,
				Version:   tt.fields.Version,
				ScEui:     tt.fields.ScEui,
				Vendor:    tt.fields.Vendor,
				Model:     tt.fields.Model,
				Name:      tt.fields.Name,
				SwVersion: tt.fields.SwVersion,
				Info:      tt.fields.Info,
				SnResume:  tt.fields.SnResume,
				SnScUuid:  tt.fields.SnScUuid,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConRsp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConCmp(t *testing.T) {
	type args struct {
		opId int64
	}
	tests := []struct {
		name string
		args args
		want ConCmp
	}{
		{
			name: "conCmp",
			args: args{
				1,
			},
			want: ConCmp{
				Command: structs.MsgConCmp,
				OpId:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConCmp(tt.args.opId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConCmp_GetOpId(t *testing.T) {
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
			name: "conCmp",
			fields: fields{
				Command: structs.MsgConCmp,
				OpId:    1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ConCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetOpId(); got != tt.want {
				t.Errorf("ConCmp.GetOpId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConCmp_GetCommand(t *testing.T) {
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
			name: "conCmp",
			fields: fields{
				Command: structs.MsgConCmp,
				OpId:    1,
			},
			want: structs.MsgConCmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ConCmp{
				Command: tt.fields.Command,
				OpId:    tt.fields.OpId,
			}
			if got := m.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConCmp.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
