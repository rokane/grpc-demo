package v1

import (
	"context"
	db "github.com/rokane/grpc-demo/pkg/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
)

func (s *service) SayHello(ctx context.Context, req *v1.SayHelloRequest) (*v1.SayHelloResponse, error) {
	criteria := db.RegisterDetailsCriteria{}
	resp, err := s.serializer.RegisterDetails(ctx, criteria)
	if err != nil {
		return nil, status.Error(codes.Internal, "unable to register user")
	}
	return resp.ToV1()
}
