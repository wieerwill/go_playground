package main

import "fmt"

// Go supports pointers, allowing you to pass references to values and records within your program

// zeroval has an int parameter, so arguments will be passed to it by value
// zeroval will get a copy of ival distinct from the one in the calling function
func zeroval(ival int) {
	ival = 0
}

// zeroptr has an *int parameter, meaning that it takes an int pointer
// the *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address
// assigning a value to a dereferenced pointer changes the value at the referenced address
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i syntax gives the memory address
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed
	fmt.Println("pointer:", &i)
}
