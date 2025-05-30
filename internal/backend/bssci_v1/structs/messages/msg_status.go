package messages

import (
	"errors"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"
)

//go:generate msgp
//msgp:tuple common.GeoLocation

// Status
//
// The status operation can be initiated by the Service Center to retrieve status
// information from the Base Station.
//
// Service Center -> Basestation
type Status struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewStatus(opId int64) Status {
	return Status{OpId: opId, Command: structs.MsgStatus}
}

func NewStatusFromProto(opId int64, pb *bs.RequestStatus) (*Status, error) {
	if pb != nil {
		m := NewStatus(opId)
		return &m, nil
	}
	return nil, errors.New("invalid RequestStatus command")
}

func (m *Status) GetOpId() int64 {
	return m.OpId
}

func (m *Status) GetCommand() structs.Command {
	return structs.MsgStatus
}

// implements ServerMessage
func (m *Status) SetOpId(opId int64) {
	m.OpId = opId
}

// Status response
//
// Basestation -> Service Center
type StatusRsp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
	// Status code, using POSIX error numbers, 0 for ok
	Code uint32 `msg:"code" json:"code"`
	// Status message
	Message string `msg:"message" json:"message"`
	// Unix UTC system time, 64 bit, ns resolution
	Time uint64 `msg:"time" json:"time"`
	// Fraction of TX time, sliding window over one hour
	DutyCycle float32 `msg:"dutyCycle" json:"dutyCycle"`
	// Geographic location [Latitude, Longitude, Altitude], optional
	GeoLocation *GeoLocation `msg:"geoLocation,omitempty" json:"geoLocation,omitempty"`
	// System uptime in seconds, optional
	Uptime *uint64 `msg:"uptime,omitempty" json:"uptime,omitempty"`
	// System temperature in degree Celsius, optional
	Temp *float64 `msg:"temp,omitempty" json:"temp,omitempty"`
	// CPU utilization, normalized to 1.0 for all cores, optional
	CpuLoad *float64 `msg:"cpuLoad,omitempty" json:"cpuLoad,omitempty"`
	// Memory utilization, normalized to 1.0, optional
	MemLoad *float64 `msg:"memLoad,omitempty" json:"memLoad,omitempty"`
}

func NewStatusRsp(opId int64, code uint32, message string, time uint64, dutyCycle float32, geoLocation *GeoLocation, uptime *uint64, temp *float64, cpuLoad *float64, memLoad *float64) StatusRsp {
	return StatusRsp{
		Command:     structs.MsgStatusRsp,
		OpId:        opId,
		Code:        code,
		Message:     message,
		Time:        time,
		DutyCycle:   dutyCycle,
		GeoLocation: geoLocation,
		Uptime:      uptime,
		Temp:        temp,
		CpuLoad:     cpuLoad,
		MemLoad:     memLoad,
	}
}

func (m *StatusRsp) GetOpId() int64 {
	return m.OpId
}

func (m *StatusRsp) GetCommand() structs.Command {
	return structs.MsgStatusRsp
}

// implements BasestationMessage.GetEventType()
func (m *StatusRsp) GetEventType() events.EventType {
	return events.EventTypeBsStatus
}

// implements BasestationMessage.IntoProto()
func (m *StatusRsp) IntoProto(bsEui *common.EUI64) *bs.BasestationUplink {

	var message bs.BasestationUplink

	if m != nil && bsEui != nil {
		bsEuiB := bsEui.String()
		ts := TimestampNsToProto(int64(m.Time))

		message = bs.BasestationUplink{
			BsEui: bsEuiB,
			Message: &bs.BasestationUplink_Status{
				Status: &bs.BasestationStatus{
					StatusCode:  m.Code,
					StatusMsg:   m.Message,
					Ts:          ts,
					DutyCycle:   m.DutyCycle,
					Uptime:      m.Uptime,
					Temp:        m.Temp,
					Cpu:         m.CpuLoad,
					Memory:      m.MemLoad,
					GeoLocation: m.GeoLocation.IntoProto(),
				},
			},
		}
	}

	return &message
}

// Status complete
//
// Service Center -> Basestation
type StatusCmp struct {
	Command structs.Command `msg:"command" json:"command"`
	// ID of the operation
	OpId int64 `msg:"opId" json:"opId"`
}

func NewStatusCmp(opId int64) StatusCmp {
	return StatusCmp{OpId: opId, Command: structs.MsgStatusCmp}
}

func (m *StatusCmp) GetOpId() int64 {
	return m.OpId
}

func (m *StatusCmp) GetCommand() structs.Command {
	return structs.MsgStatusCmp
}
