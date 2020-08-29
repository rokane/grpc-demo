package v2

import (
	"context"

	v2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
)

func (s *Service) SayHello(ctx context.Context, requ *v2.SayHelloRequest) (*v2.SayHelloResponse, error) {
	return nil, nil
}