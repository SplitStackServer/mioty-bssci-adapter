package forwarder

import (
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/events"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/config"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/integration"

	"github.com/SplitStackServer/splitstack/api/go/v5/bs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Setup configures the forwarder.
func Setup(conf config.Config) error {
	b := backend.GetBackend()
	i := integration.GetIntegration()

	if b == nil {
		return errors.New("backend is not set")
	}

	if i == nil {
		return errors.New("integration is not set")
	}

	// setup backend callbacks
	b.SetSubscribeEventHandler(gatewaySubscribeEventHandler)
	b.SetBasestationMessageHandler(basestationMessageHandler)
	b.SetEndnodeMessageHandler(endnodeMessageHandler)

	// setup integration callbacks
	i.SetServerCommandHandler(serverCommandHandler)
	i.SetServerResponseHandler(serverResponseHandler)

	return nil
}

func gatewaySubscribeEventHandler(pl events.Subscribe) {
	go func(pl events.Subscribe) {
		if err := integration.GetIntegration().SetBasestationSubscription(pl.Subscribe, pl.BasestationEui); err != nil {
			log.Error().Err(err).Msg("set basestation subscription error")
		}
	}(pl)
}

func basestationMessageHandler(eui common.EUI64, event events.EventType, pb *bs.BasestationUplink) {
	go func(eui common.EUI64, event events.EventType, pb *bs.BasestationUplink) {
		if err := integration.GetIntegration().PublishBasestationEvent(eui, string(event), pb); err != nil {
			log.Error().Err(err).Str("bs_eui", eui.String()).Str("event", string(event)).Msg("publish basestation event error")
		}
	}(eui, event, pb)
}

func endnodeMessageHandler(eui common.EUI64, event events.EventType, pb *bs.EndnodeUplink) {
	go func(eui common.EUI64, event events.EventType, pb *bs.EndnodeUplink) {

		if err := integration.GetIntegration().PublishEndnodeEvent(eui, string(event), pb); err != nil {
			log.Error().Err(err).Str("bs_eui", eui.String()).Str("event", string(event)).Msg("publish endnode event error")
		}
	}(eui, event, pb)
}

func serverCommandHandler(pb *bs.ServerCommand) {
	go func(pb *bs.ServerCommand) {
		if err := backend.GetBackend().HandleServerCommand(pb); err != nil {
			log.Error().Err(err).Msg("failed to handle server command")
		}
	}(pb)
}

func serverResponseHandler(pb *bs.ServerResponse) {
	go func(pb *bs.ServerResponse) {
		if err := backend.GetBackend().HandleServerResponse(pb); err != nil {
			log.Error().Err(err).Msg("failed to handle server response")
		}
	}(pb)
}
