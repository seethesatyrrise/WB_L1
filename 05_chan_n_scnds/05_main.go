package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ch := make(chan int)

	// количество секунд работы пересылки
	var n int64
	fmt.Fscanln(os.Stdin, &n)

	ctx, cancel := context.WithCancel(context.Background())

	go receive(ctx, ch)
	go send(ctx, ch)

	time.Sleep(time.Second * time.Duration(n))
	cancel()
	close(ch)
}

// отправляем данные в канал
func send(ctx context.Context, ch chan int) {
	val := 0
	for {
		select {
		case <-ctx.Done():
			// выход по отмене контекста
			return
		default:
			ch <- val
			fmt.Println("send ", val)
			val++
			time.Sleep(time.Millisecond * 100)
		}
	}
}

// принимаем данные из канала и выводим
func receive(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			// выход по отмене контекста
			return
		case val := <-ch:
			fmt.Println("get ", val)
		}
	}
}
