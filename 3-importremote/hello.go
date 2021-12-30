package main

import (
	"fmt"

	"example/user/importremote/importpackage"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(importpackage.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}