package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ch := make(chan int)

	var n int64
	fmt.Fscanln(os.Stdin, &n)

	ctx, cancel := context.WithCancel(context.Background())

	go receive(ctx, ch)
	go send(ctx, ch)

	time.Sleep(time.Second * time.Duration(n))
	cancel()
	close(ch)
}

func send(ctx context.Context, ch chan int) {
	val := 0
	for {
		ch <- val
		fmt.Println("send ", val)
		val++
		time.Sleep(time.Millisecond * 100)
	}
}

func receive(ctx context.Context, ch chan int) {
	for val := range ch {
		fmt.Println("get ", val)
	}
}
