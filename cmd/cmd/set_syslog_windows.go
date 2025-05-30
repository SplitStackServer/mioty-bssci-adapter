//go:build windows
// +build windows

package cmd

import (
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/config"
	"github.com/rs/zerolog/log"
)

func setSyslog() error {
	if !config.C.General.LogToSyslog {
		log.Fatal().Msg("syslog logging is not supported on Windows")
	}

	return nil
}
