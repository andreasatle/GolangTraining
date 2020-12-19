package main

import "fmt"

func main() {
	fmt.Printf("true && true = %v\n", true && true)
	fmt.Printf("true && false = %v\n", true && false)
	fmt.Printf("false && true = %v\n", false && true)
	fmt.Printf("false && false = %v\n", false && false)

	fmt.Printf("true || true = %v\n", true || true)
	fmt.Printf("true || false = %v\n", true || false)
	fmt.Printf("false || true = %v\n", false || true)
	fmt.Printf("false || false = %v\n", false || false)

	fmt.Printf("!true = %v\n", !true)
	fmt.Printf("!false = %v\n", !false)
}
