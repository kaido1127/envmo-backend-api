package logger

import (
	"envmo/app_config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugaredLogger *zap.SugaredLogger

//var isSugarType bool

func CreateLoggerByConfig(cfg app_config.LoggerConfig) error {
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return err
	}

	config := zap.Config{
		Development:       cfg.Development,
		DisableCaller:     cfg.DisableCaller,
		DisableStacktrace: cfg.DisableStacktrace,
		Encoding:          cfg.Encoding,
		Level:             zap.NewAtomicLevelAt(level),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		EncoderConfig:     zap.NewProductionEncoderConfig(),
	}
	zapLogger, err := config.Build()
	if err != nil {
		return err
	}
	sugaredLogger = zapLogger.Sugar()
	//isSugarType = cfg.SugarType
	return nil
}

func log(level zapcore.Level, msg string, fields ...interface{}) {
	sugaredLogger.With(fields...).Log(level, msg)
}

func Debug(msg string, fields ...interface{}) {
	log(zapcore.DebugLevel, msg, fields...)
}

func Info(msg string, fields ...interface{}) {
	log(zapcore.InfoLevel, msg, fields...)
}

func Warn(msg string, fields ...interface{}) {
	log(zapcore.WarnLevel, msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	log(zapcore.ErrorLevel, msg, fields...)
}

func Fatal(msg string, fields ...interface{}) {
	log(zapcore.FatalLevel, msg, fields...)
}

func Sync() {
	sugaredLogger.Sync()
}
