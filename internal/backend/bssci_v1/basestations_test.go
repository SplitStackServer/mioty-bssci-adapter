package bssci_v1

import (
	"net"
	"testing"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestBasestationsSuite struct {
	suite.Suite

	eui        common.EUI64
	connection connection

	basestations basestations
}

func TestBasestations(t *testing.T) {
	suite.Run(t, new(TestBasestationsSuite))
}

func (ts *TestBasestationsSuite) SetupSuite() {}

func (ts *TestBasestationsSuite) SetupTest() {

	ts.eui = common.EUI64{1}

	serverConn, _ := net.Pipe()
	ts.connection = newConnection(serverConn, structs.NewSessionUuid(uuid.New()))

	ts.basestations = basestations{
		basestations:          map[common.EUI64]*connection{ts.eui: &ts.connection},
		subscribeEventHandler: func(events.Subscribe) {},
	}
}

func (ts *TestBasestationsSuite) TestBasestations_get() {
	t := ts.T()

	type args struct {
		eui common.EUI64
	}
	tests := []struct {
		name    string
		args    args
		want    *connection
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				eui: ts.eui,
			},
			want:    &ts.connection,
			wantErr: false,
		},
		{
			name: "not_found",
			args: args{
				eui: common.EUI64{2},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got, err := ts.basestations.get(tt.args.eui)

			if tt.wantErr {
				assert.Error(err)
			} else {
				if assert.NoError(err) != tt.wantErr {
					assert.Equal(tt.want, got)
				}
			}
		})
	}
}

func (ts *TestBasestationsSuite) TestBasestations_set() {
	t := ts.T()

	serverConn, _ := net.Pipe()
	newConnection := newConnection(serverConn, structs.NewSessionUuid(uuid.New()))

	type args struct {
		eui common.EUI64
		c   *connection
	}
	tests := []struct {
		name    string
		args    args
		want    *connection
		wantErr bool
	}{
		{
			name: "existing",
			args: args{
				eui: ts.eui,
				c:   &ts.connection,
			},
			wantErr: false,
		},
		{
			name: "overwrite",
			args: args{
				eui: ts.eui,
				c:   &newConnection,
			},
			wantErr: false,
		},
		{
			name: "new",
			args: args{
				eui: common.EUI64{2},
				c:   &newConnection,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			err := ts.basestations.set(tt.args.eui, tt.args.c)

			if tt.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
				_, err := ts.basestations.get(tt.args.eui)
				assert.NoError(err)
			}
		})
	}
}

func (ts *TestBasestationsSuite) TestBasestations_remove() {
	t := ts.T()

	type args struct {
		eui common.EUI64
	}
	tests := []struct {
		name      string
		args      args
		remaining int
	}{
		{
			name: "no-op",
			args: args{
				eui: common.EUI64{2},
			},
			remaining: 1,
		},
		{
			name: "valid (cannot fail)",
			args: args{
				eui: ts.eui,
			},
			remaining: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			err := ts.basestations.remove(tt.args.eui)
			assert.NoError(err)
			assert.Equal(len(ts.basestations.basestations), tt.remaining)

		})
	}
}
