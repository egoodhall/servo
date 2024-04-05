package main

import (
	"context"

	"github.com/egoodhall/servo/example"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServiceImpl struct {
}

func (svc *EchoServiceImpl) Echo(ctx context.Context, req *example.EchoRequest) (*example.EchoResponse, error) {
	return &example.EchoResponse{Message: req.Message}, nil
}

func main() {
	srv := example.NewEchoServiceHttpServer(new(EchoServiceImpl))

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
