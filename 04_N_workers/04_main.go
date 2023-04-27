package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGINT)

	var n int
	fmt.Fscanln(os.Stdin, &n)

	ch := make(chan int)
	defer close(ch)
	quit := make(chan struct{})
	defer close(quit)

	runWorkers(n, ch)

	val := 0
	for {
		val++
		//time.Sleep(time.Millisecond * 100)
		select {
		// отлов сигнала о прерывании
		case <-signalChan:
			return
		default:

			ch <- val
		}
	}
}

// запуск N воркеров
func runWorkers(n int, ch chan int) {
	for i := 0; i < n; i++ {
		go func(i int) {
			for val := range ch {
				fmt.Printf("goroutine #%d\tval: %d\n", i, val)
			}
		}(i)
	}
}
