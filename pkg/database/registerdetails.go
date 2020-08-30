package database

import (
	"context"

	"github.com/google/uuid"
)

type RegisterDetailsCriteria struct {
	First string
	Last string
}

type RegisterDetailsResp struct {
	User User
	Exists bool
}

func (ms *memstore) RegisterDetails(ctx context.Context, criteria RegisterDetailsCriteria) (*RegisterDetailsResp, error) {
	id := uuid.NewSHA1(UserNamespace, []byte(criteria.First+criteria.Last)).String()
	incoming := User{
		ID:        id,
		FirstName: criteria.First,
		LastName:  criteria.Last,
	}
	count, exists := ms.seen[incoming]
	if exists {
		ms.seen[incoming] = count + 1
		return &RegisterDetailsResp{User: incoming, Exists: true}, nil
	}
	ms.seen[incoming] = 1
	return &RegisterDetailsResp{User: incoming, Exists: false}, nil
}