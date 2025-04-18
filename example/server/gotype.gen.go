// Code generated by servoc (gostruct plugin). DO NOT EDIT.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	Location *URN  `json:"location,omitempty"`
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
	Name          *string `json:"name,omitempty"`
}

func (u *Telemetry) UnmarshalJSON(p []byte) error {
	if u == nil {
		return errors.New("unable to unmarshal into nil Telemetry pointer")
	}

	var discriminator struct {
		Type string `json:"@type"`
	}
	if err := json.Unmarshal(p, &discriminator); err != nil {
		return err
	}

	switch discriminator.Type {
	case "log", "metric", "name":
		return json.Unmarshal(p, u)
	default:
		return fmt.Errorf("unknown Telemetry type \"%s\"", discriminator.Type)
	}

}

func NewTelemetryLog(value Log) *Telemetry {
	return &Telemetry{TelemetryType: "log", Log: &value}
}

func NewTelemetryMetric(value Metric) *Telemetry {
	return &Telemetry{TelemetryType: "metric", Metric: &value}
}

func NewTelemetryName(value string) *Telemetry {
	return &Telemetry{TelemetryType: "name", Name: &value}
}

type EchoService interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

type TelemetryService interface {
	Publish(context.Context, *Telemetry) error
}
