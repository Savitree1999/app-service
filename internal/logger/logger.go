package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(env string) *zap.SugaredLogger {
	logDir := "logs"

	// Set default config (development-friendly)
	cfg := zap.NewDevelopmentConfig()

	// Set encoder (JSON format)
	encoder := zapcore.NewJSONEncoder(cfg.EncoderConfig)

	// Set log level
	level := zapcore.InfoLevel
	if env == "dev" {
		level = zapcore.DebugLevel
	}

	// Create logs directory if not exists
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, 0755)
		if err != nil {
			panic(fmt.Errorf("could not create log folder: %w", err))
		}
	}

	// Open file based on environment
	logFile := filepath.Join(logDir, fmt.Sprintf("%s.log", env))
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("could not open log file: %w", err))
	}

	core := zapcore.NewCore(encoder, zapcore.AddSync(file), level)
	logger := zap.New(core)

	return logger.Sugar()
}
