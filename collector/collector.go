package collector

import "github.com/prometheus/client_golang/prometheus"

const (
	namespace = "rabbitmq_cli_consumer"
)

var (
	// ProcessCounter is a Prometheus metric describing the total number of processes executed.
	ProcessCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "process_total",
			Help:      "The total number of processes executed.",
		},
		[]string{"exit_code"},
	)

	// ProcessDuration is a Prometheus metric describing the time spent by the consumer to process the message.
	ProcessDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "process_duration_seconds",
			Help:      "The time spent by the consumer to process the message.",
		},
	)

	// MessageDuration is a Prometheus metric describing the time spent from publishing to finished processing the message.
	MessageDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "message_duration_seconds",
			Help:      "The time spent from publishing to finished processing the message.",
		},
	)
)
