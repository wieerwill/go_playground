// A goroutine is a lightweight thread of execution
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// normal function call f(s), running synchronously
	f("direct")

	// invoke functions in a goroutine, use go f(s)
	// this new goroutine will execute concurrently with the calling one
	go f("goroutine")

	// start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// the two function calls are running asynchronously in separate goroutines now
	// Wait for them to finish (alternate use WaitGroup)
	time.Sleep(time.Second)
	fmt.Println("done")
}
