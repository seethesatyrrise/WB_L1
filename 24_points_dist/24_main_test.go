package main

import "testing"

func Test_point_distanceTo(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		p2 *point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"1", fields{0, 3}, args{&point{4, 0}}, 5},
		{"2", fields{2, 5}, args{&point{2, 8}}, 3},
		{"3", fields{3, 1}, args{&point{10, 1}}, 7},
		{"4", fields{4, 6}, args{&point{4, 6}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1 := &point{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := p1.distanceTo(tt.args.p2); got != tt.want {
				t.Errorf("distanceTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
