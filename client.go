package main

import (
	"google.golang.org/grpc"
	"log"
	"time"
	"golang.org/x/net/context"

	"github.com/whitenoiseL/go-grpc/helloworld"
)

const (
	address     = "149.28.90.228:50051"
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

func sayhi() string{
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
	//defer conn.Close()
	//defer cancel()

	var re string

	for i := 0; i<10; i++{
		re = sayhi()
	}

	log.Print("Greeting: ", re)
	closeConn()
}