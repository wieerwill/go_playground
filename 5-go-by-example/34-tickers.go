// tickers are for when you want to do something repeatedly at regular intervals

package main

import (
	"fmt"
	"time"
)

func main() {
	// Tickers use a similar mechanism to timers: a channel that is sent values
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickers can be stopped like timers
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
