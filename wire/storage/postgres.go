package storage

import (
	"context"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"toki/code-golang/wire/config"
)

var DBProvider = wire.NewSet(NewPostgresDB)

type PostgresStorage struct {
	db *sqlx.DB
}

func NewPostgresDB(cfg *config.Config) *PostgresStorage {
	return &PostgresStorage{}
}

func (p *PostgresStorage) GetUserByID(ctx context.Context, id int64) (*User, error) {
	return &User{
		ID:   1,
		Name: "model",
	}, nil
}
