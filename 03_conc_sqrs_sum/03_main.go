package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	input := []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	fmt.Println(calculateSum1(input))
	fmt.Println(calculateSum2(input))
	fmt.Println(calculateSum3(input))
}

// вариант для интеллигенции (через atomic)
func calculateSum1(in []int64) int64 {
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

// вариант для извращенцев (через отправку в канал)
func calculateSum2(in []int64) int64 {
	var sum int64 = 0
	ch := make(chan int64)
	wg := sync.WaitGroup{}
	for _, val := range in {
		wg.Add(1)
		go func(v int64) {
			v *= v
			ch <- v
		}(val)
	}

	go func() {
		for val := range ch {
			sum += val
			wg.Done()
		}
	}()

	wg.Wait()
	close(ch)
	return sum
}

// вариант для серийных маньяков с пожизненным сроком
// (через каналы и мьютексы)
func calculateSum3(in []int64) int64 {
	var sum int64 = 0
	ch := make(chan int64)
	mtx := &sync.Mutex{}
	wg := sync.WaitGroup{}
	for _, val := range in {
		wg.Add(1)
		go func(v int64) {
			v *= v
			ch <- v
		}(val)
	}

	go func() {
		for val := range ch {
			go func(val int64) {
				mtx.Lock()
				sum += val
				mtx.Unlock()
				wg.Done()
			}(val)
		}
	}()

	wg.Wait()
	close(ch)
	return sum
}
