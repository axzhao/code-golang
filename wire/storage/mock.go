package storage

import (
	"context"
	"github.com/google/wire"
	"toki/code-golang/wire/config"
)

var MockDBProvider = wire.NewSet(NewMockDB)

type MockDB struct{}

func NewMockDB(cfg *config.Config) *MockDB {
	return &MockDB{}
}

func (p *MockDB) GetUserByID(ctx context.Context, id int64) (*User, error) {
	return &User{
		ID:   1,
		Name: "mock",
	}, nil
}
