package main

import (
	"context"
	"fmt"
	"github.com/NaraLuwan/grpc-greeter/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type greeterImpl struct{}

func (greeter greeterImpl) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	// 创建gRPC服务器
	greeterServer := grpc.NewServer()
	// 在gRPC服务端注册服务
	pb.RegisterGreeterServer(greeterServer, &greeterImpl{})
	// 注册grpcurl所需的reflection服务，使用grpcurl命令行工具可直接调用
	reflection.Register(greeterServer)

	// Serve方法在listener上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应
	err = greeterServer.Serve(listener)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
