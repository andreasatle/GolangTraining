package main

import "fmt"

func main() {
	vs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(vs[:3], vs[6:]...)
	fmt.Println(y)
}
