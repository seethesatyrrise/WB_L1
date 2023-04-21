package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	sleep1(3)
	fmt.Println(time.Since(t), " have passed")

	t = time.Now()
	sleep2(3)
	fmt.Println(time.Since(t), " have passed")
}

func sleep1(n int) {
	select {
	case <-time.After(time.Duration(n) * time.Second):
		return
	}
}

func sleep2(n int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(n))
	defer cancel()
	select {
	case <-ctx.Done():
		return
	}
}
