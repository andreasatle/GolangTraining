package main

import (
	"fmt"

	"github.com/andreasatle/Udemy/GolangTraining/00_Excercises/NinjaLevel12/01/dog"
)

func main() {
	humanYears := 13.0
	dogYears := dog.Years(humanYears)
	fmt.Printf("%v human years equals %v dog years.\n", humanYears, dogYears)
}
