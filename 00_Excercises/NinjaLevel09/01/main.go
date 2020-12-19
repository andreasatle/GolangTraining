package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Add(4)
	go foo("FOO1")
	go foo("FOO2")
	go foo("FOO3")
	go foo("FOO4")

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}

func foo(str string) {
	for i := 0; i < 20; i++ {
		fmt.Println(str, i)
		runtime.Gosched()
	}
	wg.Done()
}
