package main

import (
	"context"
	"fmt"
	"github.com/NaraLuwan/grpc-greeter/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to connect: %v", err)
		return
	}
	defer conn.Close()

	greeterClient := pb.NewGreeterClient(conn)

	helloReply, err := greeterClient.SayHello(context.Background(), &pb.HelloRequest{
		Name: "Tom",
	})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	fmt.Printf("Greeting: %s !\n", helloReply.Message)
}
