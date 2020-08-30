package database

import "context"

type RegisteredUser struct {
}

type RegisterDetailsCriteria struct {}

type RegisterDetailsResp struct {}

func (ms *memstore) RegisterDetails(ctx context.Context, criteria RegisterDetailsCriteria) (*RegisterDetailsResp, error) {
	return nil, nil
}