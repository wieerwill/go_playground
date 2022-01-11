// built-in synchronization features of goroutines and channels

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// state will be owned by a single goroutine
// this will guarantee that the data is never corrupted with concurrent access
// in order to read or write that state, other goroutines will send messages to the owning goroutine and receive corresponding replies
// these readOp and writeOp structs encapsulate those requests and a way for the owning goroutine to respond
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// count how many operations to perform
	var readOps uint64
	var writeOps uint64

	// reads and writes channels will be used by other goroutines
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// goroutine that owns the state, which is a map as in the previous example but now private to the stateful goroutine
	// this goroutine repeatedly selects on the reads and writes channels, responding to requests as they arrive
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// each read requires constructing a readOp, sending it over the reads channel, and the receiving the result over the provided resp channel
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// start 10 writes as well, using a similar approach
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// let the goroutines work for a second
	time.Sleep(time.Second)

	// capture and report the op counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
