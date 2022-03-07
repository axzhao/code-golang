package wireDemo

import (
	"go.uber.org/zap"
	"toki/code-golang/wire/storage"
)

type Service struct {
	logger  zap.Logger
	storage storage.Storage
}

func NewService(logger zap.Logger, storage storage.Storage) *Service {
	return &Service{logger: logger, storage: storage}
}
