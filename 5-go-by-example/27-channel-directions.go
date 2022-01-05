// specify if a channel is meant to only send or receive values
package main

import "fmt"

// ping function only accepts a channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong function accepts one channel for receives (pings) and a second for sends (pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
