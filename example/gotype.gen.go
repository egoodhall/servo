// Code generated by servoc (gostruct plugin). DO NOT EDIT.
package example

import "context"

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Message string `json:"message"`
}

type Metric struct {
	Type   MetricType        `json:"type"`
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
	Value  float64           `json:"value"`
}

type Log struct {
	Labels map[string]string `json:"labels"`
	Value  string            `json:"value"`
}

type MetricType string

const (
	MetricType_Counter MetricType = "COUNTER"
	MetricType_Gauge   MetricType = "GAUGE"
)

type Telemetry struct {
	TelemetryType string  `json:"@type"`
	Log           *Log    `json:"log,omitempty"`
	Metric        *Metric `json:"metric,omitempty"`
}

type EchoService interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

type TelemetryService interface {
	Publish(context.Context, *Telemetry) error
}