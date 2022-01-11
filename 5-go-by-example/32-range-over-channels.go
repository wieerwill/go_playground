package main

import "fmt"

func main() {
	// iterate over values received from a channel

	// here iterate over 2 values in the queue channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// range iterates over each element as it’s received from queue
	for elem := range queue {
		fmt.Println(elem)
	}
	// also shows that it’s possible to close a non-empty channel but still have the remaining values be received
}
