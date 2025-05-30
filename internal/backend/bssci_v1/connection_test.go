package bssci_v1

import (
	"net"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs/messages"

	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/tinylib/msgp/msgp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestConnectionSuite struct {
	suite.Suite

	clientConn net.Conn
	serverConn net.Conn
	connection connection
	snBsUuid   uuid.UUID
}

func TestConnection(t *testing.T) {
	suite.Run(t, new(TestConnectionSuite))
}

func (ts *TestConnectionSuite) SetupSuite() {}

func (ts *TestConnectionSuite) SetupTest() {
	ts.serverConn, ts.clientConn = net.Pipe()
	ts.snBsUuid = uuid.New()
	ts.connection = newConnection(ts.serverConn, structs.NewSessionUuid(ts.snBsUuid))
}

func (ts *TestConnectionSuite) TearDownTest() {

	assert := require.New(ts.T())

	assert.NoError(ts.serverConn.Close())
	assert.NoError(ts.clientConn.Close())
}

func (ts *TestConnectionSuite) TestConnection_Write() {
	t := ts.T()

	type args struct {
		msg     messages.Message
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				msg:     &messages.Ping{Command: structs.MsgPing, OpId: -2},
				timeout: time.Second,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			go func() {
				defer ts.clientConn.Close()

				for {
					ts.clientConn.SetReadDeadline(time.Now().Add(time.Second))
					buf := make([]byte, 12)
					ts.clientConn.Read(buf)
				}
			}()

			assert := assert.New(t)
			assert.NoError(ts.connection.Write(tt.args.msg, tt.args.timeout))
		})
	}
}

func (ts *TestConnectionSuite) TestConnection_Read() {
	t := ts.T()

	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		data    []byte
		wantCmd structs.CommandHeader
		wantRaw msgp.Raw
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				timeout: time.Second,
			},
			data: []byte{77, 73, 79, 84, 89, 66, 48, 49, 20, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			wantCmd: structs.CommandHeader{
				Command: structs.MsgPing,
				OpId:    0,
			},
			wantRaw: msgp.Raw{130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			wantErr: false,
		},
		{
			name:    "error",
			data:    []byte{},
			wantCmd: structs.CommandHeader{},
			wantRaw: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			go func() {
				defer ts.clientConn.Close()
				ts.clientConn.SetWriteDeadline(time.Now().Add(tt.args.timeout))
				ts.clientConn.Write(tt.data)
			}()

			gotCmd, gotRaw, err := ts.connection.Read()

			if tt.wantErr {
				assert.Error(err)
			} else {
				if assert.NoError(err) != tt.wantErr {
					assert.Equal(tt.wantCmd, gotCmd)
					assert.Equal(tt.wantRaw, gotRaw)
				}
			}
		})
	}
}

func (ts *TestConnectionSuite) TestConnection_GetAndDecrementOpId() {
	t := ts.T()

	tests := []struct {
		name     string
		wantOpId int64
	}{
		{
			name:     "first",
			wantOpId: -1,
		},
		{
			name:     "second",
			wantOpId: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			gotOpId := ts.connection.GetAndDecrementOpId()
			assert.Equal(tt.wantOpId, gotOpId)
		})
	}
}

func (ts *TestConnectionSuite) TestConnection_ResumeConnection() {
	t := ts.T()

	var snScOpId int64 = -10

	type args struct {
		snBsUuid uuid.UUID
		snScOpId *int64
	}
	tests := []struct {
		name            string
		args            args
		wantResume      bool
		wantNewSnScUuid bool
	}{
		{
			name: "resume",
			args: args{
				snBsUuid: ts.snBsUuid,
				snScOpId: nil,
			},
			wantResume: true,
		},
		{
			name: "resume_with_opId",
			args: args{
				snBsUuid: ts.snBsUuid,
				snScOpId: &snScOpId,
			},
			wantResume: true,
		},
		{
			name: "no_resume",
			args: args{
				snBsUuid: uuid.New(),
				snScOpId: nil,
			},
			wantResume: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			currentSnScUuid := ts.connection.SnScUuid

			resume, _ := ts.connection.ResumeConnection(tt.args.snBsUuid, tt.args.snScOpId)
			if assert.Equal(tt.wantResume, resume) {
				if resume {
					assert.Equal(currentSnScUuid, ts.connection.SnScUuid)
				} else {
					assert.NotEqual(currentSnScUuid, ts.connection.SnScUuid)
				}

			}

		})
	}
}
