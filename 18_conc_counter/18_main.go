package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	value int64
}

func (c *counter) increment() {
	atomic.AddInt64(&c.value, 1)
}

func main() {
	cntr := &counter{value: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("goroutine #", i)
			cntr.increment()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(cntr.value)
}
