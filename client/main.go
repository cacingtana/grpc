package main

import (
	"context"
	"fmt"
	pb "grpc/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)

	response, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello Everyone"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(response)
	fmt.Println(response.Msg)
}
