package cmd

import (
	"html/template"
	"os"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/config"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// when updating this template, don't forget to update config.md!
const configTemplate = `[general]
# debug=0, info=1, warn=2, error=3, fatal=4, panic=5, disabled=7
log_level={{ .General.LogLevel }}

# Log to syslog.
#
# When set to true, log messages are being written to syslog.
log_to_syslog={{ .General.LogToSyslog }}




# Gateway backend configuration.
[backend]

# Backend type.
#
# Valid options are:
#   * bssci_v1
type="{{ .Backend.Type }}"

  # BSSCI V1 backend.
  [backend.bssci_v1]

  # ip:port to bind the TCP listener to.
  bind="{{ .Backend.BssciV1.Bind }}"

  # TLS certificate and key files.
  #
  # When set, the TCP listener will use TLS to secure the connections
  # between the gateways and mioty BSSCI Adapter (optional).
  tls_cert="{{ .Backend.BssciV1.TLSCert }}"
  tls_key="{{ .Backend.BssciV1.TLSKey }}"

  # TLS CA certificate.
  #
  # When configured, mioty BSSCI Adapter will validate that the client
  # certificate of the gateway has been signed by this CA certificate.
  ca_cert="{{ .Backend.BssciV1.CACert }}"

  # Stats interval.
  #
  # This defines the interval in which the mioty BSSCI Adapter requests status messages from the connected basestations
  #
  # Valid units are 'ms', 's', 'm', 'h'. Note that these values can be combined, e.g. '24h30m15s'.
  stats_interval="{{ .Backend.BssciV1.StatsInterval }}"

  # Ping interval.
  #
  # This defines the interval in which mioty BSSCI Adapter sends ping messages to the connected basestations
  #
  # Valid units are 'ms', 's', 'm', 'h'. Note that these values can be combined, e.g. '24h30m15s'.
  ping_interval="{{ .Backend.BssciV1.PingInterval }}"


  # Keep alive period .
  #
  # This interval must be greater than the configured ping interval.
  #
  # Valid units are 'ms', 's', 'm', 'h'. Note that these values can be combined, e.g. '24h30m15s'.
  keep_alive_period="{{ .Backend.BssciV1.KeepAlivePeriod }}"

# Integration configuration.
[integration]
# Payload marshaler.
#
# This defines how the MQTT payloads are encoded. Valid options are:
# * protobuf:  Protobuf encoding
# * json:      JSON encoding (for debugging)
marshaler="{{ .Integration.Marshaler }}"

  # MQTT integration configuration.
  [integration.mqtt_v3]

  # Keep alive will set the amount of time (in seconds) that the client should
  # wait before sending a PING request to the broker. This will allow the client
  # to know that a connection has not been lost with the server.
  # Valid units are 'ms', 's', 'm', 'h'. Note that these values can be combined, e.g. '24h30m15s'.
  keep_alive="{{ .Integration.MQTTV3.KeepAlive }}"

  # Maximum interval that will be waited between reconnection attempts when connection is lost.
  # Valid units are 'ms', 's', 'm', 'h'. Note that these values can be combined, e.g. '24h30m15s'.
  max_reconnect_interval="{{ .Integration.MQTTV3.MaxReconnectInterval }}"

  # Terminate on connect error.
  #
  # When set to true, instead of re-trying to connect, the mioty BSSCI Adapter
  # process will be terminated on a connection error.
  terminate_on_connect_error={{ .Integration.MQTTV3.TerminateOnConnectError }}

  # State retained.
  #
  # By default this value is set to true and states are published as retained
  # MQTT messages. Setting this to false means that states will not be retained
  # by the MQTT broker.
  state_retained={{ .Integration.MQTTV3.StateRetained }}

  # State topic template.
  #
  # States are sent by the gateway as retained MQTT messages (by default)
  # so that the last message will be stored by the MQTT broker.
  # Only enabled when 'state_retained' is set to true.
  #
  # The following variables can be used in the template:
  #   * .BsEui - basestation EUI64
  #
  # Default: bssci/{{ .BsEui }}/state
  state_topic_template = "{{ .Integration.MQTTV3.StateTopicTemplate }}"


  # Event topic template.
  #
  # Events from basestations and endnodes are published on this topic.
  #
  # The following variables can be used in the template:
  #   * .BsEui        - basestation EUI64
  #   * .EventSource  - event source (ep=endpoint, bs=basestation)
  #   * .EventType    - event type (e.g. BasestationUplink, EndnodeUplink, etc.)
  #
  # Default: bssci/{{ .BsEui }}/event/{{ .EventSource }}/{{ .EventType }}
  event_topic_template = "{{ .Integration.MQTTV3.EventTopicTemplate }}"

  # Command topic template.
  #
  # Commands from SplitStack Server are received on this topic.
  #
  # The following variables can be used in the template:
  #   * .BsEui - basestation EUI64
  #
  # Default: bssci/{{ .BsEui }}/command/#
  command_topic_template = "{{ .Integration.MQTTV3.CommandTopicTemplate }}"

  # Response topic template.
  #
  # Responses from SplitStack Server are received on this topic. Responses are similar to commands 
  # but are used to directly reply to events sent by a basestation. 
  #
  # The following variables can be used in the template:
  #   * .BsEui - basestation EUI64
  #
  # Default: bssci/{{ .BsEui }}/response/#
  response_topic_template = "{{ .Integration.MQTTV3.ResponseTopicTemplate }}"

  # MQTT authentication.
  [integration.mqtt_v3.auth]
  # Type defines the MQTT authentication type to use.
  #
  # Set this to the name of one of the sections below.
  type="{{ .Integration.MQTTV3.Auth.Type }}"

    # Generic MQTT authentication.
    [integration.mqtt.auth.generic]
    # MQTT servers.
    #
    # Configure one or multiple MQTT server to connect to. Each item must be in
    # the following format: scheme://host:port where scheme is tcp, ssl or ws.
    servers=[{{ range $index, $elm := .Integration.MQTTV3.Auth.Generic.Servers }}
      "{{ $elm }}",{{ end }}
    ]

    # Connect with the given username (optional)
    username="{{ .Integration.MQTTV3.Auth.Generic.Username }}"

    # Connect with the given password (optional)
    password="{{ .Integration.MQTTV3.Auth.Generic.Password }}"

    # Quality of service level
    #
    # 0: at most once
    # 1: at least once
    # 2: exactly once
    #
    # Note: an increase of this value will decrease the performance.
    # For more information: https://www.hivemq.com/blog/mqtt-essentials-part-6-mqtt-quality-of-service-levels
    qos={{ .Integration.MQTTV3.Auth.Generic.QOS }}

    # Clean session
    #
    # Set the "clean session" flag in the connect message when this client
    # connects to an MQTT broker. By setting this flag you are indicating
    # that no messages saved by the broker for this client should be delivered.
    clean_session={{ .Integration.MQTTV3.Auth.Generic.CleanSession }}

    # Client ID
    #
    # Set the client id to be used by this client when connecting to the MQTT
    # broker. A client id must be no longer than 23 characters. When left blank,
    # a random id will be generated. This requires clean_session=true.
    client_id="{{ .Integration.MQTTV3.Auth.Generic.ClientID }}"

    # CA certificate file (optional)
    #
    # Use this when setting up a secure connection (when server uses ssl://...)
    # but the certificate used by the server is not trusted by any CA certificate
    # on the server (e.g. when self generated).
    ca_cert="{{ .Integration.MQTTV3.Auth.Generic.CACert }}"

    # mqtt TLS certificate file (optional)
    tls_cert="{{ .Integration.MQTTV3.Auth.Generic.TLSCert }}"

    # mqtt TLS key file (optional)
    tls_key="{{ .Integration.MQTTV3.Auth.Generic.TLSKey }}"


# Metrics configuration.
[metrics]

  # Metrics stored in Prometheus.
  #
  # These metrics expose information about the state of the mioty BSSCI Adapter
  # instance like number of messages processed, number of function calls, etc.
  [metrics.prometheus]
  # Expose Prometheus metrics endpoint.
  endpoint_enabled={{ .Metrics.Prometheus.EndpointEnabled }}

  # The ip:port to bind the Prometheus metrics server to for serving the
  # metrics endpoint.
  bind="{{ .Metrics.Prometheus.Bind }}"

`

var configCmd = &cobra.Command{
	Use:   "configfile",
	Short: "Print the mioty BSSCI Adapter configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(configTemplate))
		err := t.Execute(os.Stdout, config.C)
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}
