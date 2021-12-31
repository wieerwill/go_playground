package main

import "fmt"

// for is Go’s only looping construct

func main() {

	// basic type with a single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic initial/condition/after
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// without a condition will loop repeatedly until break or return
	for {
		fmt.Println("loop")
		break
	}

	// continue to the next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
