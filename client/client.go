package main

// Welcome to channel go guruji

// Topic grpc server streaming

import (
	"context"
	"fmt"
	proto "grpc/protoc"
	"io"
	"net/http"
	"time"

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
	stream, err := client.ServerReply(context.TODO(), &proto.HelloRequest{SomeString: "m rutha hua hu"})
	if err != nil {
		fmt.Println("Something error")
		return
	}
	count := 0
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("girlfriend messages:- ", message)
		time.Sleep(1 * time.Second)
		count++
	}
	c.JSON(http.StatusOK, gin.H{
		"message_count": count,
	})
}
