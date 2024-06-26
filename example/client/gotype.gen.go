// Code generated by servoc (gostruct plugin). DO NOT EDIT.
package main

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type (
	URN string
	Bit bool
)

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Message   string    `json:"message"`
	Id        uuid.UUID `json:"id"`
	Signature []byte    `json:"signature"`
}

type Metric struct {
	Type   MetricType        `json:"type"`
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
	Value  float64           `json:"value"`
	At     time.Time         `json:"at"`
}

type Binary struct {
	Location URN   `json:"location"`
	Bits     []Bit `json:"bits"`
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
