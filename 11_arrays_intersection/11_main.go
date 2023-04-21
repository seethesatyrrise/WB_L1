package main

import "fmt"

func main() {
	arr1 := []int{3, 62, 5, 85, 4, 2, 79}
	arr2 := []int{62, 11, 3, 8, 9, 10, 4, 7}

	fmt.Println(intersection(arr1, arr2))
}

func intersection(a1, a2 []int) []int {
	m := make(map[int]struct{})
	res := []int{}
	for _, val := range a1 {
		m[val] = struct{}{}
	}

	for _, val := range a2 {
		_, ok := m[val]
		if ok {
			res = append(res, val)
		}
	}

	return res
}
