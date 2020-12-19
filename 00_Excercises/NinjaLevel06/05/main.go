package main

import (
	"fmt"
	"math"
)

type shape interface 
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("Area of shape:", s.area())
}
func main() {
	sq := square{side: 2.0}
	ci := circle{radius: 3.0}
	info(sq)
	info(ci)
}
