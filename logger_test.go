package tiologger

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	regExp = `logger\.logger{log:\(\*zap.Logger\)\([a-zA-Z\d]{12}\)}`
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
			defer func() { recover() }()
			Panic(tt.args.msg, tt.args.err, tt.args.tags...)
		})
	}
}

func Test_logger_Printf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		l    logger
		args args
	}{
		{"Test 1", logger{}, args{format: "", v: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Printf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_logger_Print(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		l    logger
		args args
	}{
		{"Test 1", logger{}, args{v: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Print(tt.args.v...)
		})
	}
}

func Test_getLevel(t *testing.T) {
	tests := []struct {
		name string
		want zapcore.Level
	}{
		{"test", zap.ErrorLevel},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLevel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
