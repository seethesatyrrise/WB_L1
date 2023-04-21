package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []int{2, 4, 6, 8, 10, 12}
	calculateSquares(input)
}

func calculateSquares(in []int) {
	wg := sync.WaitGroup{}
	for _, val := range in {
		wg.Add(1)
		go func(v int) {
			fmt.Println(v * v)
			wg.Done()
		}(val)
	}
	wg.Wait()
}
