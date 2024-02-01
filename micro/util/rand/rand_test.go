package main

import (
	"fmt"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				length: 32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRandomString(tt.args.length)
			fmt.Println(got)
		})
	}
}
