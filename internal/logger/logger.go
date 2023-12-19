package logger

import (
	"bytes"
	"flag"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logz *zap.Logger

func New() {
	dcCore := newDiscordCore()
	csCore := newConsoleCore()
	logz = zap.New(zapcore.NewTee(dcCore, csCore), zap.AddCaller(), zap.AddCallerSkip(1))

	defer logz.Sync()
	Info("[LOGGER] 😎 logger has been initialize 🙄")
}

func newDiscordCore() zapcore.Core {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.StacktraceKey = ""

	return zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(&discordSink{}), zapcore.ErrorLevel)
}

func newConsoleCore() zapcore.Core {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.StacktraceKey = ""
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(os.Stdout), zap.InfoLevel)
}

type discordSink struct{}

func (s *discordSink) Write(p []byte) (n int, err error) {
	if flag.Lookup("test.v") != nil {
		return
	}
	go http.Post(os.Getenv("LOG_SERVER"), "application/json", bytes.NewBuffer([]byte(string(p))))
	return len(p), nil
}

func Info(message string, fields ...zap.Field) {
	logz.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	logz.Debug(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logz.Warn(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	if flag.Lookup("test.v") != nil {
		return
	}
	switch v := message.(type) {
	case error:
		logz.Error(v.Error(), fields...)
	case string:
		logz.Error(v, fields...)
	}
}
