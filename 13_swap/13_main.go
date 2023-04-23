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

// метод счастливых go-шников
func swap1(a, b *int) {
	*a, *b = *b, *a
}

// метод для грустных людей, использующих языки,
// в которых нет множественного присваивания
func swap2(a, b *int) {
	*a = *a ^ *b
	*b = *b ^ *a
	*a = *a ^ *b
}
