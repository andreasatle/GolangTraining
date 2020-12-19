// Embedded structs
package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (v vehicle) String() string {
	str := fmt.Sprintf("vehicle{doors: %d, color: %s}", v.doors, v.color)
	return str
}

func (t truck) String() string {
	str := fmt.Sprintf("truck{doors: %d, color: %s, fourWheel: %t}", t.doors, t.color, t.fourWheel)
	return str
}

func (s sedan) String() string {
	str := fmt.Sprintf("sedan{doors: %d, color: %s, luxury: %t}", s.doors, s.color, s.luxury)
	return str
}

func main() {
	v := vehicle{
		doors: 4,
		color: "black",
	}

	f150 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}
	s40 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		luxury: false,
	}
	fmt.Println(f150, s40, v)
}
