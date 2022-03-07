package storage

import (
	"context"
)

type User struct {
	ID   int64
	Name string
}

type Storage interface {
	GetUserByID(ctx context.Context, id int64) (*User, error)
}
