package v2

import (
	"context"

	"github.com/rokane/grpc-demo/pkg/serializer"
	v2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
)

type Service struct {
	ctx        context.Context
	serializer serializer.DatabaseSerializer
}

func NewService() (*v2.GreeterService, error) {
	serializer, err := serializer.NewDBSerializer()
	if err != nil {
		return nil, err
	}
	s := &Service{
		ctx:        context.Background(),
		serializer: serializer,
	}
	return &v2.GreeterService{
		SayHello:   s.SayHello,
		SayGoodbye: s.SayGoodbye,
	}, nil
}
