package main

import "fmt"

func main() {
	add3 := addN(3)
	add4 := addN(4)
	x := 4
	fmt.Println(x, add3(x))
	fmt.Println(x, add4(x))
}

func addN(n int) func(int) int {
	return func(x int) int { return x + n }
}
