package main

import (
	"fmt"
)

type Coord struct{ x, y int }

func (c Coord) String() string {
	return fmt.Sprintf("Coord{%d,%d}", c.x, c.y)
}

const maxNum = 10

func main() {
	for v := range coordGenerator() {
		fmt.Println(v)
	}
}

func coordGenerator() <-chan Coord {
	cxy := make(chan Coord)

	pointGenerator := func() <-chan int {
		cp := make(chan int)
		go func() {
			for p := 0; p < maxNum; p++ {
				cp <- p
			}
			close(cp)
		}()
		return cp
	}

	go func() {
		for x := 0; x < maxNum; x++ {
			for y := range pointGenerator() {
				cxy <- Coord{x, y}
			}
		}
		close(cxy)
	}()
	return cxy
}
