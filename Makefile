build:
	go build server.go

	go build client.go

	protoc -I. --go_out=plugins=grpc:$(GOPATH)roomone/microservices-demo/ \
	  protoc/numerical.proto
