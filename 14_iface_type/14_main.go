package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		a int
		b string
		c bool
		d chan int
	)
	fmt.Println(getType(a))
	fmt.Println(getType(b))
	fmt.Println(getType(c))
	fmt.Println(getType(d))
	printType(a)
	printType(b)
	printType(c)
	printType(d)
}

// вариант с рефлексией
func getType(val interface{}) string {
	return reflect.TypeOf(val).String()
}

// вариант для демонстрации конструкции type-switch,
// в целом не имеющей особого смысла в данном случае
func printType(val interface{}) {
	var res string
	switch val.(type) {
	case int:
		res = "int"
	case string:
		res = "string"
	case bool:
		res = "bool"
	case chan int, chan string, chan bool:
		res = "chan"
	default:
		res = "undefined"
	}
	fmt.Println("value type is", res)
}
