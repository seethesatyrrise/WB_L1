package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(deleteElement(arr, 2))
}

func deleteElement(in []int, i int) []int {
	if i < 0 || i > len(in)-1 {
		return in
	}
	return append(in[:i], in[i+1:]...)
}
