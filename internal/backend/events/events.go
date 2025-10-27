package events

import "github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

type EventType string

const (
	EventTypeBsStatus   EventType = "status"
	EventTypeBsCon      EventType = "con"
	EventTypeBsVmStatus EventType = "vm"
	EventTypeBsDl       EventType = "dl"
	EventTypeBsPrpAck   EventType = "prp_ack"
	EventTypeEpOtaa     EventType = "otaa"
	EventTypeEpUl       EventType = "ul"
	EventTypeEpRx       EventType = "rx"
)

// Subscribe event
type Subscribe struct {
	// Basestation EUI64.
	BasestationEui common.EUI64

	// Subscribe (true) or unsubscribe (false) the gateway.
	Subscribe bool
}
