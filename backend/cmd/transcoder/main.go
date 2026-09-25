
package main

import (
	"context"
	"log"
	"net/http"

	todov1 "example.com/todo-platform/backend/api/todo/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	mux := runtime.NewServeMux()

	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	}

	err := todov1.RegisterTodoServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:50051",
		dialOptions,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(
		"HTTP → gRPC transcoder listening on :8081",
	)

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}