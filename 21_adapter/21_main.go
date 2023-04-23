package main

import "fmt"

func main() {
	// путешественник
	drvr := &driver{}
	auto := &auto{}

	drvr.travel(auto)

	camel := &camel{}
	camelTransport := &camelToTransportAdapter{c: camel}

	// продолжаем путь по пескам пустыни
	drvr.travel(camelTransport)
}

// интерфейс транспорта
type transport interface {
	drive()
}

// путешественник
type driver struct {
}

func (d *driver) travel(t transport) {
	t.drive()
}

// машина
type auto struct {
	transport
}

func (a *auto) drive() {
	fmt.Println("driving auto")
}

// верблюд
type camel struct {
	transport
}

func (c *camel) move() {
	fmt.Println("moving camel")
}

// адаптер
type camelToTransportAdapter struct {
	c *camel
	transport
}

func (a *camelToTransportAdapter) drive() {
	a.c.move()
}
