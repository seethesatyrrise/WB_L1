package main

import "testing"

func Test_set0(t *testing.T) {
	type args struct {
		val int64
		n   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{26, 3}, 18},
		{"2", args{26, 10}, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := set0(tt.args.val, tt.args.n); got != tt.want {
				t.Errorf("set0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_set1(t *testing.T) {
	type args struct {
		val int64
		n   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{10, 4}, 26},
		{"2", args{10, 0}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := set1(tt.args.val, tt.args.n); got != tt.want {
				t.Errorf("set1() = %v, want %v", got, tt.want)
			}
		})
	}
}
