package main

import (
	"context"
	"log"

	pb "github.com/Lazaro-Barros/gRPC/greet/proto"
)

func (s Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("The function Greet was invoked with the parameter: %v\n", in)
	return &pb.GreetResponse{
		ResultMessage: "hello " + in.FirstName,
	}, nil
}
