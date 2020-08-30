package v2

import (
	"context"
	db "github.com/rokane/grpc-demo/pkg/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	v2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
)

func (s *Service) SayHello(ctx context.Context, req *v2.SayHelloRequest) (*v2.SayHelloResponse, error) {
	criteria := db.RegisterDetailsCriteria{
		First: req.FirstName,
		Last: req.LastName,
	}
	log.Printf("processing SayHello request: %v", req)

	resp, err := s.serializer.RegisterDetails(ctx, criteria)
	if err != nil {
		return nil, status.Error(codes.Internal, "unable to register user")
	}
	return resp.ToV2()
}
