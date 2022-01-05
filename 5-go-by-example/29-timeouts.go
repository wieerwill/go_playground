// bound execution time
package main

import (
	"fmt"
	"time"
)

func main() {
	// returns its result on a channel c1 after 2s
	// channel is buffered, so the send in the goroutine is nonblocking
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// select implementing a timeout
	// res awaits the result and <-time
	// After awaits a value to be sent after the timeout
	// select proceeds with the first receive that’s ready
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// allow a longer timeout of 3s, then the receive from c2 will succeed and print the result
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
