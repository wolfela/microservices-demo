build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)roomone/microservices-demo/ \
	  protoc/numerical.proto
