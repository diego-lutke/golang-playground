# golang-playground

Repo for the author to play with golang for the first time,
everything in here is based on the material available in
[https://golang-for-python-programmers.readthedocs.io/en/latest/index.html](https://golang-for-python-programmers.readthedocs.io/en/latest/index.html)

## go installation
`sudo apt install gcc libc6-dev mercurial`

[https://go.dev/doc/install](https://go.dev/doc/install)

## version
`go version`

## go hello world
[https://go.dev/doc/tutorial/getting-started](https://go.dev/doc/tutorial/getting-started)

## go packages
[https://pkg.go.dev/](https://pkg.go.dev/)

## go docs
`go install golang.org/x/tools/cmd/godoc@latest`

## environment variables
`export GOPATH=~/go`
`export PATH=$PATH:$GOPATH/bin`

## workspace preparation

Create a directory for your Go project inside the src directory:
`mkdir -p ~/go/src/github.com/yourusername/yourproject`

In your repo root directory, make a new file called `go.mod` (this file manages
the project dependencies). In the repo folter, run:
`go mod init example.com/myproject`

## indentation
The recommended indentation is to use tabs instead of spaces.

## formatter
`gofmt -w <filename>.go  # -w option updates the file(s) in place`

## running a file
e.g.: `go run 01_hello.go`
