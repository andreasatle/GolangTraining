package main

import "fmt"

func main() {
	// Not the golang way of doing this, but I want to use make...
	persons := make([][]string, 2)
	persons[0] = []string{"James", "Bond", "Shaken, not stirred"}
	persons[1] = []string{"Miss", "Moneypenny", "Hello James"}

	for _, person := range persons {
		for _, word := range person {
			fmt.Printf("%v ", word)
		}
		fmt.Println()
	}
}
