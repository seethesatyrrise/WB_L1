package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []int{2, 4, 6, 8, 10, 12}
	wg := sync.WaitGroup{}
	for _, val := range input {
		wg.Add(1)
		go func(v int) {
			fmt.Println(v * v)
			wg.Done()
		}(val)
	}
	wg.Wait()
}
