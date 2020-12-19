package main

import "fmt"

const (
	a         = 3
	b float64 = 2
)

func main() {
	fmt.Printf("Integer: %d\t%T\n", a, a)
	fmt.Printf("Float: %f\t%T\n", b, b)
}
