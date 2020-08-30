package database

import "context"

type DeleteUserCriteria struct {}

type DeleteUserResp struct {}

func (ms *memstore) DeleteUser(ctx context.Context, criteria DeleteUserCriteria) (*DeleteUserResp, error) {
	return nil, nil
}