package zapwrap

import (
	"time"

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
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
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

// wrap zap logger Fatal function; will trigger an os.exit(1) call
func Fatal(message string, fields ...zap.Field) {
	zl.Fatal(message, fields...)
}
