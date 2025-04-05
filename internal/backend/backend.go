package backend

import (
	"fmt"
	// "mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/api/rsp"
	"mioty-bssci-adapter/internal/backend/bssci_v1"
	"mioty-bssci-adapter/internal/backend/events"
	"mioty-bssci-adapter/internal/common"
	"mioty-bssci-adapter/internal/config"

	"github.com/pkg/errors"
)

var backend Backend

// Setup configures the backend.
func Setup(conf config.Config) error {
	var err error

	switch conf.Backend.Type {
	case "bssci_v1":
		backend, err = bssci_v1.NewBackend(conf)
	default:
		return fmt.Errorf("unknown backend type: %s", conf.Backend.Type)
	}

	if err != nil {
		return errors.Wrap(err, "new backend error")
	}

	return nil
}

// GetBackend returns the backend.
func GetBackend() Backend {
	return backend
}

// Backend defines the interface that a backend must implement
type Backend interface {
	// Stop closes the backend.
	Stop() error

	// Start starts the backend.
	Start() error

	// Set handler for Subscribe events.
	SetSubscribeEventHandler(func(events.Subscribe))

	// Set handler messages from basestations
	SetBasestationMessageHandler(func(common.EUI64, events.EventType, *msg.ProtoBasestationMessage))

	// Set handler for messages from endnodes
	SetEndnodeMessageHandler(func(common.EUI64, events.EventType, *msg.ProtoEndnodeMessage))

	// Handler for server command messages
	HandleServerCommand(*cmd.ProtoCommand) error

	// Handler for server response messages
	HandleServerResponse(*rsp.ProtoResponse) error
}
