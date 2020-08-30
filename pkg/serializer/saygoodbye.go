package serializer

import (
	"context"
	v1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
	v2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
	db "github.com/rokane/grpc-demo/pkg/database"
)

type deleteUserWrapper struct {
	*db.DeleteUserResp
}

func (w *deleteUserWrapper) ToV1() (*v1.SayGoodbyeResponse, error) {
	return nil, nil
}

func (w *deleteUserWrapper) ToV2() (*v2.SayGoodbyeResponse, error) {
	return nil, nil
}

func (dbs *dbserializer) DeleteUser(ctx context.Context, criteria db.DeleteUserCriteria) (SayGoodbyeSerializer, error) {
	resp, err := dbs.storer.DeleteUser(ctx, criteria)
	if err != nil {
		return nil, err
	}
	return &deleteUserWrapper{resp}, nil
}