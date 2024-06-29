package main

import (
	"context"
	"fmt"
	"net/http"
)

//go:generate servoc generate ../example.srvo --option=gohttp.server.enabled=false

func main() {
	client := NewDelegatingEchoServiceHttpClient(
		"http://localhost:8080",
		new(http.Client),
	)

	response, err := client.Echo(context.Background(), &EchoRequest{Message: "hello world!"})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Message)
}
