package v1

import (
	"context"

	v1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
)

func (s *service) SayHello(ctx context.Context, req *v1.SayHelloRequest) (*v1.SayHelloResponse, error) {
	return nil, nil
}
