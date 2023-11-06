package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Config struct {
	Namespace   string
	Development bool
	Filepath    string
	Level       zapcore.Level
}

var (
	logger        *zap.Logger
	encoderConfig zapcore.EncoderConfig
	writeSyncer   zapcore.WriteSyncer
)

func InitLogger(config Config) error {
	if config.Development {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"
	encoderConfig.TimeKey = "time"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder

	if len(config.Filepath) != 0 {
		logFile, err := os.OpenFile(config.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return errors.Wrap(err, "fail to open log file")
		}
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(logFile), zapcore.AddSync(os.Stdout))
	} else {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writeSyncer,
		zap.NewAtomicLevelAt(config.Level),
	)

	logger = zap.New(core, zap.AddStacktrace(zap.ErrorLevel))

	logger = logger.With(zap.String("namespace", config.Namespace))
	defer logger.Sync()

	return nil
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}
