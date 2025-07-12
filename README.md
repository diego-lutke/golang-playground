# golang-playground

Repo for the author to play with golang for the first time.

## how to get starting

[official installation reference](https://go.dev/doc/install)

- go tarball
  `curl -L -O https://go.dev/dl/go1.24.5.linux-amd64.tar.gz`

- remove previous installations
  `rm -rf /usr/local/go`

- extract the tarball into `usr/local`
  `sudo tar -C /usr/local -xzf go1.24.5.linux-amd64.tar.gz`

- requirements
  `sudo apt install gcc libc6-dev mercurial`

- docs
  `go install golang.org/x/tools/cmd/godoc@latest`

- version
  `go version`

- environment variables
  `export PATH=$PATH:/usr/local/go/bin`
  `export GOPATH=~/go`
  `export PATH=$PATH:$GOPATH/bin`

## workspace preparation

- In your repo root directory, make a new file called `go.mod` (this file manages
  the project dependencies). In the repo root folter, run:
  `go mod init github.com/diego-lutke/golang-playground`

## basic usage

- running a file
  `go run <filename>.go`

- running a project
  `go run .`

- formatting a file
  `gofmt -w <filename>.go  # -w option updates the file(s) in place`

- installing a package
  `go get github.com/onsi/gomega/...`

- cleaning and updating go.mod/go.sum to reflect required dependencies
  `go mod tidy`

## cool resources

- [hello world](https://go.dev/doc/tutorial/getting-started)
- [tour](https://go.dev/tour/list)
- [packages](https://pkg.go.dev/)
- [playground](https://go.dev/play/)
- [project layouts - official](https://go.dev/doc/modules/layout)
- [project layouts - unofficial](https://github.com/golang-standards/project-layout)

## general notes

- The recommended indentation is to use tabs instead of spaces.
- Go is a strongly typed static programming language.
