package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
