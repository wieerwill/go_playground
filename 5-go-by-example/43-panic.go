package main

import "os"

// a panic typically means something went unexpectedly wrong
func main() {

	// this is the only code here designed to panic
	panic("a problem")

	// a common use of panic is to abort if a function returns an error value that one doesn’t know how to handle
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
