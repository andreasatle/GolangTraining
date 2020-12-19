package main

import "fmt"

func main() {
	funcWithVariadicArg(1, 2, 4)
	funcWithVariadicArg([]int{1, 2, 4}...)
	q := make([]int, 10, 20)
	fmt.Println(q)

	m := map[string]string{}
	m[""] = ""
	delete(m, "foo")
	if v, ok := m[""]; ok {
		fmt.Println(v, "Came here?!")
	}
}

func funcWithVariadicArg(vs ...int) {
	fmt.Printf("%v\t%T\n", vs, vs)
}
