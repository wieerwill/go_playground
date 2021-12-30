# Testing
Go has a lightweight test framework composed of the go test command and the testing package.

Write a test by creating a file with a name ending in `_test.go` that contains functions named TestXXX with signature `func (t *testing.T)`. The test framework runs each such function; if the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed. 

```bash
cp -r 3-importremote 4-testing
# change importremote to testing
nano go.mod
nano hello.go
# create test for reverse
cd importpackage
nano reverse_test.go
>   package morestrings
>   
>   import "testing"
>   
>   func TestReverseRunes(t *testing.T) {
>   	cases := []struct {
>   		in, want string
>   	}{
>   		{"Hello, world", "dlrow ,olleH"},
>   		{"Hello, 世界", "界世 ,olleH"},
>   		{"", ""},
>   	}
>   	for _, c := range cases {
>   		got := ReverseRunes(c.in)
>   		if got != c.want {
>   			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
>   		}
>   	}
>   }
go build
# run test code
go test
```