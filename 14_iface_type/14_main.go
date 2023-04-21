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
}

func getType(val interface{}) reflect.Type {
	return reflect.TypeOf(val)
}
