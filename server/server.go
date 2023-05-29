// Welcome to your channel go guruji

// Topic how to implement uniary operation in grpc

package main

import (
	"fmt"
	proto "grpc/protoc"
	"net"
	"time"

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

func (s *server) ServerReply(req *proto.HelloRequest, strem proto.Example_ServerReplyServer) error {
	fmt.Println(req.SomeString)
	time.Sleep(5 * time.Second)
	girlReply := []*proto.HelloResponse{
		{Reply: "Kya hua"},
		{Reply: "Sorry"},
		{Reply: "Man jao na"},
		{Reply: "Please"},
	}
	for _, msg := range girlReply {
		err := strem.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
