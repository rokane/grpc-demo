package v1

import (
	"context"

	v1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
)

func (s *service) SayGoodbye(context.Context, *v1.SayGoodbyeRequest) (*v1.SayGoodbyeResponse, error) {
	return nil, nil
}
