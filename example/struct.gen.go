package example

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
	Ids    []int64           `json:"ids"`
}

type MetricType string

const (
	Counter MetricType = "COUNTER"
	Gauge   MetricType = "GAUGE"
)

type EchoService interface {
	Echo(*EchoRequest) (*EchoResponse, error)
}
type MetricsPublisher interface {
	Publish(*Metric) error
}
