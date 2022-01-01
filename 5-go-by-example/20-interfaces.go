// Interfaces are named collections of method signatures

package main

import (
	"fmt"
	"math"
)

// basic interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}

// implement interface on rect and circle types
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// implement geometry on rects
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// implementation for circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// call methods that are in the named interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// circle and rect struct types both implement the geometry interface
	// use instances of these structs as arguments to measure
	measure(r)
	measure(c)
}
