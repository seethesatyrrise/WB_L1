package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
}

func main() {
	a := &Action{}
	a.SetName("Hazel")
	a.SetAge(2)
	fmt.Println(a)
}

func (h *Human) SetName(name string) {
	h.name = name
}

func (h *Human) SetAge(age int) {
	if age >= 0 {
		h.age = age
	}
}
