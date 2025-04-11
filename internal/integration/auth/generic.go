package auth

import (
	"crypto/tls"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"

	"github.com/rs/zerolog/log"

	"mioty-bssci-adapter/internal/common"
	"mioty-bssci-adapter/internal/config"
)

// GenericAuthentication implements a generic MQTT authentication.
type GenericAuthentication struct {
	servers      []string
	username     string
	password     string
	cleanSession bool
	clientID     string

	tlsConfig *tls.Config
}

// NewGenericAuthentication creates a GenericAuthentication.
func NewGenericAuthentication(conf config.Config) (Authentication, error) {
	tlsConfig, err := newTLSConfig(
		conf.Integration.MQTTV3.Auth.Generic.CACert,
		conf.Integration.MQTTV3.Auth.Generic.TLSCert,
		conf.Integration.MQTTV3.Auth.Generic.TLSKey,
	)
	if err != nil {
		return nil, errors.Wrap(err, "mqtt/auth: new tls config error")
	}

	return &GenericAuthentication{
		tlsConfig:    tlsConfig,
		servers:      conf.Integration.MQTTV3.Auth.Generic.Servers,
		username:     conf.Integration.MQTTV3.Auth.Generic.Username,
		password:     conf.Integration.MQTTV3.Auth.Generic.Password,
		cleanSession: conf.Integration.MQTTV3.Auth.Generic.CleanSession,
		clientID:     conf.Integration.MQTTV3.Auth.Generic.ClientID,
	}, nil
}

// Init applies the initial configuration.
func (a *GenericAuthentication) Init(opts *mqtt.ClientOptions) error {
	for _, server := range a.servers {
		opts.AddBroker(server)
	}
	opts.SetUsername(a.username)
	opts.SetPassword(a.password)
	opts.SetCleanSession(a.cleanSession)
	opts.SetClientID(a.clientID)

	if a.tlsConfig != nil {
		opts.SetTLSConfig(a.tlsConfig)
	}

	return nil
}

// Return the basestation EUI64 if available
func (a *GenericAuthentication) GetBasestationEui() *common.EUI64 {
	if a.clientID == "" {
		return nil
	}

	// Try to decode the client ID as gateway ID.
	var gatewayID common.EUI64
	if err := gatewayID.UnmarshalText([]byte(a.clientID)); err != nil {
		log.Warn().Err(err).Str("client", a.clientID).Msg("could not decode client ID to gateway ID")

		return nil
	}

	return &gatewayID
}

// Update updates the authentication options.
func (a *GenericAuthentication) Update(opts *mqtt.ClientOptions) error {
	return nil
}

// ReconnectAfter returns a time.Duration after which the MQTT client must re-connect.
// Note: return 0 to disable the periodical re-connect feature.
func (a *GenericAuthentication) ReconnectAfter() time.Duration {
	return 0
}
