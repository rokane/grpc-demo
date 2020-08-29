package v1alpha1

import (
	"context"

	v1alpha1 "github.com/rokane/pkg/api/greeter/v1alpha1"
)

func (s *Server) SayGoodbye(ctx context.Context, req *v1alpha1.SayGoodbyeRequest) (*v1alpha1.SayGoodbyeResponse, error) {
	return nil, nil
}
