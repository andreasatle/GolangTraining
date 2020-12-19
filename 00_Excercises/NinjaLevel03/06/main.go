package main

import "fmt"

func main() {
	if 2 < 3 {
		fmt.Println("2 < 3 is true")
	}

	if !(2 < 3) {
		fmt.Println("2 < 3 is not true")
	}
}
