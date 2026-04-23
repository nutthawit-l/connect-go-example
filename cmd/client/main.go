package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	greetv1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"

	"connectrpc.com/connect"
)

func main() {
	_, err := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	).Greet(
		context.Background(),
		&greetv1.GreetRequest{Name: "Jane"},
	)
	if connectErr := new(connect.Error); errors.As(err, &connectErr) {
		fmt.Println(connectErr.Meta().Get("Greet-Version"))
	}
}
