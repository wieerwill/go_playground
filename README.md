# GO Playground
This repo ist to learn and play with GO.

All Code is written and tested on a Ubuntu 20.4 Linux. 

Learn more about [GO](https://go.dev/)

## Install from Release
```bash
# Download release, change the number to latest release
curl -OL https://golang.org/dl/go1.17.5.linux-amd64.tar.gz
# verify integrity of download, compare with website's shown hash
sha256sum go1.16.7.linux-amd64.tar.gz
# extract tarball
sudo tar -C /usr/local -xvf go1.17.5.linux-amd64.tar.gz
# set up paths
sudo nano ~/.profile
>   export PATH=$PATH:/usr/local/go/bin
source ~/.profile
# confirm installation
which go
go version
```

## Install from Source
```bash
# install requirements
sudo apt install git gcc 
# get source repository
git clone https://go.googlesource.com/go goroot
cd goroot
#  <tag> is the version string of the release
git checkout <tag>
cd src
# install
./all.bash
# test installation
which go
go version
```

The compilation environment can be customized by environment variables (but not required). Heres a short list
- `$GOROOT`: root of the GO tree, often `$HOME/go1.X`
- `$GOPATH`: directory where Go projects outside the Go distribution are typically checked out. Executables outside the Go distribution are installed in $GOPATH/bin (or $GOBIN, if set)
- `$GOBIN`: directory where executables outside the Go distribution are installed using the go command

To change any variable, edit your `~/.profile` file. Scroll all the way down to the end of the file and add the following with your specifications (sample here shows default values)
```bash
GOROOT=$HOME/go1.X
GOPATH=$HOME/go
GOBIN=$GOPATH/bin
```
Feel free to change it to your needs. Finally make the system aware of the new profile and run `source ~/.profile`


### Update your GO Version
GO releases can be fetched with the installed repository above. To never forget any release you can join the [golang-announce mailing list](https://groups.google.com/group/golang-announce).

To update to the latest release, simply run
```bash
cd go/src
git fetch
# change <tag> to latest release tag e.g. go1.9
git checkout <tag>
./all.bash
```

## First programm
```go
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```
Build and run this program:
```bash
go run hello.go
```

## Code to binary
The `go run` command is used as a shortcut for compiling and running a program that requires frequent changes. 

When you’ve finished your code and want to run it without compiling each time, use `go build` to turn your code into an executable binary. Building code into an executable binary consolidates the application into a single file with all the (support) code necessary to execute the binary. 

Once built, run `go install` to place the program on an executable file path to run it from anywhere on the system.

To build and install executables you need an `go.mod` file with your executable definitions. Follow the [helloworld program](1-hello-world/hello.go) to see how.

```bash
# make sure to run from same directory as .go file
go build
# execute binary
./hello
# install to system
go install
# execute global binary
hello
```

## Install additional tools
Several Go tools are kept in the `golang.org/x/tools` repository. To install one of the tools, e.g. `gopls`:
```bash
go install golang.org/x/tools/gopls@latest
```