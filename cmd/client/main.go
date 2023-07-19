package main

import (
	"context"
	"fmt"
	"log"

	"github.com/davimelovasc/go-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	sPort := 9000 //grpc server port
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", sPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{
		Name: "Davi",
	}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
