package main

import (
	"fmt"
	"sort"
)

// in order to sort by a custom function in Go, you need a corresponding type
type byLength []string

// implement sort.Interface - Len, Less, and Swap - on the type
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// implement the custom sort by converting the original fruits slice to byLength and then use sort
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
