package v1alpha2

import (
	"context"

	v1alpha2 "github.com/rokane/pkg/api/greeter/v1alpha2"
)

func (s *Server) SayGoodbye(ctx context.Context, req *v1alpha2.SayGoodbyeRequest) (*v1alpha2.SayGoodbyeResponse, error) {
	return nil, nil
}
