package main

import "fmt"

func main() {
	vs := []int{1, 3, 5, 7, 9}
	fmt.Println(foo(vs...))
	fmt.Println(bar(vs))
}

func foo(vs ...int) int {
	sum := 0
	for _, v := range vs {
		sum += v
	}
	return sum
}

func bar(vs []int) int {
	sum := 0
	for _, v := range vs {
		sum += v
	}
	return sum
}
