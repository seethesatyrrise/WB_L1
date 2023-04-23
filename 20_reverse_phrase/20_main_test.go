package main

import "testing"

func Test_reversePhrase(t *testing.T) {
	type args struct {
		phrase string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"snow dog sun morning"}, "morning sun dog snow"},
		{"2", args{""}, ""},
		{"3", args{"cat sleeping on table"}, "table on sleeping cat"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePhrase(tt.args.phrase); got != tt.want {
				t.Errorf("reversePhrase() = %v, want %v", got, tt.want)
			}
		})
	}
}
