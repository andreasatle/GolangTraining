// Anonomous struct
package main

import "fmt"

func main() {
	p := struct {
		firstName string
		lastName  string
	}{
		firstName: "Bo",
		lastName:  "Ek",
	}

	fmt.Println(p)
}
