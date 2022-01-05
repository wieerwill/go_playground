// synchronize execution across goroutines
package main

import (
	"fmt"
	"time"
)

// the done channel will be used to notify another goroutine that this function’s work is done
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// Send a value to notify "done"
	done <- true
}

func main() {
	// start a worker goroutine, giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)
	// block until receive notification from the worker on the channel
	<-done
}
