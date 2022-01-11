// wait for multiple goroutines to finish

package main

import (
	"fmt"
	"sync"
	"time"
)

// function to run in every goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// WaitGroup is used to wait for all the goroutines launched here to finish
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// Avoid re-use of the same i value in each goroutine closure
		i := i
		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	// Block until the WaitGroup counter goes back to 0
	wg.Wait()

}
