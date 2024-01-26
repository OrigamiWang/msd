package logutil

import (
	"testing"
)

func TestInfo(t *testing.T) {
	type args struct {
		arg0 interface{}
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				arg0: "info...",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.arg0, tt.args.args...)
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		arg0 interface{}
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				arg0: "debug...",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.arg0, tt.args.args...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		arg0 interface{}
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				arg0: "error...",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.arg0, tt.args.args...)
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		arg0 interface{}
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.arg0, tt.args.args...)
		})
	}
}
