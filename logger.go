package tiologger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zapcore.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel = zapcore.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zapcore.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zapcore.ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = zapcore.DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel = zapcore.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zapcore.FatalLevel
)

// init, initialises zap logger
func init() {
	var (
		err error
	)
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

// GetLogger, returns *zap.Logger
func GetLogger() *zap.Logger {
	return log
}

// Debug, creates log entry under "debug". Requires string
func Debug(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	if err := log.Sync(); err != nil {
		log.Error(err.Error(), tags...)
	}
}

// Info, creates log entry under "info". Requires string
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	if err := log.Sync(); err != nil {
		log.Error(err.Error(), tags...)
	}
}

// Info, creates log entry under "info". Requires string
func Warn(msg string, tags ...zap.Field) {
	log.Warn(msg, tags...)
	if err := log.Sync(); err != nil {
		log.Error(err.Error(), tags...)
	}
}

// Error, creates log entry under "ERROR". Requires string and error
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("ERROR", err))
	log.Error(msg, tags...)
	if err := log.Sync(); err != nil {
		log.Error(err.Error(), tags...)

	}
}

// Critical, creates log entry under "CRITICAL". Requires string and error
func Critical(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("CRITICAL", err))
	log.Error(msg, tags...)
	if err := log.Sync(); err != nil {
		log.Error(err.Error(), tags...)
	}
}

// DPanic, creates log entry under "D-PANIC". Requires string and error
func DPanic(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("D-PANIC", err))
	log.Error(msg, tags...)
	if err := log.Sync(); err != nil {
		log.Error(err.Error(), tags...)

	}
}

// Panic, creates log entry under "PANIC". Requires string and error
func Panic(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("PANIC", err))
	log.Error(msg, tags...)
	if err := log.Sync(); err != nil {
		log.Error(err.Error(), tags...)

	}
}
