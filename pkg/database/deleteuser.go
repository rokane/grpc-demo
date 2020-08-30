package database

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteUserCriteria struct {
	First string
	Last string
}

type DeleteUserResp struct {
	User User
	Count int
}

func (ms *memstore) DeleteUser(ctx context.Context, criteria DeleteUserCriteria) (*DeleteUserResp, error) {
	toDelete := User{
		ID: uuid.NewSHA1(UserNamespace, []byte(criteria.First+criteria.Last)).String(),
		FirstName: criteria.First,
		LastName: criteria.Last,
	}
	count, exists := ms.seen[toDelete]
	if !exists {
		return nil, status.Error(codes.NotFound,
			fmt.Sprintf("have not seen user with name: %s %s", criteria.First, criteria.Last))
	}
	delete(ms.seen, toDelete)
	return &DeleteUserResp{User: toDelete, Count:count}, nil
}