package main

import "testing"

func Test_someFunc(t *testing.T) {
	type args struct {
		hugeStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"00000000001000000000200000000030000000004000000000500000000060000000007000000000800000000090000000001000000000"},
			"0000000000100000000020000000003000000000400000000050000000006000000000700000000080000000009000000000"},
		{"2", args{""}, ""},
		{"3", args{"3456786534567"}, "3456786534567"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := someFunc(tt.args.hugeStr); got != tt.want {
				t.Errorf("someFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
