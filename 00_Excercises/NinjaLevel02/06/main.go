package main

import "fmt"

const (
	thisYear = 2020 + iota
	nextYear
	inTwoYears
)

func main() {
	fmt.Println(thisYear, nextYear, inTwoYears)
}
