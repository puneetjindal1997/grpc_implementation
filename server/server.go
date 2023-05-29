// Welcome to your channel go guruji

// Topic how to implement uniary operation in grpc

package main

import (
	"fmt"
	proto "grpc/protoc"
	"io"
	"net"
	"strconv"

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

func (s *server) ServerReply(strem proto.Example_ServerReplyServer) error {
	total := 0 // total messages client
	for {
		request, err := strem.Recv()
		if err == io.EOF {
			return strem.SendAndClose(&proto.HelloResponse{
				Reply: strconv.Itoa(total),
			})
		}
		if err != nil {
			return err
		}

		total++
		fmt.Println(request)
	}
}
