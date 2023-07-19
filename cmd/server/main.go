package main

import (
	"context"
	"fmt"
	"net"

	"github.com/davimelovasc/go-grpc/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello %v", in.GetName())}, nil
}

func main() {
	port := 9000
	fmt.Printf("Running gRPC server on port %v\n", port)
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Server{})

	if err = s.Serve(l); err != nil {
		panic(err)
	}
}
