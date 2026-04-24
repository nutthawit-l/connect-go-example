package main

import (
	"context"
	greetv1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
)

func main() {
	ctx, callInfo := connect.NewClientContext(context.Background())
	_, err := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	).Greet(
		ctx,
		&greetv1.GreetRequest{
			Name: "Jane",
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Doesn't contain "Greet-Version" because any HTTP headers prefixed with
	// Trailer- are treated as trailers.
	fmt.Println(callInfo.RequestHeader())
	// Prefixes are automatically stripped.
	fmt.Println(callInfo.ResponseTrailer().Get("Greet-Version"))
}
