package v2

import (
	"context"

	v2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
)

func (s *Service) SayGoodbye(ctx context.Context, req *v2.SayGoodbyeRequest) (*v2.SayGoodbyeResponse, error) {
	return nil, nil
}