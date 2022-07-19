package main

import (
	"log"
	"net"

	pb "github.com/Lazaro-Barros/gRPC/greet/proto"
	"google.golang.org/grpc"
)

var address = "127.0.0.1:5005"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Fail to listen on %v\n", err)
	}

	log.Printf("Listenig on %s\n", address)

	server := grpc.NewServer()
	pb.RegisterGreetServiceServer(server, &Server{})
	if err = server.Serve(listener); err != nil {
		log.Fatalf("Fail to listen on %v\n", err)
	}
}
