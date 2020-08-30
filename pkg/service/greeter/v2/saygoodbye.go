package v2

import (
	"context"
	db "github.com/rokane/grpc-demo/pkg/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
)

func (s *Service) SayGoodbye(ctx context.Context, req *v2.SayGoodbyeRequest) (*v2.SayGoodbyeResponse, error) {
	criteria := db.DeleteUserCriteria{}
	resp, err := s.serializer.DeleteUser(ctx, criteria)
	if err != nil {
		return nil, status.Error(codes.Internal, "unable to remove user")
	}
	return resp.ToV2()
}
