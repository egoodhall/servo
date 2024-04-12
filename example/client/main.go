package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/egoodhall/servo/example"
)

func main() {
	client := example.NewDelegatingEchoServiceHttpClient(
		"http://localhost:8080",
		new(http.Client),
	)

	response, err := client.Echo(context.Background(), &example.EchoRequest{Message: "hello world!"})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Message)
}
