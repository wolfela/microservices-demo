# microservices-demo

Simple Go implementation of a "Numerical" Microservice using protobuf to compile protoc files into gRPC service code.

# Installation 

This project requires a working version of Protoc/gRPC to be installed. 
Instructions can be found here:

https://grpc.io/docs/quickstart/go.html

As well as the go libraries:

`go get -u google.golang.org/grpc`
`go get -u github.com/golang/protobuf/protoc-gen-go`

# Deployment 

Run `make build` to compile all of the code. 

The server can be then launched with `./server` and the client with `./client`