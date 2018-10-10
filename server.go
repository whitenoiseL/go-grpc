package main

import (
	pb "./helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"fmt"
)
const (
	grpcPort = ":50051"
)

type Server struct{}

func (s *Server) SayHello(context context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error){
	return &pb.HelloResponse{Response: "Hello " + in.Request}, nil
}

func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloworldServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)


}