package main

import (
	"context"
	pb "grpc/proto"
	"log"

	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func main() {

	go func() {
		//mux
		mux := runtime.NewServeMux()

		//Register
		pb.RegisterTestApiHandlerServer(context.Background(), mux, &testApiServer{})

		//http Server
		log.Fatalln(http.ListenAndServe("localhost:8000", mux))

	}()

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
