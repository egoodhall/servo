package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:generate servoc generate ../example.srvo --option=gohttp.client.enabled=false

type EchoServiceImpl struct {
}

func (svc *EchoServiceImpl) Echo(ctx echo.Context, req *EchoRequest) (*EchoResponse, error) {
	return &EchoResponse{Message: req.Message}, nil
}

func main() {
	srv := NewEchoServiceHttpServer(new(EchoServiceImpl))

	srv.Use(
		middleware.RequestID(),
		middleware.Logger(),
		middleware.RemoveTrailingSlash(),
		middleware.BodyLimit("4KB"),
		middleware.Recover(),
	)

	if err := srv.Start("localhost:8080"); err != nil {
		panic(err)
	}
}
