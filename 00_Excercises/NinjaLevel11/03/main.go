package main

import "fmt"

type customErr struct {
	msg string
}

func (ce customErr) Error() string {
	return ce.msg
}

func foo(err error) {
	fmt.Println(err.Error())
}

func main() {
	foo(customErr{"Fake error"})
}
