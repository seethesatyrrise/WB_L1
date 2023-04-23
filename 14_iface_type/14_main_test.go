package main

import (
	"reflect"
	"testing"
)

func Test_getType(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"ksvisortb"}, "string"},
		{"2", args{make(chan int)}, "chan int"},
		{"3", args{235}, "int"},
		{"4", args{false}, "bool"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getType(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}
