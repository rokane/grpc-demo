package main

import (
	"log"
	"net"

	apiv1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
	apiv2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
	greeterv1 "github.com/rokane/grpc-demo/pkg/service/greeter/v1"
	greeterv2 "github.com/rokane/grpc-demo/pkg/service/greeter/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("unable to listen on port ", port)
	}
	server := grpc.NewServer()
	reflection.Register(server)

	// Register API v1
	greeterV1, err := greeterv1.NewService()
	if err != nil {
		log.Fatal("unable to initialise v1 service")
	}
	apiv1.RegisterGreeterService(server, greeterV1)

	// Register API v2
	greeterV2, err := greeterv2.NewService()
	if err != nil {
		log.Fatal("unable to initialise v2 service")
	}
	apiv2.RegisterGreeterService(server, greeterV2)

	log.Printf("listening on port %s", port)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
