# golang-playground

Repo for the author to play with golang for the first time.

## how to get starting

[official installation reference](https://go.dev/doc/install)

- requirements
  `sudo apt install gcc libc6-dev mercurial`

- docs
  `go install golang.org/x/tools/cmd/godoc@latest`

- version
  `go version`

- environment variables
  `export GOPATH=~/go`
  `export PATH=$PATH:$GOPATH/bin`

## cool resources

- [hello world](https://go.dev/doc/tutorial/getting-started)
- [tour](https://go.dev/tour/list)
- [packages](https://pkg.go.dev/)
- [playground](https://go.dev/play/)
- [project layouts](https://github.com/golang-standards/project-layout)

## workspace template

- Create a directory for your Go project inside the src directory:
  `mkdir -p ~/go/src/github.com/yourusername/yourproject`

- In your repo root directory, make a new file called `go.mod` (this file manages
  the project dependencies). In the repo folter, run:
  `go mod init example.com/myproject`

## basic usage

- running a file
  `go run <filename>.go`

- formatting a file
  `gofmt -w <filename>.go  # -w option updates the file(s) in place`

## general notes

- The recommended indentation is to use tabs instead of spaces.
- Go is a strongly typed static programming language.
