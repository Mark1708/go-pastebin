package logger

import (
	"fmt"
	"os"

	"github.com/Mark1708/go-pastebin/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	log *zap.Logger
}

func SetDefaultZapLogger() {
	stdout := zapcore.AddSync(os.Stdout)
	level := getLevel("INFO")
	consoleEncoderCore := zapcore.NewConsoleEncoder(encoderConfig(true))
	core := zapcore.NewCore(consoleEncoderCore, stdout, level)

	log := zap.New(core)

	SetLogger(&zapLogger{log: log})
}

func SetZapLogger(cfg config.LoggerConfig) {
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Filename(),
		MaxSize:    cfg.MaxSize(), // megabytes
		MaxBackups: cfg.MaxBackups(),
		MaxAge:     cfg.MaxAge(), // days
	})

	level := getLevel(cfg.Level())

	fileEncoderCore := zapcore.NewJSONEncoder(encoderConfig(false))
	consoleEncoderCore := zapcore.NewConsoleEncoder(encoderConfig(true))

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoderCore, stdout, level),
		zapcore.NewCore(fileEncoderCore, file, level),
	)

	log := zap.New(core)

	SetLogger(
		&zapLogger{
			log: log.With(
				zap.String("service", cfg.ServiceName()),
			),
		},
	)
}

func (z *zapLogger) Debug(msg string) {
	z.log.Debug(msg + " ")
}

func (z *zapLogger) Debugf(template string, args ...interface{}) {
	z.Debug(fmt.Sprintf(template, args...))
}

func (z *zapLogger) Error(msg string) {
	z.log.Error(msg + " ")
}

func (z *zapLogger) Errorf(template string, args ...interface{}) {
	z.Error(fmt.Sprintf(template, args...))
}

func (z *zapLogger) Fatal(msg string) {
	z.log.Fatal(msg + " ")
}

func (z *zapLogger) Fatalf(template string, args ...interface{}) {
	z.Fatal(fmt.Sprintf(template, args...))
}

func (z *zapLogger) Info(msg string) {
	z.log.Info(msg + " ")
}

func (z *zapLogger) Infof(template string, args ...interface{}) {
	z.Info(fmt.Sprintf(template, args...))
}

func (z *zapLogger) Panic(msg string) {
	z.log.Panic(msg + " ")
}

func (z *zapLogger) Panicf(template string, args ...interface{}) {
	z.Panic(fmt.Sprintf(template, args...))
}

func (z *zapLogger) Warn(msg string) {
	z.log.Warn(msg + " ")
}

func (z *zapLogger) Warnf(template string, args ...interface{}) {
	z.Warn(fmt.Sprintf(template, args...))
}

func (z *zapLogger) With(args ...interface{}) Logger {
	z.log = z.log.With(toFields(args...)...)
	return z
}

func (z *zapLogger) Close() error {
	return z.log.Sync()
}

func toFields(args ...interface{}) []zap.Field {
	if args != nil {
		size := len(args)
		fields := make([]zap.Field, size)
		for i := 0; i < size; i++ {
			var ok bool
			fields[i], ok = args[i].(zap.Field)
			if !ok {
				Log.Fatal("incorrect field type")
			}
		}
		return fields
	}
	return []zap.Field{}
}

func encoderConfig(isColored bool) zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = "timestamp"
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeDuration = zapcore.MillisDurationEncoder
	if isColored {
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	return cfg
}

func getLevel(levelStr string) zapcore.Level {
	switch levelStr {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "PANIC":
		return zapcore.PanicLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
