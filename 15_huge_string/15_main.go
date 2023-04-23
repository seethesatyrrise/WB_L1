package main

import "fmt"

func main() {
	v := "ddfrrffffgfdbtrebrdsg"
	fmt.Println(someFunc(v))
	v = "00000000001000000000200000000030000000004000000000500000000060000000007000000000800000000090000000001000000000"
	fmt.Println(someFunc(v))
}

/*
вариант с v[:100] упадет в случаях,
когда длина строки меньше 100.
в данной реализации строка режется
либо до 100, либо возвращается вся,
если ее длины не хватает
*/
func someFunc(hugeStr string) string {
	right := 100
	if len(hugeStr) < right {
		return hugeStr
	}
	return hugeStr[:right]
}
