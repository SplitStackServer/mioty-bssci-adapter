package bssci_v1

import (
	"bytes"
	"encoding/hex"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs/messages"
	"reflect"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestHex(t *testing.T) {
	// t.Skip()
	s := "8da7636f6d6d616e64a9766d2e756c44617461a46f70496404a76d61635479706501a87573657244617461dc00565544cca511465279cc817607cc8c00ccb7cc900f002c253e351100cc98cccaccecccc1124225707a1bcc933207102a44ccfbcca948ccf1ccd8ccfa1879cc87cca0cc920bcc89cc8d38ccde03ccfacc9fcce876cce6502dccb5cca4672f5fcc843535225d31cc9b68cc9f555721ccb6cc902b7746cc8a6ca774727854696d65cf0009533f0d4cc1c7a773797354696d65cf183555b86b09db8ea7667265714f6666cb41c9df844e000000a472737369cbc0601c8880000000a3736e72cb400a2c0bc0000000a963617272537061636501a77061747447727000a7706174744e756d03a36372639257cc8e"

	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	t.Errorf("%v", data)

}

func TestReadBssciMessage(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantCmd structs.CommandHeader
		wantRaw msgp.Raw
		wantErr bool
	}{
		{
			name: "valid",
			data: []byte{77, 73, 79, 84, 89, 66, 48, 49, 20, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			wantCmd: structs.CommandHeader{
				Command: structs.MsgPing,
				OpId:    0,
			},
			wantRaw: msgp.Raw{130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			wantErr: false,
		},
		{
			name:    "io read error on header",
			data:    []byte{},
			wantCmd: structs.CommandHeader{},
			wantRaw: nil,
			wantErr: true,
		},
		{
			name:    "header error",
			data:    []byte{77, 73, 79, 84, 89, 66, 48, 21, 20, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			wantCmd: structs.CommandHeader{},
			wantRaw: nil,
			wantErr: true,
		},
		{
			name:    "io read error on message",
			data:    []byte{77, 73, 79, 84, 89, 66, 48, 49, 1, 0, 0, 0},
			wantCmd: structs.CommandHeader{},
			wantRaw: nil,
			wantErr: true,
		},
		{
			name:    "command error",
			data:    []byte{77, 73, 79, 84, 89, 66, 48, 49, 20, 0, 0, 0, 130},
			wantCmd: structs.CommandHeader{},
			wantRaw: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			reader := bytes.NewReader(tt.data)

			gotCmd, gotRaw, err := ReadBssciMessage(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBssciMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCmd, tt.wantCmd) {
				t.Errorf("ReadBssciMessage() gotCmd = %v, want %v", gotCmd, tt.wantCmd)
			}
			if !reflect.DeepEqual(gotRaw, tt.wantRaw) {
				t.Errorf("ReadBssciMessage() gotRaw = %v, want %v", gotRaw, tt.wantRaw)
			}
		})
	}
}

func TestWriteBssciMessage(t *testing.T) {
	type args struct {
		msg messages.Message
	}
	tests := []struct {
		name    string
		args    args
		wantB   []byte
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				msg: &messages.Ping{
					Command: structs.MsgPing,
					OpId:    0,
				},
			},
			wantB:   []byte{77, 73, 79, 84, 89, 66, 48, 49, 20, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := &bytes.Buffer{}
			if err := WriteBssciMessage(w, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("WriteBssciMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(w.Bytes(), tt.wantB) {
				t.Errorf("WriteBssciMessage() = %v, want %v", w.Bytes(), tt.wantB)
			}
		})
	}
}

func TestMarshalBssciMessage(t *testing.T) {
	type args struct {
		msg messages.Message
	}
	tests := []struct {
		name    string
		args    args
		wantBuf []byte
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				msg: &messages.Ping{
					Command: structs.MsgPing,
					OpId:    0,
				},
			},
			wantBuf: []byte{77, 73, 79, 84, 89, 66, 48, 49, 20, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBuf, err := MarshalBssciMessage(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalBssciMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBuf, tt.wantBuf) {
				t.Errorf("MarshalBssciMessage() = %v, want %v", gotBuf, tt.wantBuf)
			}
		})
	}
}

func TestUnmarshalBssciMessage(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		args    args
		wantCmd structs.CommandHeader
		wantRaw msgp.Raw
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				buf: []byte{77, 73, 79, 84, 89, 66, 48, 49, 20, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			},
			wantCmd: structs.CommandHeader{
				Command: structs.MsgPing,
				OpId:    0,
			},
			wantRaw: msgp.Raw{130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			wantErr: false,
		},
		{
			name: "header_error",
			args: args{
				buf: []byte{77, 73, 79, 84, 89, 66, 48, 55, 20, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			},
			wantCmd: structs.CommandHeader{},
			wantRaw: nil,
			wantErr: true,
		},
		{
			name: "command_error",
			args: args{
				buf: []byte{77, 73, 79, 84, 89, 66, 48, 49, 20, 0, 0, 0, 130, 162, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			},
			wantCmd: structs.CommandHeader{},
			wantRaw: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmd, gotRaw, err := UnmarshalBssciMessage(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalBssciMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCmd, tt.wantCmd) {
				t.Errorf("UnmarshalBssciMessage() gotCmd = %v, want %v", gotCmd, tt.wantCmd)
			}
			if !reflect.DeepEqual(gotRaw, tt.wantRaw) {
				t.Errorf("UnmarshalBssciMessage() gotRaw = %v, want %v", gotRaw, tt.wantRaw)
			}
		})
	}
}

func Test_prepareBssciMessage(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		wantBuf []byte
	}{
		{
			name: "valid_header",
			args: args{
				length: 0,
			},
			wantBuf: []byte{77, 73, 79, 84, 89, 66, 48, 49, 0, 0, 0, 0},
		},
		{
			name: "valid_header",
			args: args{
				length: 10,
			},
			wantBuf: []byte{77, 73, 79, 84, 89, 66, 48, 49, 10, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBuf := prepareBssciMessage(tt.args.length)

			if !reflect.DeepEqual(gotBuf, tt.wantBuf) {
				t.Errorf("prepareBssciMessage() = %v, want %v", gotBuf, tt.wantBuf)
			}
		})
	}
}

func Test_getBssciMessageLengthFromHeader(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		args    args
		wantL   int32
		wantErr bool
	}{
		{
			name: "valid_header",
			args: args{
				buf: []byte{77, 73, 79, 84, 89, 66, 48, 49, 20, 0, 0, 0},
			},
			wantL:   20,
			wantErr: false,
		},
		{
			name: "invalid_header_size",
			args: args{
				buf: []byte{77, 73, 79, 84, 89, 66, 48, 49, 20},
			},
			wantL:   0,
			wantErr: true,
		},
		{
			name: "invalid_identifier",
			args: args{
				buf: []byte{77, 73, 81, 84, 22, 66, 48, 49, 20, 0, 0, 0},
			},
			wantL:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotL, err := getBssciMessageLengthFromHeader(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBssciMessageLengthFromHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotL != tt.wantL {
				t.Errorf("getBssciMessageLengthFromHeader() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
