package main

import (
	"fmt"

	"github.com/andreasatle/Udemy/GolangTraining/00_Excercises/NinjaLevel13/01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
