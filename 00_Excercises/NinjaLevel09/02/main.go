package main

import "fmt"

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println(p.name, "says hi!")
}

func main() {
	p1 := person{name: "Foo"}
	p1.speak()
	saySomething(&p1)

	p2 := &person{name: "Bar"}
	p2.speak()
	saySomething(p2)
}
