package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter atomic.Uint64

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < 100_000; j++ {
				counter.Add(1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Load())
}

/* func main() {
	var wg sync.WaitGroup
	var counter int64

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 10_000)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
} */
