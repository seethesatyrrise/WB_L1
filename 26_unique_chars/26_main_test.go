package main

import "testing"

func Test_isUnique(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"asdfghjklxcvbnm"}, true},
		{"2", args{"aaaaaaaaaaaa"}, false},
		{"3", args{"asdfghjklSxcvbnm"}, false},
		{"4", args{"AgtyBunRjO"}, true},
		{"5", args{"asdfghjASDFGHJ"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique(tt.args.s); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
