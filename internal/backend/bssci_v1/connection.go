package bssci_v1

import (
	"net"
	"sync"
	"time"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs/messages"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
)

type connection struct {
	sync.RWMutex
	conn net.Conn
	opId int64
	// Base Station session UUID, used to resume session
	SnBsUuid uuid.UUID
	// Service Center session UUID, used to resume session
	SnScUuid uuid.UUID
}

func newConnection(conn net.Conn, snBsUuid structs.SessionUuid) connection {
	snScUuid := uuid.New()

	conn.SetReadDeadline(time.Time{})

	return connection{
		conn: conn,
		// stats:      stats.NewCollector(),
		opId:     -1,
		SnBsUuid: snBsUuid.ToUuid(),
		SnScUuid: snScUuid,
	}
}

// Send the message to this connection
func (conn *connection) Write(msg messages.Message, timeout time.Duration) (err error) {
	conn.Lock()
	defer conn.Unlock()

	bb, err := MarshalBssciMessage(msg)

	if err != nil {
		return errors.Wrap(err, "marshal msgp error")
	}

	conn.conn.SetWriteDeadline(time.Now().Add(timeout))
	_, err = conn.conn.Write(bb)
	if err != nil {
		// conn.conn.Close()
		return errors.Wrap(err, "write error")
	}

	return
}

// Read a message from this connection
func (conn *connection) Read() (cmd structs.CommandHeader, raw msgp.Raw, err error) {
	// conn.Lock()
	// defer conn.Unlock()

	cmd, raw, err = ReadBssciMessage(conn.conn)
	if err != nil {
		return
	}

	return
}

// Should be called when a message chain is initialized by the server.
//
// returns the current opId before decrement by 1
func (conn *connection) GetAndDecrementOpId() (opId int64) {
	conn.Lock()
	defer conn.Unlock()

	opId = conn.opId
	conn.opId = conn.opId - 1
	return
}

// Check if this connection is resumed after a Con message
//
// returns true and the current snScUuid if the connection is resumable, else false and a new snScUuid
func (conn *connection) ResumeConnection(snBsUuid uuid.UUID, snScOpId *int64) (resume bool, snScUuid uuid.UUID) {
	conn.Lock()
	defer conn.Unlock()

	conn.opId = -1

	if conn.SnBsUuid == snBsUuid {
		if snScOpId != nil {
			conn.opId = *snScOpId
		}
		return true, conn.SnScUuid
	}
	snScUuid = uuid.New()
	conn.SnScUuid = snScUuid

	return false, snScUuid

}
