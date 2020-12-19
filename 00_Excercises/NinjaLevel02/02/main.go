package main

import "fmt"

func main() {
	a, b, c := 42, 65, 87
	fmt.Println("a:", a, "b:", b, "c:", c)
	fmt.Println("a == a", a == a)
	fmt.Println("a == b", a == b)

	fmt.Println("a <= a", a <= a)
	fmt.Println("a <= b", a <= b)
	fmt.Println("b <= a", b <= a)

	fmt.Println("a >= a", a >= a)
	fmt.Println("a >= b", a >= b)
	fmt.Println("b >= a", b >= a)

	fmt.Println("a != a", a != a)
	fmt.Println("a != b", a != b)
	fmt.Println("b != a", b != a)

	fmt.Println("a < a", a < a)
	fmt.Println("a < b", a < b)
	fmt.Println("b < a", b < a)

	fmt.Println("a > a", a > a)
	fmt.Println("a > b", a > b)
	fmt.Println("b > a", b > a)
}
