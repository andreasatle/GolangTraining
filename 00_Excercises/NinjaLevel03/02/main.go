package main

import "fmt"

func main() {
	for c := 65; c <= 90; c++ {
		fmt.Println(c)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", c)
		}
	}
}
