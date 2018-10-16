package main

import (
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/go-redis/redis"

	pb "github.com/whitenoiseL/go-grpc/helloworld"
)

const (
	grpcPort = ":50051"
)

type Server struct{}
var client *redis.Client

func (s *Server) SayHello(context context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error){
	//operation()
	return &pb.HelloResponse{Response: "Hello " + in.Request}, nil
}

func grpc_start(){
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetingServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}

func redis_start() (*redis.Client, error){
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client,nil
	// Output: PONG <nil>
}

func operation() {
	err := client.Set("key", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	// Output: key value
	// key2 does not exist
}

func main() {
	go grpc_start()
	//go redis_start()
	c := make(chan int)
	fmt.Println("grpc service started...")
	c <- 0

}