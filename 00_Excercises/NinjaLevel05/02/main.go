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
		lastName:  "Ekstrom",
		favIce:    "Vanilla",
	}

	p2 := person{
		firstName: "Bo",
		lastName:  "Ek",
		favIce:    "Strawberry",
	}

	ps := []person{p1, p2}
	m := map[string]person{}

	for _, p := range ps {
		m[p.lastName] = p
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
