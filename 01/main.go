// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	name string
}

func (h *Human) sayHello() {
	fmt.Println("Hello!")
}

func (h *Human) sayMyName() {
	fmt.Printf("My name is %s!\n", h.name)
}

type Action struct {
	Human // Встраивание Human в Action
}

func (a *Action) doSmth() {
	fmt.Println("I'm busy doing something useful!")
}

func main() {
	var f Action
	f.name = "Edward"
	f.sayHello()
	f.sayMyName()
	f.doSmth()
}
