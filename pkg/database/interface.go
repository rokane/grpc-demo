package database

import "context"

type Storer interface {
	RegisterDetails(context.Context, RegisterDetailsCriteria) (*RegisterDetailsResp, error)
	DeleteUser(context.Context, DeleteUserCriteria) (*DeleteUserResp, error)
}

type memstore struct {
	users map[string]int
}

func NewMemstore() (Storer, error) {
	return &memstore{users: make(map[string]int)}, nil
}
