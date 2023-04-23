package main

import (
	"fmt"
)

func main() {
	arr := []int{93, 3, 64, 11, 8, 45}
	runConveyor(arr)
}

// запуск конвейера
// создает нужные каналы и запускает все
// необходимые горутины
func runConveyor(data []int) {
	firstChan := make(chan int)
	secondChan := make(chan int)
	done := make(chan struct{})
	go firstStep(data, firstChan)
	go secondStep(firstChan, secondChan)
	go printData(secondChan, done)
	<-done
}

// отправка данных из массива в первый канал
func firstStep(data []int, outCh chan int) {
	for _, val := range data {
		outCh <- val
	}
	close(outCh)
}

// принятие данных из первого канала, вычисление 2*x,
// отправка значения во второй канал
func secondStep(inCh, outCh chan int) {
	for val := range inCh {
		outCh <- doMath(val)
	}
	close(outCh)
}

// вычисление 2*x
func doMath(val int) int {
	return val * 2
}

// принятие данных из второго канала и их вывод
func printData(inCh chan int, done chan struct{}) {
	for val := range inCh {
		fmt.Println(val)
	}
	done <- struct{}{}
}
