package mqtt

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	pc = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "integration_mqtt_event_count",
		Help: "The number of gateway events published by the MQTT integration (per subscriber, source, event).",
	}, []string{"basestation", "source", "event"})

	sc = promauto.NewCounter(prometheus.CounterOpts{
		Name: "integration_mqtt_state_count",
		Help: "The number of gateway states published by the MQTT integration",
	})

	cc = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "integration_mqtt_command_count",
		Help: "The number of commands received by the MQTT integration (per subscriber).",
	}, []string{"basestation"})

	rr = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "integration_mqtt_response_count",
		Help: "The number of responsess received by the MQTT integration (per subscriber).",
	}, []string{"basestation"})

	mqttc = promauto.NewCounter(prometheus.CounterOpts{
		Name: "integration_mqtt_connect_count",
		Help: "The number of times the integration connected to the MQTT broker.",
	})

	mqttd = promauto.NewCounter(prometheus.CounterOpts{
		Name: "integration_mqtt_disconnect_count",
		Help: "The number of times the integration disconnected from the MQTT broker.",
	})

	mqttr = promauto.NewCounter(prometheus.CounterOpts{
		Name: "integration_mqtt_reconnect_count",
		Help: "The number of times the integration reconnected to the MQTT broker (this also increments the disconnect and connect counters).",
	})
)

func mqttEventCounter(c string, s string, e string) prometheus.Counter {
	return pc.With(prometheus.Labels{"basestation": c, "source": s, "event": e})
}

func mqttStateCounter() prometheus.Counter {
	return sc
}

func mqttCommandCounter(c string) prometheus.Counter {
	return cc.With(prometheus.Labels{"basestation": c})
}

func mqttResponseCounter(c string) prometheus.Counter {
	return rr.With(prometheus.Labels{"basestation": c})
}

func mqttConnectCounter() prometheus.Counter {
	return mqttc
}

func mqttDisconnectCounter() prometheus.Counter {
	return mqttd
}

func mqttReconnectCounter() prometheus.Counter {
	return mqttr
}
