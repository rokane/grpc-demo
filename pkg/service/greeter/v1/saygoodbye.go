package v1

import (
	"context"
	db "github.com/rokane/grpc-demo/pkg/database"
	"log"

	v1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
)

func (s *service) SayGoodbye(ctx context.Context, req *v1.SayGoodbyeRequest) (*v1.SayGoodbyeResponse, error) {
	criteria := db.DeleteUserCriteria{
		First: req.Name,
	}
	log.Printf("processing SayGoodbye request: %v", req)

	resp, err := s.serializer.DeleteUser(ctx, criteria)
	if err != nil {
		return nil, err
	}
	return resp.ToV1()
}
