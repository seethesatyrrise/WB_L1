package main

import "fmt"

func main() {
	arr := []int{2, 5, 9, 10, 12, 15, 20, 28, 30, 45}
	target := 9
	res := binarySearch(arr, target)
	if res >= 0 {
		fmt.Println("num ", target, " placed at pos #", binarySearch(arr, target))
	} else {
		fmt.Println("not found")
	}
}

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := low + (high-low)/2

		if arr[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != target {
		return -1
	}

	return low
}
