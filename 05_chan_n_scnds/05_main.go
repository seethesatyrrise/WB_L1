package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	tCh := make(chan struct{})
	wg := sync.WaitGroup{}
	var n int64
	fmt.Fscanln(os.Stdin, &n)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		val := 0
		for {
			select {
			case <-tCh:
				return
			default:
				ch <- val
				fmt.Println("send value ", val)
				val++
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	go func() {
		for val := range ch {
			fmt.Println("get value ", val)
		}
	}()

	time.Sleep(time.Second * time.Duration(n))
	tCh <- struct{}{}
	wg.Wait()

}
