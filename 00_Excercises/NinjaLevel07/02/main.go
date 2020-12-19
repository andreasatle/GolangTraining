package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person, name string) {
	p.name = name
}

func (p *person) set(name string) {
	p.name = name
}
func main() {
	p := person{name: "Sven"}
	fmt.Println(p)

	changeMe(&p, "Bert")
	fmt.Println(p)

	p.set("Arne")
	fmt.Println(p)
}
