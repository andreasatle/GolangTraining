package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("I'm false")
	case true:
		fmt.Println("I'm true")
	}
}
