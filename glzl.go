package gzl

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zap logger global variable
var zl *zap.Logger

func init() {
	// initialise a zap logger
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	// configure zap logger encoding
	// configure timestamp format
	zapcore.TimeEncoderOfLayout("ISO8601")
	// hide stack trace information
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	// build zap logger configuration
	zl, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

// wrap zap logger Info function
func Info(message string, fields ...zap.Field) {
	zl.Info(message, fields...)
}

// wrap zap logger Debug function
func Debug(message string, fields ...zap.Field) {
	zl.Debug(message, fields...)
}

// wrap zap logger Error function
func Error(message string, fields ...zap.Field) {
	zl.Error(message, fields...)
}

// wrap zap logger Fatal function
func Fatal(message string, fields ...zap.Field) {
	zl.Fatal(message, fields...)
}

// wrap zap logger Panic function
func Panic(message string, fields ...zap.Field) {
	zl.Panic(message, fields...)
}
