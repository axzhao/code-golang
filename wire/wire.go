//go:build wireinject

package wireDemo

import (
	"github.com/google/wire"
	"toki/code-golang/wire/config"
	"toki/code-golang/wire/storage"
)

func InitializeService(path string) (*Service, error) {
	panic(wire.Build(config.Provider, storage.DBProvider, NewService, wire.Bind(new(storage.Storage), new(*storage.PostgresStorage))))
}

func InitializeMockService(path string) (*Service, error) {
	panic(wire.Build(config.Provider, storage.MockDBProvider, NewService, wire.Bind(new(storage.Storage), new(*storage.MockDB))))
}
