package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var mu sync.Mutex
var wg sync.WaitGroup

func main() {
	const n = 100
	var counter int64
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
}
