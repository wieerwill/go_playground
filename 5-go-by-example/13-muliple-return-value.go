package main

import "fmt"

// (int, int) in this function signature shows that the function returns 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {

	// use the 2 different return values
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// use the blank identifier _
	_, c := vals()
	fmt.Println(c)
}
