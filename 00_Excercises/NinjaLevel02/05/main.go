package main

import "fmt"

func main() {
	str := "Hello, world!"
	fmt.Printf("%s\n", str)
	bs := []byte(str)
	fmt.Println(bs)
	fmt.Printf("%U\n", bs)
	fmt.Printf("%#U\n", bs)
}
