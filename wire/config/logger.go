package config

import (
	"fmt"
	"go.uber.org/zap"
)

type Logger zap.Logger

func NewLogger() zap.Logger {
	fmt.Println("init logger")
	return zap.Logger{}
}
