package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	stopByTimeout(3)
	stopByCancel(4)
	stopByChan(5)
}

// функция с возвратом по сигналу завершения контекста
func stopWithCtx(ctx context.Context) {
	t := time.Now()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stopped by ctx after %fs\n", time.Since(t).Seconds())
			return
		}
	}
}

// функция с возвратом по сигналу из канала
func stopWithChan(ch chan struct{}) {
	t := time.Now()
	for {
		select {
		case <-ch:
			fmt.Printf("stopped by channel after %fs\n", time.Since(t).Seconds())
			return
		}
	}
}

// запуск горутины с отменой выполнения по истечению n секунд
// использует контекст WithCancel
// отмена происходит при вызове функции отмены cancel
func stopByCancel(n int) {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("start goroutine with cancelling")
	go stopWithCtx(ctx)

	time.Sleep(time.Second * time.Duration(n))
	cancel()
	time.Sleep(time.Second)
}

// запуск горутины с отменой выполнения по истечению n секунд
// использует контекст WithTimeout
// также есть вариант с контекстом WithDeadline, но они взаимозаменяемы
func stopByTimeout(n int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(n))

	fmt.Println("start goroutine with timeout")
	go stopWithCtx(ctx)

	time.Sleep(time.Second*time.Duration(n) + 2)
	cancel()
	time.Sleep(time.Second)
}

// запуск горутины с отменой выполнения по истечению n секунд
// использует канал, в который отправляется любое значение
// для отмены выполнения
func stopByChan(n int) {
	ch := make(chan struct{})

	fmt.Println("start goroutine with stop channel")
	go stopWithChan(ch)

	time.Sleep(time.Second * time.Duration(n))
	ch <- struct{}{}
	time.Sleep(time.Second)
}
