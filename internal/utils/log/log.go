package log

import (
	"go.uber.org/zap"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func NewLog() *zap.SugaredLogger {
	if sugar != nil {
		return sugar
	}
	Init()
	return sugar
}

func Init() {
	logger, _ = zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger) // flushes buffer, if any
	sugar = logger.Sugar()
}
