package crypto

import "testing"

func TestMd5Encode(t *testing.T) {
	type args struct {
		rawStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{rawStr: "abc123"},
			want: "e99a18c428cb38d5f260853678922e03",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Encode(tt.args.rawStr); got != tt.want {
				t.Errorf("Md5Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
