package model

import (
	"context"
	"database/sql"
	"xorm.io/xorm"
)

type PostgresStorage struct {
	db *xorm.Engine
}

func NewPostgresStorage(db *xorm.Engine) *PostgresStorage {
	return &PostgresStorage{db: db}
}

type User struct {
	ID   int64  `json:"id,string"`
	Name string `json:"name"`
}

func (*User) TableName() string {
	return "user"
}

func (c *PostgresStorage) GetUser(ctx context.Context, id int64) (u User, err error) {
	has, err := c.db.Table(User{}).ID(id).Get(&u)
	if !has {
		return User{}, sql.ErrNoRows
	}
	return
}

func (c *PostgresStorage) AddUser(u User) (id int64, err error) {
	id, err = c.db.Insert(&u)
	return
}
