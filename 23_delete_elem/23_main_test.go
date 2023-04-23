package main

import (
	"reflect"
	"testing"
)

func Test_deleteElement(t *testing.T) {
	type args struct {
		in []int
		i  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5, 6}, 0}, []int{2, 3, 4, 5, 6}},
		{"2", args{[]int{1, 2, 3, 4, 5, 6}, 5}, []int{1, 2, 3, 4, 5}},
		{"3", args{[]int{1, 2, 3, 4, 5, 6}, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"4", args{[]int{1, 2, 3, 4, 5, 6}, -1}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteElement(tt.args.in, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
