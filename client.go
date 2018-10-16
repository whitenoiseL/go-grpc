package main

import (
	"google.golang.org/grpc"
	"log"
	"time"
	"golang.org/x/net/context"

	"github.com/whitenoiseL/go-grpc/helloworld"
)

const (
	address     = "localhost:50051"
	hi = "world!"
)

var (
	c helloworld.GreetingClient
	ctx context.Context
	cancel context.CancelFunc
	conn *grpc.ClientConn
	err error
)

func clientStart() {
	// Set up a connection to the server.
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c = helloworld.NewGreetingClient(conn)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)

}

func sayhi() string{
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Request:hi})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Response
}

func main() {

	clientStart()
	defer conn.Close()
	defer cancel()

	re := sayhi()

	log.Printf("Greeting: %s", re)
}