package main

import (
	"log"

	pb "github.com/Lazaro-Barros/gRPC/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address = "127.0.0.1:5005"

func main() {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	defer connection.Close()

	client := pb.NewGreetServiceClient(connection)
	DoGreet(client)

}
