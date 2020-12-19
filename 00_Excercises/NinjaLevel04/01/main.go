package main

import "fmt"

func main() {
	vs := [5]int{1, 2, 3, 4, 5}
	vs[3] = 7
	for i, v := range vs {
		fmt.Println(i, v)
	}
	fmt.Printf("Type of array: %T\n", vs)
}
