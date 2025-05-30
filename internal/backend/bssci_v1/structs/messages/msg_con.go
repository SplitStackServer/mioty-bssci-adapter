package messages

import (
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
	"github.com/google/uuid"
)

//go:generate msgp

//msgp:shim common.EUI64 as:uint64 using:common.Eui64toUnsignedInt/common.Eui64FromUnsignedInt
//msgp:shim uuid.UUID as:[16]int8 using:UuidFromInt8/UuidToInt8

// Connect
//
// The connect operation is initiated by the Base Station immediately after establishing the
// network connection with the Service Center. No other operations may be started by
// either the Base Station or the Service Center until the connect operation is completed.
// The initial connect operation must use ID 0. This still applies if a previous session shall
// be resumed.
//
// Basestation -> Service Center
type Con struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// Requested protocol version, major.minor.patch
	Version string `msg:"version" json:"version"`
	// Base Station EUI64
	BsEui common.EUI64 `msg:"bsEui" json:"bsEui"`
	// Vendor of the Base Station, optional
	Vendor *string `msg:"vendor,omitempty" json:"vendor,omitempty"`
	// Model of the Base Station, optional
	Model *string `msg:"model,omitempty" json:"model,omitempty"`
	// Name of the Base Station, optional
	Name *string `msg:"name,omitempty" json:"name,omitempty"`
	// Software version, optional
	SwVersion *string `msg:"swVersion,omitempty" json:"swVersion,omitempty"`
	// Additional Base Station information object, might contain arbitrary key-value-pairs, optional
	Info map[string]any `msg:"info,omitempty" json:"info,omitempty"`
	// True if Base Station is bidirectional
	Bidi bool `msg:"bidi" json:"bidi"`
	// Geographic location [Latitude, Longitude, Altitude], optional
	GeoLocation *GeoLocation `msg:"geoLocation,omitempty" json:"geoLocation,omitempty"`
	// Base Station session UUID, must match with previous connect to resume session
	SnBsUuid structs.SessionUuid `msg:"snBsUuid" json:"snBsUuid"`
	// Minimum required known Base Station operation ID to resume previous session, optional
	SnBsOpId *int64 `msg:"snBsOpId,omitempty" json:"snBsOpId,omitempty"`
	// Maximum known Service Center operation ID to resume previous session, optional
	SnScOpId *int64 `msg:"snScOpId,omitempty" json:"snScOpId,omitempty"`
}

func (m *Con) GetOpId() int64 {
	return m.OpId
}

func (m *Con) GetCommand() structs.Command {
	return structs.MsgCon
}

func (m *Con) GetEui() common.EUI64 {
	return m.BsEui
}

// implements BasestationMessage.GetEventType()
func (m *Con) GetEventType() events.EventType {
	return events.EventTypeBsCon
}

// implements BasestationMessage.IntoProto()
func (m *Con) IntoProto(bsEui *common.EUI64) *bs.BasestationUplink {
	_ = bsEui
	var message bs.BasestationUplink

	if m != nil {
		bsEuiB := m.BsEui.String()

		now := getNow().UnixNano()
		ts := TimestampNsToProto(now)

		message = bs.BasestationUplink{
			BsEui: bsEuiB,
			Message: &bs.BasestationUplink_Con{
				Con: &bs.BasestationConnection{
					Ts:        ts,
					Version:   m.Version,
					Bidi:      m.Bidi,
					Vendor:    m.Vendor,
					Model:     m.Model,
					Name:      m.Name,
					SwVersion: m.SwVersion,
				},
			},
		}
	}

	return &message
}

// Connect response
//
// Basestation -> Service Center
type ConRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// Requested protocol version, major.minor.patch
	Version string `msg:"version" json:"version"`
	// Service Center EUI64
	ScEui common.EUI64 `msg:"scEui" json:"scEui"`
	// Vendor of the Service Center, optional
	Vendor *string `msg:"vendor,omitempty" json:"vendor,omitempty"`
	// Model of the Service Center, optional
	Model *string `msg:"model,omitempty" json:"model,omitempty"`
	// Name of the Service Center, optional
	Name *string `msg:"name,omitempty" json:"name,omitempty"`
	// Software version, optional
	SwVersion *string `msg:"swVersion,omitempty" json:"swVersion,omitempty"`
	// Additional Service Center information object, might contain arbitrary key-value-pairs, optional
	Info map[string]interface{} `msg:"info,omitempty" json:"info,omitempty"`
	// True if a previous session is resumed
	SnResume bool `msg:"snResume" json:"snResume"`
	// Service Center session UUID, must match with previous connect to resume session
	SnScUuid structs.SessionUuid `msg:"snScUuid" json:"snScUuid"`
}

func NewConRsp(opId int64, version string, snScUuid uuid.UUID) ConRsp {
	session := structs.NewSessionUuid(snScUuid)
	vendor := "SplitStack"
	name := "SplitStack"
	model := "mioty BSSCI Adapter"
	swVersion := "1.0"
	return ConRsp{
		Command:   structs.MsgConRsp,
		OpId:      opId,
		ScEui:     common.EUI64{1, 1, 1, 1, 1, 1, 1, 1},
		Version:   version,
		SnScUuid:  session,
		SnResume:  false,
		Vendor:    &vendor,
		Name:      &name,
		Model:     &model,
		SwVersion: &swVersion,
	}
}

func (m *ConRsp) ResumeConnection(snScUuid uuid.UUID) {
	session := structs.NewSessionUuid(snScUuid)
	m.SnScUuid = session
	m.SnResume = true
}

func (m *ConRsp) GetOpId() int64 {
	return m.OpId
}

func (m *ConRsp) GetCommand() structs.Command {
	return structs.MsgConRsp
}

// Connect complete
//
// Basestation -> Service Center
type ConCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewConCmp(opId int64) ConCmp {
	return ConCmp{OpId: opId, Command: structs.MsgConCmp}
}

func (m *ConCmp) GetOpId() int64 {
	return m.OpId
}

func (m *ConCmp) GetCommand() structs.Command {
	return structs.MsgConCmp
}
