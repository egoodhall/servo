package main

import (
	"fmt"

	"github.com/egoodhall/servo/example"
	"gopkg.in/h2non/gentleman.v2"
)

type EchoServiceImpl struct {
}

func (svc *EchoServiceImpl) Echo(req *example.EchoRequest) (*example.EchoResponse, error) {
	return &example.EchoResponse{Message: req.Message}, nil
}

func main() {
	client := example.NewDelegatingEchoServiceHttpClient(
		gentleman.New().URL("http://localhost:8080"),
	)

	response, err := client.Echo(&example.EchoRequest{Message: "hello world!"})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Message)
}
