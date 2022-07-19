package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Lazaro-Barros/gRPC/greet/proto"
)

func DoGreet(c pb.GreetServiceClient) {
	var firstName string = ""
	fmt.Print("Insert your first name:")
	fmt.Scan(&firstName)
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: firstName,
	})
	if err != nil {
		log.Fatalf("Fail on use greet service: %v\n", err)
	}
	log.Printf("Greeting Response: %v\n", res.ResultMessage)
}
