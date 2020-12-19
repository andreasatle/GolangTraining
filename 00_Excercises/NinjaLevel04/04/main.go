package main

import "fmt"

func main() {
	vs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	vs = append(vs, 52)
	fmt.Println(vs)
	vs = append(vs, 53, 54, 55)
	fmt.Println(vs)
	y := []int{56, 57, 58, 59, 60}
	vs = append(vs, y...)
	fmt.Println(vs)
}
