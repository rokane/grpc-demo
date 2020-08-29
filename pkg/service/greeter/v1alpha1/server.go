package v1alpha1

import (
	"context"

	"github.com/rokane/grpc-demo/pkg/serializer"
)

type Server struct {
	ctx        context.Context
	serializer serializer.DatabaseSerializer
}

func NewServer() (*Server, error) {
	return &Server{
		ctx:        context.Background(),
		serializer: nil,
	}, nil
}
