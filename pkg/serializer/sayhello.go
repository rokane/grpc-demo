package serializer

import (
	"context"
	v1 "github.com/rokane/grpc-demo/pkg/api/greeter/v1"
	v2 "github.com/rokane/grpc-demo/pkg/api/greeter/v2"
	db "github.com/rokane/grpc-demo/pkg/database"
)

type registerDetailsWrapper struct {
	*db.RegisterDetailsResp
}

func (w *registerDetailsWrapper) ToV1() (*v1.SayHelloResponse, error) {
	return nil, nil
}

func (w *registerDetailsWrapper) ToV2() (*v2.SayHelloResponse, error) {
	return nil, nil
}

func (dbs *dbserializer) RegisterDetails(ctx context.Context, criteria db.RegisterDetailsCriteria) (SayHelloSerializer, error) {
	resp, err := dbs.storer.RegisterDetails(ctx, criteria)
	if err != nil {
		return nil, err
	}
	return &registerDetailsWrapper{resp}, nil
}