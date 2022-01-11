package main

import (
	"fmt"
	"time"
)

// worker, of which will run several concurrent instances
// those workers will receive work on the jobs channel and send the corresponding results on results
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// send workers work and collect their results
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// starts up 3 workers, initially blocked because there are no jobs
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs and then close that channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collect all the results of the work
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
