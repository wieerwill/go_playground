package main

import "fmt"

// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

// this person struct type has name and age fields
type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct with the given name
// safely return a pointer to local variable as a local variable will survive the scope of the function
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// create a new struct
	fmt.Println(person{"Bob", 20})

	// name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// encapsulate new struct creation in constructor functions (ideomatic)
	fmt.Println(newPerson("Jon"))

	// access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// use dots with struct pointers
	// the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
