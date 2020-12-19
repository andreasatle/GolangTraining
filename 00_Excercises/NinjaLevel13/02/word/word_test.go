package word

import (
	"fmt"
	"testing"

	"github.com/andreasatle/Udemy/GolangTraining/00_Excercises/NinjaLevel13/02/quote"
)

func TestCountSunAlso(t *testing.T) {
	count := Count(quote.SunAlso)
	if count != 1349 {
		fmt.Errorf("Got: %d Expected: %d", count, 1349)
	}
}

func ExampleCount() {
	fmt.Println(Count("Gosh, I'm good!"))
	// output:
	// 3
}

func ExampleUseCount() {
	fmt.Println(UseCount("How do you do"))
	// output:
	// map[How:1 do:2 you:1]
}
