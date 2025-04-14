package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewRequestLogger() (*zap.Logger, *zap.SugaredLogger, error) {
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			return nil, nil, fmt.Errorf("failed to create log directory: %w", err)
		}
	}

	logFile := fmt.Sprintf("%s/request-%s.log", logDir, time.Now().Format("2006-01-02"))
	file, err := os.Create(logFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create log file: %w", err)
	}

	writeSyncer := zapcore.AddSync(file)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "T"
	encoderConfig.LevelKey = "L"
	encoderConfig.MessageKey = "M"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writeSyncer, zap.InfoLevel)
	logger := zap.New(core)
	sugar := logger.Sugar()

	return logger, sugar, nil
}
