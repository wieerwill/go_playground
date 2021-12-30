# Import package from remote modules
An import path can describe how to obtain the package source code using a revision control system such as Git or Mercurial. The go tool uses this property to automatically fetch packages from remote repositories. 

```bash
cp -r 2-importpackage 3-importremote
# change importpackage to importremote
nano go.mod
nano importpackage/reverse.go
# update modules in hello.go
nano hello.go
>   package main
>   
>   import (
>   	"fmt"
>   
>   	"example/user/importremote/importpackage"
>   
>   	"github.com/google/go-cmp/cmp"
>   )
>   
>   func main() {
>   	fmt.Println(importpackage.ReverseRunes("!oG ,olleH"))
>   	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
>   }
# get all requirements and download modules
go mod tidy
# run code
go run hello.go
```