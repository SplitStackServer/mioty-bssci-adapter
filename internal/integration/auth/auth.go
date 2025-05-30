package auth

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"
)

// Authentication defines the authentication interface.
type Authentication interface {
	// Init applies the initial configuration.
	Init(*mqtt.ClientOptions) error

	// Return the basestation EUI64 if available
	GetBasestationEui() *common.EUI64

	// Update updates the authentication options.
	Update(*mqtt.ClientOptions) error

	// ReconnectAfter returns a time.Duration after which the MQTT client must re-connect.
	// Note: return 0 to disable the periodical re-connect feature.
	ReconnectAfter() time.Duration
}

func newTLSConfig(cafile string, certFile string, certKeyFile string) (*tls.Config, error) {
	if cafile == "" && certFile == "" && certKeyFile == "" {
		return nil, nil
	}

	tlsConfig := &tls.Config{}

	if cafile != "" {
		rawCACert, err := os.ReadFile(cafile)
		if err != nil {
			return nil, errors.Wrap(err, "read ca cert error")
		}
		certpool := x509.NewCertPool()
		certpool.AppendCertsFromPEM(rawCACert)

		tlsConfig.RootCAs = certpool // RootCAs = certs used to verify server cert.
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	}

	if certFile != "" && certKeyFile != "" {
		tlsCert, err := tls.LoadX509KeyPair(certFile, certKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "read tls cert error")
		}
		tlsConfig.Certificates = []tls.Certificate{tlsCert}
	}

	return tlsConfig, nil
}
