package main

import "fmt"

func main() {
	defer foo(1)
	defer foo(2)
	foo(3)
}

func foo(x int) {
	fmt.Println(x)
}
