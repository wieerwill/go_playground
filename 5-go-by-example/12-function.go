package main

import "fmt"

// function that takes two ints and returns their sum as an int
// Go requires explicit returns
func plus(a int, b int) int {
	return a + b
}

// multiple consecutive parameters of the same type,
// may omit the type name for the like-typed parameters up to the final parameter that declares the type
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	// Call a function
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
