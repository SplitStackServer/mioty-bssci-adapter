//go:build !windows
// +build !windows

package cmd

import (
	"log/syslog"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func setSyslog() error {
	if !config.C.General.LogToSyslog {
		return nil
	}

	zsyslog, err := syslog.New(syslog.LOG_EMERG|syslog.LOG_ERR|syslog.LOG_INFO|syslog.LOG_CRIT|syslog.LOG_WARNING|syslog.LOG_NOTICE|syslog.LOG_DEBUG, "mioty-bssci-adapter")
	if err != nil {
		panic(err)
	}

	log.Logger = zerolog.New(zsyslog).With().Caller().Logger()

	return nil
}
