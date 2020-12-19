package main

import "fmt"

func main() {
	c := generator()
	receiver(c)
}

func generator() <-chan int {
	c := make(chan int)
	go func() {
		for fn, fn1 := 0, 1; fn1 <= 10000; fn, fn1 = fn1, fn+fn1 {
			c <- fn1
		}
		close(c)
	}()
	return c
}

func receiver(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
