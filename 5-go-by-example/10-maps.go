package main

import "fmt"

// Maps are Go’s built-in associative data type
func main() {

	// create an empty map, use the builtin: make(map[key-type]val-type).
	m := make(map[string]int)

	// set key/value pairs using typical name[key]=val syntax
	m["k1"] = 7
	m["k2"] = 13

	// printing a map will show all of its key/value pairs
	fmt.Println("map:", m)

	// get a value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// len returns the number of key/value pairs
	fmt.Println("len:", len(m))

	// delete removes key/value pairs
	delete(m, "k2")
	fmt.Println("map:", m)

	// optional second return value when getting a value from a map indicates if the key was present in the map
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
