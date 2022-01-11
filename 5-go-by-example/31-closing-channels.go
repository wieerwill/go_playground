// closing a channel indicates that no more values will be sent on it
package main

import "fmt"

func main() {
	// jobs channel to communicate work to be done
	jobs := make(chan int, 5)
	done := make(chan bool)

	// repeatedly receives from jobs
	// more value will be false if jobs has been closed and all values in the channel have already been received
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// sends 3 jobs to the worker over the jobs channel, then closes it
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// await the worker using the synchronization approach
	<-done
}
