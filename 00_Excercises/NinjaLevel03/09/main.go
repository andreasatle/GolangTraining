package main

import "fmt"

func main() {
	str := "favSport"
	switch str {
	case "FavSport":
		fmt.Println("Wrong case")
	case "favSport":
		fmt.Println("Correct case")
	}
}
