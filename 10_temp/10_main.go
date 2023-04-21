package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(getRanges(temps))
}

func getRanges(temps []float64) map[float64][]float64 {
	m := make(map[float64][]float64)
	for _, temp := range temps {
		key := math.Floor(temp/10) * 10
		values := m[key]
		values = append(values, temp)
		m[key] = values
	}
	return m
}
