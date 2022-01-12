package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

// common string formatting tasks
func main() {
	// prints an instance of point struct
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// the %+v variant will include the struct’s field names
	fmt.Printf("struct2: %+v\n", p)

	// the %#v variant prints a Go syntax representation of the value
	fmt.Printf("struct3: %#v\n", p)

	// the %T variant prints the type of a value
	fmt.Printf("type: %T\n", p)

	// formatting booleans straight-forward
	fmt.Printf("bool: %t\n", true)

	// use %d for standard, base-10 formatting
	fmt.Printf("int: %d\n", 123)

	// prints a binary representation
	fmt.Printf("bin: %b\n", 14)

	// prints the character corresponding to the given integer
	fmt.Printf("char: %c\n", 33)

	// %x provides hex encoding
	fmt.Printf("hex: %x\n", 456)

	// decimal formatting use %f
	fmt.Printf("float1: %f\n", 78.9)

	// %e and %E format the float in (different versions of) scientific notation
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// basic string printing use %s
	fmt.Printf("str1: %s\n", "\"string\"")

	// to double-quote strings use %q
	fmt.Printf("str2: %q\n", "\"string\"")

	// %x renders the string in base-16, with two output characters per byte of input
	fmt.Printf("str3: %x\n", "hex this")

	// use %p to print a representation of a pointer
	fmt.Printf("pointer: %p\n", &p)

	// to specify the width of an integer, use a number after the % in the verb
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// to specify the width of printed floats
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// to left-justify, use the - flag
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// specify width when formatting strings
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// and left-justify strings with - too
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats and returns a string without printing it anywhere
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// format+print to io.Writers other than os.Stdout using Fprintf
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
