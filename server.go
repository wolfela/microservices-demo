package main

import (
	"log"
	"net"

	// Import the generated protobuf code
	pb "./protoc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// Return matching the `Response` message we created in our
	// protobuf definition.
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

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterNumericalServiceServer(s, &service{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
