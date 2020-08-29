package v1alpha2

import (
	"context"

	v1alpha2 "github.com/rokane/pkg/api/greeter/v1alpha2"
)

func (s *Server) SayHello(ctx context.Context, req *v1alpha2.SayHelloRequest) (*v1alpha2.SayHelloResponse, error) {
	return nil, nil
}
