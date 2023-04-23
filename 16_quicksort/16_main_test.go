package main

import (
	"reflect"
	"testing"
)

func Test_quickSortStart(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{}}, []int{}},
		{"2", args{[]int{4, 29, 6, 13, 644, 36, 75, 49}}, []int{4, 6, 13, 29, 36, 49, 75, 644}},
		{"3", args{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"4", args{[]int{63, -46, 2, 56, -3}}, []int{-46, -3, 2, 56, 63}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSortStart(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSortStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
