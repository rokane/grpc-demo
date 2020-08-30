package serializer

import (
	"context"

	v1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
	v2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
	db "github.com/rokane/grpc-demo/pkg/database"
)

type SayHelloSerializer interface {
	ToV1() (*v1.SayHelloResponse, error)
	ToV2() (*v2.SayHelloResponse, error)
}

type SayGoodbyeSerializer interface {
	ToV1() (*v1.SayGoodbyeResponse, error)
	ToV2() (*v2.SayGoodbyeResponse, error)
}

type DatabaseSerializer interface {
	RegisterDetails(context.Context, db.RegisterDetailsCriteria) (SayHelloSerializer, error)
	DeleteUser(context.Context, db.DeleteUserCriteria) (SayGoodbyeSerializer, error)
}

type dbserializer struct {
	storer db.Storer
}

func NewDBSerializer() (DatabaseSerializer, error) {
	storer, err := db.NewMemstore()
	if err != nil {
		return nil, err
	}
	return &dbserializer{storer}, nil
}
