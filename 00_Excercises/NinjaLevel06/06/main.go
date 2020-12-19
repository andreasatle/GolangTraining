package main

import "fmt"

func main() {
	func(n int) {
		fmt.Println("Meaning of life:", n)
	}(42)
}
