package main

// Welcome to channel go guruji

// Topic grpc client streaming

import (
	"context"
	"fmt"
	proto "grpc/protoc"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.ExampleClient

func main() {
	// Connection to internal grpc server
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = proto.NewExampleClient(conn)
	// implement rest api
	r := gin.Default()
	r.GET("/sent", clientConnectionServer)
	r.Run(":8000") // 8080

}

func clientConnectionServer(c *gin.Context) {

	req := []*proto.HelloRequest{
		{SomeString: "Request 1"},
		{SomeString: "Request 2"},
		{SomeString: "Request 3"},
		{SomeString: "Request 4"},
		{SomeString: "Request 5"},
		{SomeString: "Request 6"},
	}

	stream, err := client.ServerReply(context.TODO())
	if err != nil {
		fmt.Println("Something error")
		return
	}
	for _, re := range req {
		err = stream.Send(re)
		if err != nil {
			fmt.Println("request not fulfil")
			return
		}

	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("there is some error occure ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message_count": response,
	})

}
