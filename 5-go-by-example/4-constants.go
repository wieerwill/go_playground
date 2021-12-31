package main

import (
	"fmt"
	"math"
)

// const declares a constant value.
const s string = "constant"

func main() {
	fmt.Println(s)

	// can appear anywhere a var statement can
	const n = 500000000

	// expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// numeric constants have no type until it’s given one
	fmt.Println(int64(d))

	// numbers can be given a type by using it in a context that requires one
	fmt.Println(math.Sin(n))
}
