package configs

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger    *zap.Logger
	zapConfig zapcore.EncoderConfig
)

func init() {
	// Logger, _ = zap.NewProduction()
	zapConfig = zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(zapConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, os.Stdout, zapcore.DebugLevel),
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	// Logger.Info("📝 Initializing the zap logger ...")
}
