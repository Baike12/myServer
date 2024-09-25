package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

var Logger *zap.Logger

var logFilePath string = "../log/myServer.log"

func InitLogger() error {
	if err := os.MkdirAll(filepath.Dir(logFilePath), 0755); err != nil {
		return err
	}

	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Encoding:          "json",
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       []string{file.Name()},
		ErrorOutputPaths:  []string{"stderr"},
		DisableCaller:     false,
		DisableStacktrace: true,
	}

	logger, err := config.Build()
	if err != nil {
		return err
	}

	Logger = logger

	return nil
}

func Sync() error {
	return Logger.Sync()
}

func InfoLog(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func DebugLog(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func WarnLog(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func ErrorLog(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}
