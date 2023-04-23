package main

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_sum(t *testing.T) {
	type args struct {
		a *big.Int
		b *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"1", args{makeBigInt("2000000000"), makeBigInt("100000000")}, makeBigInt("2100000000")},
		{"2", args{makeBigInt("812373256495"), makeBigInt("23466876886954")}, makeBigInt("24279250143449")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_div(t *testing.T) {
	type args struct {
		a *big.Int
		b *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"1", args{makeBigInt("2000000000"), makeBigInt("100000000")}, makeBigInt("20")},
		{"2", args{makeBigInt("22647043290948885"), makeBigInt("634578345")}, makeBigInt("35688333")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := div(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sub(t *testing.T) {
	type args struct {
		a *big.Int
		b *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"1", args{makeBigInt("2000000000"), makeBigInt("100000000")}, makeBigInt("1900000000")},
		{"2", args{makeBigInt("812373256495"), makeBigInt("6488695433")}, makeBigInt("805884561062")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sub(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mul(t *testing.T) {
	type args struct {
		a *big.Int
		b *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"1", args{makeBigInt("2000000000"), makeBigInt("100000000")}, makeBigInt("200000000000000000")},
		{"2", args{makeBigInt("226470432"), makeBigInt("634578345")}, makeBigInt("143713231929995040")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mul(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mul() = %v, want %v", got, tt.want)
			}
		})
	}
}
