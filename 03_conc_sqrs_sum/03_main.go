package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	input := []int64{2, 4, 6, 8, 10}

	fmt.Println(calculateSum(input))
}

func calculateSum(in []int64) int64 {
	var sum int64 = 0
	wg := sync.WaitGroup{}
	for _, val := range in {
		wg.Add(1)
		go func(v int64) {
			v *= v
			atomic.AddInt64(&sum, v)
			wg.Done()
		}(val)
	}
	wg.Wait()
	return sum
}
