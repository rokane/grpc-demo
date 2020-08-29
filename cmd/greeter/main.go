package main

import (
	api_v1alpha1 "github.com/rokane/grpc-demo/pkg/api/proto/greeter/v1alpha1"
	api_v1alpha2 "github.com/rokane/grpc-demo/pkg/api/proto/greeter/v1alpha2"
	v1alpha1 "github.com/rokane/grpc-demo/pkg/service/greeter/v1alpha1"
	v1alpha2 "github.com/rokane/grpc-demo/pkg/service/greeter/v1alpha2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()
	reflection.Register(s)

	service, err := v1alpha1.NewServer()
	if err != nil {
		// Handle
	}
	api_v1alpha1.RegisterTransactionsAPIServer(s, service)

	service, err = v1alpha2.NewServer()
	if err != nil {
		// Handle
	}
	api_v1alpha2.RegisterTransactionsAPIServer(s, service)

}
