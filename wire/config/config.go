package config

import (
	"fmt"
	"github.com/google/wire"
	"time"
)

var Provider = wire.NewSet(New, NewLogger)

type DB struct {
	Dsn             string        `koanf:"dsn"`
	MigrationPath   string        `koanf:"migration_path"`
	MaxIdleConns    int           `koanf:"max_idle_conns"`
	MaxOpenConns    int           `koanf:"max_open_conns"`
	ConnMaxLifeTime time.Duration `koanf:"conn_max_life_time"`
	ConnMaxIdleTime time.Duration `koanf:"conn_max_idle_time"`
}

type Config struct {
	DB DB
}

func New(path string) (*Config, error) {
	fmt.Println("path: ", path)
	return &Config{}, nil
}
