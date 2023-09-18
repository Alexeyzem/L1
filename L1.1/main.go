// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)

package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) GetName() string {
	return h.name
}
func (h *Human) GetAge() int {
	return h.age
}

type Action struct { // Action получает все поля и методы как и Human
	Human
}

func (a *Action) GetNameAndAge() (string, int) {
	return a.name, a.age
}

func main() {
	var h Human
	h.age = 10
	h.name = "ugiv"
	var a Action
	a.age = 100
	a.name = "ionoiononn"
	fmt.Println(a.GetAge(), a.GetName(), h.GetAge())
	fmt.Println(a.GetNameAndAge())
}
