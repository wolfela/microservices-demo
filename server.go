package main

import (
	"log"
	"net"

	pb "./protoc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type service struct {
}

func (s *service) SumArray(ctx context.Context, req *pb.Array) (*pb.Response, error) {
	var sum int32
	sum = 0

	for i := range req.Ints {
		sum += req.Ints[i]
	}

	return &pb.Response{Result: sum}, nil
}

func (s *service) SumWords(ctx context.Context, req *pb.StringArray) (*pb.String, error) {
	var sum string
	for i := range req.S {
		sum += req.S[i]
	}

	return &pb.String{S: sum}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNumericalServiceServer(s, &service{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
