package database

import (
	"context"

	"github.com/google/uuid"
)

var (
	UserNamespace = uuid.Must(uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
)

type User struct {
	ID        string
	FirstName string
	LastName  string
}

type Storer interface {
	RegisterDetails(context.Context, RegisterDetailsCriteria) (*RegisterDetailsResp, error)
	DeleteUser(context.Context, DeleteUserCriteria) (*DeleteUserResp, error)
}

type memstore struct {
	seen map[User]int
}

func NewMemstore() (Storer, error) {
	return &memstore{seen: make(map[User]int)}, nil
}
