// get code working using
// 1) func literal, aka, anonymous self-executing func
// 2) a buffered channel
package main

import (
	"fmt"
)

func main() {
	// 1) Solution with go-routine for sending
	c := make(chan int)
	go func() { c <- 42 }()
	fmt.Println(<-c)

	// 2) Solution with a buffered channel
	d := make(chan int, 1)
	d <- 43
	fmt.Println(<-d)
}
