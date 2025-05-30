package integration

import (
	"context"

	"github.com/pkg/errors"

	"github.com/SplitStackServer/splitstack/api/go/v4/bs"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/config"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/integration/mqtt"
)

// Event types.
const (
	EventUp    = "up"
	EventStats = "stats"
	EventAck   = "ack"
	EventRaw   = "raw"
)

var integration Integration

// Setup configures the integration.
func Setup(conf config.Config) error {
	var err error
	integration, err = mqtt.NewIntegration(conf)
	if err != nil {
		return errors.Wrap(err, "setup mqtt integration error")
	}

	return nil
}

// GetIntegration returns the integration.
func GetIntegration() Integration {
	return integration
}

// Integration defines the interface that an integration must implement.
type Integration interface {
	// Updates the subscription for the given EUI.
	SetBasestationSubscription(subscribe bool, bsEui common.EUI64) error

	// Publishes the current state as retained message.
	PublishState(ctx context.Context, bsEui common.EUI64, pb *bs.BasestationState) error

	// Publish endnode messages.
	PublishEndnodeEvent(bsEui common.EUI64, event string, pb *bs.EndnodeUplink) error

	// Publish basestation messages.
	PublishBasestationEvent(bsEui common.EUI64, event string, pb *bs.BasestationUplink) error

	// Set handler for server command messages
	SetServerCommandHandler(func(*bs.ServerCommand))

	// Set handler for server command messages
	SetServerResponseHandler(func(*bs.ServerResponse))

	// Start starts the integration.
	Start() error

	// Stop stops the integration.
	Stop() error
}
