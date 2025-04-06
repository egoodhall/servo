package main

import (
	"context"
	"fmt"
)

//go:generate servoc generate ../example.srvo --option=gohttp.server.enabled=false

func main() {
	client := NewEchoServiceHttpClient(
		"http://localhost:8080",
	)

	response, err := client.Echo(context.Background(), &EchoRequest{Message: "hello world!"})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Message)
}
