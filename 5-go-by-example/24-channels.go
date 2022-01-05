// Channels are the pipes that connect concurrent goroutines
package main

import "fmt"

func main() {
	// create new channel with make(chan val-type)
	messages := make(chan string)

	// send a value into a channel, here "ping"
	go func() { messages <- "ping" }()

	// receives a value from the channel
	msg := <-messages
	fmt.Println(msg)
}
