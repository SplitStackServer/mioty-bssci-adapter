package common

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// EUI64 data type
type EUI64 [8]byte

// helper function to parse a int64 formatted eui64
func Eui64FromInt(in int64) EUI64 {
	b := [8]byte{
		byte(0xff & (in >> 56)),
		byte(0xff & (in >> 48)),
		byte(0xff & (in >> 40)),
		byte(0xff & (in >> 32)),
		byte(0xff & (in >> 24)),
		byte(0xff & (in >> 16)),
		byte(0xff & (in >> 8)),
		byte(0xff & in),
	}

	return b
}

// helper function to parse a int64 formatted eui64
func Eui64FromUnsignedInt(in uint64) EUI64 {
	b := [8]byte{
		byte(0xff & (in >> 56)),
		byte(0xff & (in >> 48)),
		byte(0xff & (in >> 40)),
		byte(0xff & (in >> 32)),
		byte(0xff & (in >> 24)),
		byte(0xff & (in >> 16)),
		byte(0xff & (in >> 8)),
		byte(0xff & in),
	}

	return b
}

// helper function to parse a int64 formatted eui64
func Eui64FromHexString(in string) (EUI64, error) {
	var eui EUI64

	b, err := hex.DecodeString(strings.TrimPrefix(in, "0x"))
	if err != nil {
		return eui, err
	}

	copy(eui[:], b)
	return eui, nil

}


// helper function to parse eui64 into a int64
func (e *EUI64) ToInt() int64 {
	return int64(e[7]) |
		int64(e[6])<<8 |
		int64(e[5])<<16 |
		int64(e[4])<<24 |
		int64(e[3])<<32 |
		int64(e[2])<<40 |
		int64(e[1])<<48 |
		int64(e[0])<<56
}

// helper function to parse eui64 into a uint64
func (e *EUI64) ToUnsignedInt() uint64 {
	return uint64(e[7]) |
		uint64(e[6])<<8 |
		uint64(e[5])<<16 |
		uint64(e[4])<<24 |
		uint64(e[3])<<32 |
		uint64(e[2])<<40 |
		uint64(e[1])<<48 |
		uint64(e[0])<<56
}

// String implement fmt.Stringer.
func (e EUI64) String() string {
	return hex.EncodeToString(e[:])
}

// MarshalText implements encoding.TextMarshaler.
func (e EUI64) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *EUI64) UnmarshalText(text []byte) error {
	b, err := hex.DecodeString(strings.TrimPrefix(string(text), "0x"))
	if err != nil {
		return err
	}
	if len(e) != len(b) {
		return fmt.Errorf("eui64: exactly %d bytes are expected", len(e))
	}
	copy(e[:], b)
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (e EUI64) MarshalBinary() ([]byte, error) {
	out := make([]byte, len(e))
	copy(out, e[:])
	return out, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (e *EUI64) UnmarshalBinary(data []byte) error {
	if len(data) != len(e) {
		return fmt.Errorf("eui64: exactly %d bytes are expected", len(e))
	}
	copy(e[:], data )
	return nil
}

// helper function to parse eui64 into a int64
func Eui64toInt(e EUI64) int64 {
	return e.ToInt()
}

// helper function to parse eui64 into a int64
func Eui64toUnsignedInt(e EUI64) uint64 {
	return e.ToUnsignedInt()
}
