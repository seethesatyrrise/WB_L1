package main

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"абырвалг"}, "главрыба"},
		{"2", args{""}, ""},
		{"3", args{"猫の睡眠"}, "眠睡の猫"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.str); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
