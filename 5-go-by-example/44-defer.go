package main

import (
	"fmt"
	"os"
)

// Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup
func main() {
	// suppose you wanted to create a file, write to it, and then close when done
	// immediately after getting a file object with createFile, defer the closing of that file with closeFile
	// this will be executed at the end of the enclosing function (main), after writeFile has finished
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

// it’s important to check for errors when closing a file, even in a deferred function
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
