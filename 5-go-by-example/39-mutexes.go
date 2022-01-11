package main

import (
	"fmt"
	"sync"
)

// holds a map of counters
// add a Mutex to synchronize access
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// lock the mutex before accessing counters
	c.mu.Lock()
	// unlock it at the end of the function using defer
	defer c.mu.Unlock()
	c.counters[name]++
}

// mutex to safely access data across multiple goroutines
func main() {
	c := Container{
		// the zero value of a mutex is usable as-is
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// increments a named counter in a loop
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// run several goroutines concurrently
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// wait a for the goroutines to finish
	wg.Wait()
	fmt.Println(c.counters)
}
