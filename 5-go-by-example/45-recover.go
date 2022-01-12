package main

import "fmt"

// this function panics
func mayPanic() {
	panic("a problem")
}

// it is possible to recover from a panic, by using the recover built-in function
func main() {

	// recover must be called within a deferred function
	// when the enclosing function panics, the defer will activate and a recover call within it will catch the panic
	defer func() {
		if r := recover(); r != nil {
			// the return value of recover is the error raised in the call to panic
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// this code will not run, because mayPanic panics
	// the execution of main stops at the point of the panic and resumes in the deferred closure
	fmt.Println("After mayPanic()")
}
