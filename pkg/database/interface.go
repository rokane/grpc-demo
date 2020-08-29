package database

import "context"

type User struct {
}

type RegisteredUser struct {
}

type Storer interface {
	RegisterDetails(context.Context, User) (RegisteredUser, error)
	DeleteUser(context.Context, RegisteredUser) (RegisteredUser, error)
}
