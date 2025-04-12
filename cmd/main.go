package main

import (
	"mioty-bssci-adapter/cmd/cmd"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type pahoLogWrapper struct {
	ln func(...interface{})
	f  func(string, ...interface{})
}

func (d pahoLogWrapper) Println(v ...interface{}) {
	d.ln(v...)
}

func (d pahoLogWrapper) Printf(format string, v ...interface{}) {
	d.f(format, v...)
}

func enableClientLogging() {
	l := log.Logger.With().Str("module", "mqtt").Logger()
	
	paho.ERROR = pahoLogWrapper{func(v ...interface{}) {l.Error().Msgf("%v", v)}, l.Error().Msgf}
	paho.WARN = pahoLogWrapper{func(v ...interface{}) {l.Warn().Msgf("%v", v)}, l.Warn().Msgf}
	paho.CRITICAL = pahoLogWrapper{func(v ...interface{}) {l.Error().Msgf("%v", v)}, l.Error().Msgf}
}

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano

	enableClientLogging()
}

var version string // set by the compiler

func main() {
	cmd.Execute(version)
}
