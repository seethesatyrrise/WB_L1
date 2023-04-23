package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := makeBigInt("2000000000")
	b := makeBigInt("1000000000")

	fmt.Println(sum(a, b))
	fmt.Println(div(a, b))
	fmt.Println(mul(a, b))
	fmt.Println(sub(a, b))
}

func makeBigInt(s string) *big.Int {
	res := &big.Int{}
	res.SetString(s, 10)
	return res
}

func sum(a, b *big.Int) *big.Int {
	res := &big.Int{}
	res.Add(a, b)
	return res
}

func div(a, b *big.Int) *big.Int {
	res := &big.Int{}
	res.Div(a, b)
	return res
}

func sub(a, b *big.Int) *big.Int {
	res := &big.Int{}
	res.Sub(a, b)
	return res
}

func mul(a, b *big.Int) *big.Int {
	res := &big.Int{}
	res.Mul(a, b)
	return res
}
