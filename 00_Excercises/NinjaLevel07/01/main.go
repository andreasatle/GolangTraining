package main

import "fmt"

func main() {
	x := 1
	fmt.Printf("%v\t%T\n", x, x)
	fmt.Printf("%v\t%T\n", &x, &x)
}
