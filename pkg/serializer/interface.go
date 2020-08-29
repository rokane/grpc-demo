package serializer

import (
	"context"

	v1alpha1 "github.com/rokane/pkg/api/greeter/v1alpha1"
	v1alpha2 "github.com/rokane/pkg/api/greeter/v1alpha2"
	db "github.com/rokane/pkg/database"
)

type SayHelloSerializer interface {
	ToV1Alpha1() (*v1alpha1.SayHelloResponse, error)
	ToV1Alpha2() (*v1alpha2.SayHelloResponse, error)
}

type SayGoodbyeSerializer interface {
	ToV1Alpha1() (*v1alpha1.SayGoodbyeResponse, error)
	ToV1Alpha2() (*v1alpha2.SayGoodbyeResponse, error)
}

type DatabaseSerializer interface {
	RegisterDetails(context.Context, db.User) (SayHelloSerializer, error)
	DeleteUser(context.Context, db.RegisteredUser) (SayGoodbyeSerializer, error)
}
