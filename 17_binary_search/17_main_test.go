package main

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 5, 9, 10, 12, 15, 20, 28, 30, 45}, 4}, -1},
		{"2", args{[]int{2, 5, 9, 10, 12, 15, 20, 28, 30, 45}, 2}, 0},
		{"3", args{[]int{2, 5, 9, 10, 12, 15, 20, 28, 30, 45}, 45}, 9},
		{"4", args{[]int{2, 5, 9, 10, 12, 15, 20, 28, 30, 45}, 10}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
