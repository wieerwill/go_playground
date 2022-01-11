package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// atomic counters accessed by multiple goroutines
func main() {
	// unsigned integer to represent our (always-positive) counter
	var ops uint64

	// WaitGroup will help wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// start 50 goroutines that each increment the counter exactly 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// atomically increment the counter, giving it the memory address of ops counter with the & syntax
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// wait until all the goroutines are done
	wg.Wait()

	// it’s safe to access ops now because it's known no other goroutine is writing to it
	fmt.Println("ops:", ops)
}
