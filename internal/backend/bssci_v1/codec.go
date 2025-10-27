package bssci_v1

import (
	// "encoding/binary"
	"io"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs/messages"

	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
)

const (
	bssciHeaderSize           = 12
	bssciHeaderIdentifierSize = 8
	bssciHeaderLengthOffset   = 8
	bssciHeaderLengthSize     = 4
)

var (
	bssciIdentifier = [8]byte{0x4D, 0x49, 0x4F, 0x54, 0x59, 0x42, 0x30, 0x31}
)

func ReadBssciMessage(r io.Reader) (cmd structs.CommandHeader, raw msgp.Raw, err error) {
	// reader the bssci header
	buf := make([]byte, bssciHeaderSize)
	_, err = r.Read(buf)

	if err != nil {
		err = errors.Wrap(err, "io read error on header")
		return
	}

	// get the size of the message
	length, err := getBssciMessageLengthFromHeader(buf)
	if err != nil {
		return
	}
	buf = make([]byte, length)
	_, err = r.Read(buf)

	if err != nil {
		err = errors.Wrap(err, "io read error on message")
		return
	}

	// parse out command
	_, err = cmd.UnmarshalMsg(buf)
	if err != nil {
		err = errors.Wrap(err, "command error")
		return
	}

	// get raw message
	raw.UnmarshalMsg(buf)

	return
}

func WriteBssciMessage(w io.Writer, msg messages.MessageMsgp) (err error) {
	// marshal message
	buf, err := MarshalBssciMessage(msg)
	if err != nil {
		return
	}

	// write message
	_, err = w.Write(buf)

	if err != nil {
		err = errors.Wrap(err, "io write error")
		return
	}

	return
}

// convert msg to message pack and attach bssci header
func MarshalBssciMessage(msg messages.MessageMsgp) ([]byte, error) {
	msgBuf, err := msg.MarshalMsg(nil)
	if err != nil {
		err = errors.Wrap(err, "message marshal error")
		return nil, err
	}

	msgLength := len(msgBuf)

	buf := prepareBssciMessage(msgLength)

	// add message pack data
	buf = append(buf, msgBuf...)
	return buf, err
}

// read the bssci header and extract the raw message pack data
func UnmarshalBssciMessage(buf []byte) (cmd structs.CommandHeader, raw msgp.Raw, err error) {

	// get length
	header_buf := buf[:bssciHeaderSize]

	length, err := getBssciMessageLengthFromHeader(header_buf)
	if err != nil {
		err = errors.Wrap(err, "header error")
		return
	}

	// slice off header
	buf = buf[bssciHeaderSize : bssciHeaderSize+length]

	// parse out command
	_, err = cmd.UnmarshalMsg(buf)
	if err != nil {
		err = errors.Wrap(err, "command error")
		return
	}

	// read rest of the message
	raw.UnmarshalMsg(buf)

	return
}

// build header and allocate buffer
func prepareBssciMessage(length int) []byte {
	// allocate buf
	buf := make([]byte, 0, bssciHeaderSize+length)
	// add identifier and length to header
	buf = append(buf, bssciIdentifier[:]...)
	// add length to header
	length_buf := []byte{
		byte(0xff & length),
		byte(0xff & (length >> 8)),
		byte(0xff & (length >> 16)),
		byte(0xff & (length >> 24)),
	}
	buf = append(buf, length_buf...)

	return buf
}

func getBssciMessageLengthFromHeader(buf []byte) (l int32, err error) {

	if len(buf) != bssciHeaderSize {
		err = errors.Errorf("invalid header size: %v", len(buf))
		return
	}

	identifier := [8]byte(buf[:bssciHeaderIdentifierSize])
	if identifier != bssciIdentifier {
		err = errors.Errorf("message header error: invalid identifier in buffer %v", buf)
		return
	}

	length_buf := [4]byte(buf[bssciHeaderLengthOffset : bssciHeaderLengthOffset+bssciHeaderLengthSize])
	l = int32(length_buf[0]) | int32(length_buf[1])<<8 | int32(length_buf[2])<<16 | int32(length_buf[3])<<24

	return
}
