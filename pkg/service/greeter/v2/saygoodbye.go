package v2

import (
	"context"
	db "github.com/rokane/grpc-demo/pkg/database"
	"log"

	v2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
)

func (s *Service) SayGoodbye(ctx context.Context, req *v2.SayGoodbyeRequest) (*v2.SayGoodbyeResponse, error) {
	criteria := db.DeleteUserCriteria{
		First: req.FirstName,
		Last: req.LastName,
	}
	log.Printf("processing SayGoodbye request: %v", req)

	resp, err := s.serializer.DeleteUser(ctx, criteria)
	if err != nil {
		return nil, err
	}
	return resp.ToV2()
}
