package main

import "fmt"

type Animal interface {
	Pet()
	Name() string
}

type Cat struct {
	name string
}

func (c Cat) Pet() {
	fmt.Println("prrrrrr")
}

func (c Cat) Name() string {
	return c.name
}

type Dog struct {
	name string
}

func (d Dog) Pet() {
	fmt.Println("woof woof")
}

func (d Dog) Name() string {
	return d.name
}

func Compliment(a Animal) {
	fmt.Println("Great Job", a.Name())
	a.Pet()
}

func main() {
	c := Cat{"johnny larry"}
	Compliment(c)

	d := Dog{"trixie belle"}
	Compliment(d)
}
