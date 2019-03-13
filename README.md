# go-crud
Minimal CRUD in Go with Gin framework

## Install

1. Install Go
1. Add the GOPATH to `.bash_profile`; `export GOPATH=~/go`
1. Install the package manager: `sudo apt install go-dep`
1. Add and enter the project path:
    ```
    mkdir -p $GOPATH/src/github.com/go-crud
    cd $GOPATH/src/github.com/go-crud
    ```
1. Init the project `dep init`(needs a dummy Go file in the root)
1. Install the project dependencies: `dep ensure`
    1. Or add new ones: `dep ensure -add github.com/some/project`
    1. Update existing ones: `dep ensure -update`
1. Start the project with `go run main.go`
1. Browse to http://localhost:8080/v1/hello
