package main

import (
	"fmt"
	"time"
)

// controlling resource utilization and maintaining quality of service
func main() {

	// basic rate limiting
	// want to limit handling of incoming requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// this limiter channel will receive a value every 200 milliseconds
	// this is the regulator in the rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// by blocking on a receive from the limiter channel before serving each request,
	// limit to 1 request every 200 milliseconds
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// allow short bursts of requests in rate limiting scheme while preserving the overall rate limit
	// accomplish this by buffering the limiter channel
	burstyLimiter := make(chan time.Time, 3)

	// fill up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// every 200 milliseconds try to add a new value to burstyLimiter, up to its limit of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// simulate 5 more incoming requests
	// the first 3 of these will benefit from the burst capability of burstyLimiter
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
