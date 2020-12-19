package main

import (
	"fmt"
	"runtime"
	"sync"
)

// var mu sync.Mutex
var wg sync.WaitGroup

func main() {
	const n = 100
	counter := 0
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			// mu.Lock()
			i := counter
			i++
			runtime.Gosched()
			counter = i
			fmt.Println("Counter:", counter)
			// mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
