package main

import (
	"reflect"
	"testing"
)

func Test_doSet(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{"1", args{[]string{"cat", "cat", "dog", "cat", "tree"}},
			map[string]struct{}{"cat": struct{}{}, "dog": struct{}{}, "tree": struct{}{}}},
		{"2", args{[]string{}},
			map[string]struct{}{}},
		{"3", args{[]string{"cat", "fish", "dog", "tree"}},
			map[string]struct{}{"cat": struct{}{}, "fish": struct{}{}, "dog": struct{}{}, "tree": struct{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doSet(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
