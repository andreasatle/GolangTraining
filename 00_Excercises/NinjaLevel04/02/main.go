package main

import "fmt"

func main() {
	vs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	vs[3] = 7
	for i, v := range vs {
		fmt.Println(i, v)
	}
	fmt.Printf("Type of array: %T\n", vs)

}
