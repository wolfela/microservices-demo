package main

import (
	"log"
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

func main() {
	configuration := Configuration{}
	_, dir, _, _ := runtime.Caller(0)
	err := gonfig.GetConf(path.Join(filepath.Dir(dir), "config/defaults.json"), &configuration)
	if err != nil {
		log.Fatalf("Failed getting configuration: %v", err)
	}
	address := "localhost:" + configuration.Port
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewNumericalServiceClient(conn)

	primes := [6]int32{2, 3, 5, 7, 11, 13}
	var array = primes[0:6]
	tosend := &pb.Array{Ints: array}

	words := [3]string{"hello", "three", "words"}
	var array2 = words[0:3]
	tosend2 := &pb.StringArray{S: array2}

	r, err := client.SumArray(context.Background(), tosend)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%t", r.Result)
	r2, err2 := client.SumWords(context.Background(), tosend2)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%t", r2.S)
}
