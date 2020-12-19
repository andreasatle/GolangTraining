package main

import "fmt"

func main() {
	start, stop, next := createIterator(2, 14, 4)
	for i := start(); stop(i); next(&i) {
		fmt.Println(i)
	}
}

func createIterator(start, stop, increment int) (func() int, func(int) bool, func(*int)) {

	local := start

	startFunc := func() int {
		return start
	}

	stopFunc := func(i int) bool {
		return i <= stop
	}

	nextFunc := func(i *int) {
		local += increment
		*i = local
	}

	return startFunc, stopFunc, nextFunc
}
