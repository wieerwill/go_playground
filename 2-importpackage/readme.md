# Import package
write a `importpackage` package and use it from the hello program. 

Because the `ReverseRunes` function begins with an upper-case letter, it is exported and can be used in other packages that import the importpackage package. 

Test that the package compiles with go build:
```bash
# init project
mkdir 2-importpackage
cd 2-importpackage
go mod init example/user/importpackage
# create files
touch hello.go
mkdir importpackage
cd importpackage
nano reverse.go
>   // Package importpackage implements additional functions to manipulate UTF-8
>   // encoded strings, beyond what is provided in the standard "strings" package.
package importpackage
>   
>   // ReverseRunes returns its argument string reversed rune-wise left to right.
>   func ReverseRunes(s string) string {
>   	r := []rune(s)
>   	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
>   		r[i], r[j] = r[j], r[i]
>   	}
>   	return string(r)
>   }
go build
```

This won't produce an output file. Instead it saves the compiled package in the local build cache.

After confirming that the importpackage package builds, use it from the `hello` program.
```bash
cd ..
nano hello.go
>   package main
>   
>   import (
>   	"fmt"
>   
>   	"example/user/importpackage/importpackage"
>   )
>   
>   func main() {
>   	fmt.Println(importpackage.ReverseRunes("!oG ,olleH"))
>   }
go run hello.go
```
