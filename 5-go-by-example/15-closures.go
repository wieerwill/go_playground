package main

import "fmt"

// Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

// intSeq returns another function, which is defined anonymously in the body of intSeq
// the returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// call intSeq, assigning the result (a function) to nextInt
	nextInt := intSeq()

	// see the effect of the closure
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// confirm that the state is unique to that particular function
	newInts := intSeq()
	fmt.Println(newInts())
}
