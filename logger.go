package tiologger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var (
	log logger
)

type tioLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

type logger struct {
	log *zap.Logger
}

// init, initialises zap logger
func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutput()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func getLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(envLogLevel))) {
	case "-1":
		return zap.DebugLevel
	case "0":
		return zap.InfoLevel
	case "1":
		return zap.WarnLevel
	case "2":
		return zap.ErrorLevel
	case "3":
		return zap.DPanicLevel
	case "4":
		return zap.PanicLevel
	case "5":
		return zap.FatalLevel
	default:
		return zap.ErrorLevel
	}
}

func getOutput() string {
	output := strings.TrimSpace(os.Getenv(envLogOutput))
	if output == "" {
		return "stdout"
	}
	return output
}

// GetLogger, returns *zap.Logger
func GetLogger() tioLogger {
	return log
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

// Debug, creates log entry under "debug". Requires string
func Debug(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	_ = log.log.Sync()
	// if err := log.log.Sync(); err != nil {
	// 	// TODO: update when https://github.com/uber-go/zap/issues/1000 is fixed
	// 	// log.log.Error("Debug sync error:  "+err.Error(), tags...)
	// }
}

// Info, creates log entry under "info". Requires string
func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	_ = log.log.Sync()
	// if err := log.log.Sync(); err != nil {
	// 	// TODO: update when https://github.com/uber-go/zap/issues/1000 is fixed
	// 	// log.log.Error("Info sync error:  "+err.Error(), tags...)
	// }
}

// Info, creates log entry under "info". Requires string
func Warn(msg string, tags ...zap.Field) {
	log.log.Warn(msg, tags...)
	_ = log.log.Sync()
	// if err := log.log.Sync(); err != nil {
	// 	// TODO: update when https://github.com/uber-go/zap/issues/1000 is fixed
	// 	// log.log.Error("Warn sync error:  "+err.Error(), tags...)
	// }
}

// Error, creates log entry under "ERROR". Requires string and error
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("ERROR", err))
	log.log.Error(msg, tags...)
	_ = log.log.Sync()
	// if err := log.log.Sync(); err != nil {
	// 	err = nil // TODO: update when https://github.com/uber-go/zap/issues/1000 is fixed
	// 	// log.log.Error("Error sync error:  "+err.Error(), tags...)
	// }
}

// Critical, creates log entry under "CRITICAL". Requires string and error
func Critical(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("CRITICAL", err))
	log.log.Error(msg, tags...)
	_ = log.log.Sync()
	// if err := log.log.Sync(); err != nil {
	// 	err = nil // TODO: update when https://github.com/uber-go/zap/issues/1000 is fixed
	// 	// log.log.Error("Critical sync error:  "+err.Error(), tags...)
	// }
}

// DPanic, creates log entry under "D-PANIC". Requires string and error
func DPanic(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("D-PANIC", err))
	log.log.DPanic(msg, tags...)
	_ = log.log.Sync()
	// if err := log.log.Sync(); err != nil {
	// 	err = nil // TODO: update when https://github.com/uber-go/zap/issues/1000 is fixed
	// 	// log.log.Error("DPanic sync error:  "+err.Error(), tags...)
	// }
}

// Panic, creates log entry under "PANIC". Requires string and error
func Panic(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("PANIC", err))
	log.log.Panic(msg, tags...)
	_ = log.log.Sync()
	// if err := log.log.Sync(); err != nil {
	// 	err = nil // TODO: update when https://github.com/uber-go/zap/issues/1000 is fixed
	// 	// log.log.Error("Panic sync error:  "+err.Error(), tags...)
	// }
}