package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

// Generator function
func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		q <- 0
	}()

	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Value:", v)
		case v := <-q:
			fmt.Println("About to exit", v)
			return
		}
	}
}
