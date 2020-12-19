package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last)
}
func main() {
	p := person{"Sven", "Oak", 32}
	p.speak()
}
