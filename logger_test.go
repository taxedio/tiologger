package tiologger

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"go.uber.org/zap"
)

var (
	regExp = `&zap\.Logger{core:\(\*zapcore\.ioCore\)\([a-zA-Z\d]{12}\), development:false, addCaller:true, onFatal:0x0, name:"", errorOutput:zapcore.writerWrapper{Writer:io.discard{}}, addStack:2, callerSkip:0, clock:zapcore.systemClock{}}`
)

func TestGetLogger(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Get Logger", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetLogger()
			exp := fmt.Sprintf("%#v", got)
			match, _ := regexp.MatchString(regExp, exp)
			if !match {
				t.Errorf("GetLogger() = %v, want %v", exp, tt.want)
			}

		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		msg  string
		tags []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{msg: "test", tags: []zap.Field{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.msg, tt.args.tags...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		msg  string
		err  error
		tags []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{msg: "test", err: errors.New("test"), tags: []zap.Field{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.msg, tt.args.err, tt.args.tags...)
		})
	}
}

func TestCritical(t *testing.T) {
	type args struct {
		msg  string
		err  error
		tags []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{msg: "test", err: errors.New("test"), tags: []zap.Field{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Critical(tt.args.msg, tt.args.err, tt.args.tags...)
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		msg  string
		tags []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{msg: "test", tags: []zap.Field{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.msg, tt.args.tags...)
		})
	}
}

func TestWarn(t *testing.T) {
	GetLogger()
	type args struct {
		msg  string
		tags []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{msg: "test", tags: []zap.Field{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.msg, tt.args.tags...)
		})
	}
}

func TestDPanic(t *testing.T) {
	type args struct {
		msg  string
		err  error
		tags []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{msg: "test", err: errors.New("test"), tags: []zap.Field{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DPanic(tt.args.msg, tt.args.err, tt.args.tags...)
		})
	}
}

func TestPanic(t *testing.T) {
	type args struct {
		msg  string
		err  error
		tags []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{msg: "test", err: errors.New("test"), tags: []zap.Field{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Panic(tt.args.msg, tt.args.err, tt.args.tags...)
		})
	}
}
