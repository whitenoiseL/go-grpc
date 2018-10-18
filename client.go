package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/whitenoiseL/go-grpc/helloworld"
)

const (
	address = "localhost:50051"
	//address = "localhost:50051"
	hi = "world"
)

var (
	c helloworld.GreetingClient
	ctx context.Context
	cancel context.CancelFunc
	conn *grpc.ClientConn
	err error
)

var ch = make(chan int)

func clientStart() {
	// Set up a connection to the server.
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c = helloworld.NewGreetingClient(conn)
	// Set time to wait
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	<- ch
}

func sayHi() string{
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Request:hi})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Response
}

func closeConn(){
	conn.Close()
	cancel()
}

func main() {

	go clientStart()
	ch <- 0

	var re string

	for i := 0; i<1; i++{
		re = sayHi()
	}

	log.Print("Greeting: ", re)
	closeConn()
}