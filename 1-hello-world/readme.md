# Hello World
Steps to reproduce
```bash
mkdir 1-hello-world
cd 1-hello-world
go mod init example/user/hello
cat go.mod
# first statement in a Go source file must be package name. Executable commands must always use package main. 

# to run
go run hello.go

# to build and run binary
go build
./hello

# to install binary to system
go install example/user/hello
# or
go install
```