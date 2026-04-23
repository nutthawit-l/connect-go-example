package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	greetv1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"

	"connectrpc.com/connect"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	ctx, callInfo := connect.NewClientContext(context.Background())
	callInfo.RequestHeader().Set("Acme-Tenant-Id", "1234")
	res, err := client.Greet(
		ctx,
		&greetv1.GreetRequest{Name: "Jane"},
	)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(callInfo.ResponseHeader().Get("Greet-Version"))
	log.Println(res.Greeting)
}
