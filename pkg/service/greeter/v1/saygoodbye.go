package v1

import (
	"context"
	db "github.com/rokane/grpc-demo/pkg/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
)

func (s *service) SayGoodbye(ctx context.Context, req *v1.SayGoodbyeRequest) (*v1.SayGoodbyeResponse, error) {
	criteria := db.DeleteUserCriteria{}
	resp, err := s.serializer.DeleteUser(ctx, criteria)
	if err != nil {
		return nil, status.Error(codes.Internal, "unable to remove user")
	}
	return resp.ToV1()
}
