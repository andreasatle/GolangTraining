package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	favIce    string
}

func main() {
	p1 := person{
		firstName: "Anders",
		lastName:  "Ek",
		favIce:    "Vanilla",
	}

	p2 := person{
		firstName: "Bo",
		lastName:  "Ek",
		favIce:    "Strawberry",
	}

	ps := []person{p1, p2}

	for _, p := range ps {
		fmt.Println(p.firstName, p.lastName, p.favIce)
	}
}
