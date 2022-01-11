package main

import (
	"fmt"
	"time"
)

// built-in timer and ticker features
func main() {

	// timers represent a single event in the future
	// tell the timer how long to wait and it provides a channel that will be notified at that time
	timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// to wait, you could have used time.Sleep
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// one can cancel the timer before it fires
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped
	time.Sleep(2 * time.Second)
}
