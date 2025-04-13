package logger

import (
	"go.uber.org/zap"
)

func NewLogger(env string) *zap.SugaredLogger {
	var logger *zap.Logger
	var err error

	if env == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	return logger.Sugar()
}
