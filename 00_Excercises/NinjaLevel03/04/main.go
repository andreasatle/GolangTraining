package main

import "fmt"

func main() {
	year := 1972
	for {
		fmt.Println("I was living in year", year)
		year++
		if year > 2020 {
			break
		}
	}
}
