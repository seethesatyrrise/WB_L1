package main

import (
	"reflect"
	"testing"
)

func Test_getRanges(t *testing.T) {
	type args struct {
		temps []float64
	}
	tests := []struct {
		name string
		args args
		want map[int][]float64
	}{
		{"1", args{[]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}},
			map[int][]float64{-30: {-25.4, -27, -21},
				10: {13, 19, 15.5}, 20: {24.5}, 30: {32.5}}},
		{"2", args{[]float64{-29.9, -20.1, -30.0, -20.0, 10.0, 29.9}},
			map[int][]float64{-30: {-29.9, -20.1, -30.0},
				10: {10.0}, 20: {29.9}, -20: {-20.0}}},
		{"3", args{[]float64{}},
			map[int][]float64{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRanges(tt.args.temps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
