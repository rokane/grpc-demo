package serializer

import (
	"context"

	v1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
	v2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
	db "github.com/rokane/grpc-demo/pkg/database"
)

type SayHelloSerializer interface {
	ToV1Alpha1() (*v1.SayHelloResponse, error)
	ToV1Alpha2() (*v2.SayHelloResponse, error)
}

type SayGoodbyeSerializer interface {
	ToV1Alpha1() (*v1.SayGoodbyeResponse, error)
	ToV1Alpha2() (*v2.SayGoodbyeResponse, error)
}

type DatabaseSerializer interface {
	RegisterDetails(context.Context, db.User) (SayHelloSerializer, error)
	DeleteUser(context.Context, db.RegisteredUser) (SayGoodbyeSerializer, error)
}
