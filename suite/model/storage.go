package model

//go:generate mockgen -destination=.mocks/mock_storage.go -source=./storage.go

import (
	"context"
	_ "github.com/golang/mock/mockgen/model"
)

type Storage interface {
	GetUser(ctx context.Context, id int64) (User, error)
	AddUser(u User) (int64, error)
}
