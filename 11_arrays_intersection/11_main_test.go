package main

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		a1 []int
		a2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 6, 4, 1, 5, 2}, []int{1, 5, 3, 6, 4, 2}}, []int{1, 5, 3, 6, 4, 2}},
		{"2", args{[]int{1, 2, 3, 4, 5, 6}, []int{8, 9, 10, 11, 12}}, []int{}},
		{"3", args{[]int{}, []int{}}, []int{}},
		{"4", args{[]int{3, 6, 4, 1, 5, 2}, []int{8, 3, 1, 11}}, []int{3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.a1, tt.args.a2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
