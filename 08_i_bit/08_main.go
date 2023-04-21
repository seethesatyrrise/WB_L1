package main

import "fmt"

func main() {
	var value int64
	value = 10
	fmt.Printf("%b\t%d\n", value, value) // 0b1010, 10
	value = set1(value, 4)
	fmt.Printf("%b\t%d\n", value, value) // 0b11010, 26
	value = set0(value, 3)
	fmt.Printf("%b\t%d\n", value, value) // 0b10010, 18
}

func set1(val int64, n int64) int64 {
	return val | 1<<n
}

func set0(val int64, n int64) int64 {
	return val &^ (1 << n)
}
