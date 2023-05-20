// Welcome to your channel go guruji

// Topic how to implement uniary operation in grpc

package main

import (
	"context"
	"errors"
	"fmt"
	proto "grpc/protoc"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {

	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}
	srv := grpc.NewServer() // engine
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) ServerReply(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("recieve request from client", req.SomeString)
	fmt.Println("hello from server")
	return &proto.HelloResponse{}, errors.New("")
}
