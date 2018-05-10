package main

import (
	"log"
	"net"
	"path"
	"path/filepath"
	"runtime"

	pb "./protoc"
	"github.com/tkanos/gonfig"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Configuration struct {
	Port string
}

type service struct {
}

// Implements the first service function defined in protoc file
// Sums the elements of a given array and returns the result in the defined response format
func (s *service) SumArray(ctx context.Context, req *pb.Array) (*pb.Response, error) {
	var sum int32
	sum = 0

	for i := range req.Ints {
		sum += req.Ints[i]
	}

	return &pb.Response{Result: sum}, nil
}

// Implements the second service function defined in the protoc file
// Concats given strings together and returns a single string in the defined response format
// I ran out of ideas for simple functions :)
func (s *service) SumWords(ctx context.Context, req *pb.StringArray) (*pb.String, error) {
	var sum string
	for i := range req.S {
		sum += req.S[i]
	}

	return &pb.String{S: sum}, nil
}

func main() {
	// Port is defined in config file as per 12factor standards
	// Not defined in env since it should be the same for all deployments
	configuration := Configuration{}
	_, dir, _, _ := runtime.Caller(0)
	err := gonfig.GetConf(path.Join(filepath.Dir(dir), "config/defaults.json"), &configuration)
	if err != nil {
		log.Fatalf("Failed getting configuration: %v", err)
	}

	// Launches the server to listen in the defined port
	lis, err := net.Listen("tcp", ":"+configuration.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Service listening on port: %v", configuration)

	s := grpc.NewServer()
	pb.RegisterNumericalServiceServer(s, &service{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
