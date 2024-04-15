// Code generated by servoc (gohttp plugin). DO NOT EDIT.
package example

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

/////////////////////////////
// EchoService HTTP server //
/////////////////////////////

func NewEchoServiceHttpServer(svc EchoService) *echo.Echo {
	srv := echo.New()
	RegisterEchoServiceRPCs(svc, srv)
	return srv
}

func RegisterEchoServiceRPCs(svc EchoService, srv *echo.Echo) {
	RegisterEchoServiceRPCsGroup(svc, srv.Group("/"))
}

func RegisterEchoServiceRPCsGroup(svc EchoService, srv *echo.Group) {
	compat := &echoServiceHttpServer{svc}
	srv.POST("echo-service/echo", compat.Echo)
}

type echoServiceHttpServer struct {
	svc EchoService
}

// HTTP compatibility wrapper for EchoService.echo.
func (s *echoServiceHttpServer) Echo(c echo.Context) error {
	req := new(EchoRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	res, err := s.svc.Echo(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, res)
}

/////////////////////////////
// EchoService HTTP client //
/////////////////////////////

func NewEchoServiceHttpClient(baseUrl string) EchoService {
	return NewDelegatingEchoServiceHttpClient(baseUrl, new(http.Client))
}

func NewDelegatingEchoServiceHttpClient(baseUrl string, delegate *http.Client) EchoService {
	return &echoServiceHttpClient{baseUrl, delegate}
}

var _ EchoService = new(echoServiceHttpClient)

type echoServiceHttpClient struct {
	baseUrl  string
	delegate *http.Client
}

func (client *echoServiceHttpClient) Echo(ctx context.Context, request *EchoRequest) (*EchoResponse, error) {
	u, err := url.JoinPath(client.baseUrl, "/echo-service/echo")
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(request); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.delegate.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code " + strconv.Itoa(res.StatusCode))
	}

	response := new(EchoResponse)
	return response, json.NewDecoder(res.Body).Decode(response)
}

//////////////////////////////////
// EchoService test HTTP client //
//////////////////////////////////

func NewEchoServiceTestHttpClient(svc EchoService) EchoService {
	return &echoServiceHttpTestClient{svc}
}

var _ EchoService = new(echoServiceHttpTestClient)

type echoServiceHttpTestClient struct {
	service EchoService
}

func (client *echoServiceHttpTestClient) Echo(_ context.Context, request *EchoRequest) (*EchoResponse, error) {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(request); err != nil {
		return nil, err
	}
	req := httptest.NewRequest(http.MethodPost, "/echo-service/echo", body)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	ctx := NewEchoServiceHttpServer(client.service).NewContext(req, res)
	if err := (&echoServiceHttpServer{client.service}).Echo(ctx); err != nil {
		return nil, err
	} else if res.Code != http.StatusOK {
		return nil, errors.New("unexpected status code " + strconv.Itoa(res.Code))
	}

	response := new(EchoResponse)
	return response, json.NewDecoder(res.Body).Decode(response)
}

//////////////////////////////////
// TelemetryService HTTP server //
//////////////////////////////////

func NewTelemetryServiceHttpServer(svc TelemetryService) *echo.Echo {
	srv := echo.New()
	RegisterTelemetryServiceRPCs(svc, srv)
	return srv
}

func RegisterTelemetryServiceRPCs(svc TelemetryService, srv *echo.Echo) {
	RegisterTelemetryServiceRPCsGroup(svc, srv.Group("/"))
}

func RegisterTelemetryServiceRPCsGroup(svc TelemetryService, srv *echo.Group) {
	compat := &telemetryServiceHttpServer{svc}
	srv.POST("telemetry-service/publish", compat.Publish)
}

type telemetryServiceHttpServer struct {
	svc TelemetryService
}

// HTTP compatibility wrapper for TelemetryService.publish.
func (s *telemetryServiceHttpServer) Publish(c echo.Context) error {
	req := new(Telemetry)
	if err := c.Bind(req); err != nil {
		return err
	}
	err := s.svc.Publish(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.NoContent(http.StatusNoContent)
}

//////////////////////////////////
// TelemetryService HTTP client //
//////////////////////////////////

func NewTelemetryServiceHttpClient(baseUrl string) TelemetryService {
	return NewDelegatingTelemetryServiceHttpClient(baseUrl, new(http.Client))
}

func NewDelegatingTelemetryServiceHttpClient(baseUrl string, delegate *http.Client) TelemetryService {
	return &telemetryServiceHttpClient{baseUrl, delegate}
}

var _ TelemetryService = new(telemetryServiceHttpClient)

type telemetryServiceHttpClient struct {
	baseUrl  string
	delegate *http.Client
}

func (client *telemetryServiceHttpClient) Publish(ctx context.Context, request *Telemetry) error {
	u, err := url.JoinPath(client.baseUrl, "/telemetry-service/publish")
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(request); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.delegate.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return errors.New("unexpected status code " + strconv.Itoa(res.StatusCode))
	}
	return nil
}

///////////////////////////////////////
// TelemetryService test HTTP client //
///////////////////////////////////////

func NewTelemetryServiceTestHttpClient(svc TelemetryService) TelemetryService {
	return &telemetryServiceHttpTestClient{svc}
}

var _ TelemetryService = new(telemetryServiceHttpTestClient)

type telemetryServiceHttpTestClient struct {
	service TelemetryService
}

func (client *telemetryServiceHttpTestClient) Publish(_ context.Context, request *Telemetry) error {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(request); err != nil {
		return err
	}
	req := httptest.NewRequest(http.MethodPost, "/telemetry-service/publish", body)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	ctx := NewTelemetryServiceHttpServer(client.service).NewContext(req, res)
	if err := (&telemetryServiceHttpServer{client.service}).Publish(ctx); err != nil {
		return err
	} else if res.Code != http.StatusNoContent {
		return errors.New("unexpected status code " + strconv.Itoa(res.Code))
	}
	return nil
}
