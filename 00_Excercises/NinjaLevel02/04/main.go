package main

import "fmt"

const (
	a = 4711
	b = a << 1
)

func main() {
	fmt.Printf("%d, %b, %#x\n", a, a, a)
	fmt.Printf("%d, %b, %#x\n", b, b, b)
}
