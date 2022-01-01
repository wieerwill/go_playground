package main

import "fmt"

// Variadic functions can be called with any number of trailing arguments

// take an arbitrary number of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Variadic functions can be called in the usual way
	sum(1, 2)
	sum(1, 2, 3)

	// apply multiple args to a variadic function
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
