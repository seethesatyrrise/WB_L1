package main

import "fmt"

// Паттерн Адаптер (Adapter) предназначен для преобразования интерфейса
// одного класса в интерфейс другого. Благодаря реализации данного паттерна
// мы можем использовать вместе классы с несовместимыми интерфейсами.
//
// Когда надо использовать Адаптер?
// - Когда необходимо использовать имеющийся класс, но его интерфейс не соответствует потребностям
// - Когда надо использовать уже существующий класс совместно с другими классами, интерфейсы которых не совместимы
//
// В моем примере есть человек и животные: собака и кот. Человек умеет взаимодействовать
// с животными (гладить их). Собаку можно гладить. Но кота можно гладить только с опаской.
// Поэтому нужен адаптер, чтобы человек мог взаимодействовать с котом.

func main() {
	// создаем человека
	human := &human{}
	// создаем собаку
	dog := &dog{}

	// взаимодействуем с собакой (гладим)
	human.interaction(dog)

	// создаем кота и адаптер к нему
	cat := &cat{}
	catPetting := &catToPetAdapter{c: cat}

	// гладим кота
	human.interaction(catPetting)
}

// интерфейс животного
type animal interface {
	pet()
}

// человек
type human struct{}

func (h *human) interaction(t animal) {
	t.pet()
}

// собака
type dog struct {
	animal
}

func (d *dog) pet() {
	fmt.Println("petting dog")
}

// кошка
type cat struct {
	animal
}

func (c *cat) petCautiously() {
	fmt.Println("petting cat")
}

// адаптер
type catToPetAdapter struct {
	c *cat
	animal
}

func (a *catToPetAdapter) pet() {
	a.c.petCautiously()
}
