package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println(a, b)
	swap1(&a, &b)
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)
}

func swap1(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b *int) {
	*a = *a ^ *b
	*b = *b ^ *a
	*a = *a ^ *b
}
