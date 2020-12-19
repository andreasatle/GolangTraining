package main

import "fmt"

func main() {
	vs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Slice:", vs)
	fmt.Println("Sum Even", actionWithFilter(vs, sum, even))
	fmt.Println("Sum Odd", actionWithFilter(vs, sum, odd))
	fmt.Println("Prod Even", actionWithFilter(vs, prod, even))
	fmt.Println("Prod Odd", actionWithFilter(vs, prod, odd))
}

func sum(vs []int) int {
	res := 0
	for _, v := range vs {
		res += v
	}
	return res
}

func prod(vs []int) int {
	res := 1
	for _, v := range vs {
		res *= v
	}
	return res
}

func even(vs []int) []int {
	res := []int{}
	for _, v := range vs {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

func odd(vs []int) []int {
	res := []int{}
	for _, v := range vs {
		if v%2 != 0 {
			res = append(res, v)
		}
	}
	return res
}

func actionWithFilter(vs []int, action func([]int) int, filter func([]int) []int) int {
	return action(filter(vs))
}
