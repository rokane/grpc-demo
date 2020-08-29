package v1alpha1

import (
	"context"

	v1alpha1 "github.com/rokane/pkg/api/greeter/v1alpha1"
)

func (s *Server) SayHello(ctx context.Context, req *v1alpha1.SayHelloRequest) (*v1alpha1.SayHelloResponse, error) {
	return nil, nil
}
