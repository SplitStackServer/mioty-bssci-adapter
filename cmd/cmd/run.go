package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"mioty-bssci-adapter/internal/backend"
	"mioty-bssci-adapter/internal/config"
	"mioty-bssci-adapter/internal/forwarder"
	"mioty-bssci-adapter/internal/integration"
	"mioty-bssci-adapter/internal/metrics"
)


func run(cmd *cobra.Command, args []string) error {

	tasks := []func() error{
		setLogLevel,
		setSyslog,
		printStartMessage,
		setupBackend,
		setupIntegration,
		setupForwarder,
		setupMetrics,
		startIntegration,
		startBackend,
	}

	for _, t := range tasks {
		if err := t(); err != nil {
			log.Fatal().Err(err).Msg("error during startup")
		}
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.Info().Any("signal", <-sigChan).Msg("signal received")
	log.Warn().Msg("shutting down server")

	integration.GetIntegration().Stop()

	return nil
}


func setLogLevel() error {
	log.Logger = log.Logger.With().Caller().Logger()
	zerolog.SetGlobalLevel(zerolog.Level(uint8(config.C.General.LogLevel)))
	return nil
}

func printStartMessage() error {
	log.Info().Any("config", config.C).Msg("starting server")

	return nil
}

func setupBackend() error {
	if err := backend.Setup(config.C); err != nil {
		return errors.Wrap(err, "setup backend error")
	}
	return nil
}

func setupIntegration() error {
	if err := integration.Setup(config.C); err != nil {
		return errors.Wrap(err, "setup integration error")
	}
	return nil
}

func setupForwarder() error {
	if err := forwarder.Setup(config.C); err != nil {
		return errors.Wrap(err, "setup forwarder error")
	}
	return nil
}

func setupMetrics() error {
	if err := metrics.Setup(config.C); err != nil {
		return errors.Wrap(err, "setup metrics error")
	}
	return nil
}

func startIntegration() error {
	if err := integration.GetIntegration().Start(); err != nil {
		return errors.Wrap(err, "start integration error")
	}
	return nil
}

func startBackend() error {
	if err := backend.GetBackend().Start(); err != nil {
		return errors.Wrap(err, "start backend error")
	}
	return nil
}
