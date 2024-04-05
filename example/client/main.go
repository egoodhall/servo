package main

import (
	"context"
	"fmt"

	"github.com/egoodhall/servo/example"
	"gopkg.in/h2non/gentleman.v2"
)

func main() {
	client := example.NewDelegatingEchoServiceHttpClient(
		gentleman.New().URL("http://localhost:8080"),
	)

	response, err := client.Echo(context.Background(), &example.EchoRequest{Message: "hello world!"})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Message)
}
