FROM golang:1.14

WORKDIR $GOPATH/src/github.com/ayusrina/todo2
COPY . .

