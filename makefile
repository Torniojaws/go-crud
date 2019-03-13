SHELL := /bin/bash

install:
    sudo apt install golang-go go-dep
    export GOPATH=~/go
    go get -u golang.org/x/lint/golint
    echo "Add export GOPATH to your .bash_profile"
    mkdir -p $GOPATH/src/github.com/Torniojaws/go-crud
    echo "Make sure the repository was cloned to $GOPATH/src/github.com/Torniojaws/go-crud"
    cd $GOPATH/src/github.com/Torniojaws/go-crud
    dep ensure
    
run:
    go run main.go
    
test:
    go test
