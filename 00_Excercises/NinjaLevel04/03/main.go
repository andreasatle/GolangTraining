package main

import "fmt"

func main() {
	vs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	vss := [][]int{vs[:5], vs[5:], vs[2:7], vs[1:6]}
	for i, vs := range vss {
		fmt.Println(i, vs)
	}

}
