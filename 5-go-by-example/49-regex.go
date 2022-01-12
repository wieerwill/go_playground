package main

import (
	"bytes"
	"fmt"
	"regexp"
)

// Go offers built-in support for regular expressions
func main() {
	// test whether a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// compile an optimized Regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	// match test like earlier
	fmt.Println(r.MatchString("peach"))

	// finds the match for the regex
	fmt.Println(r.FindString("peach punch"))

	// finds the first match but returns the start and end indexes for the match instead of the matching text
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// include information about both the whole-pattern matches and the submatches within those matches
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// information about the indexes of matches and submatches
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// apply to all matches in the input, not just the first
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// the All variants are available for the other functions too
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// providing a non-negative integer as the second argument to these functions will limit the number of matches
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// provide []byte arguments and drop String from the function name
	fmt.Println(r.Match([]byte("peach")))

	// MustCompile panics instead of returning an error, which makes it safer to use for global variables
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// replace subsets of strings with other values
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// transform matched text with a given functions
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
