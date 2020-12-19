package main

import (
	"fmt"

	"github.com/andreasatle/Udemy/GolangTraining/00_Excercises/NinjaLevel13/02/quote"
	"github.com/andreasatle/Udemy/GolangTraining/00_Excercises/NinjaLevel13/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
