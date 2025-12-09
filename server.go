package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main(){
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000 :%v",err)
	}

	grpcserver := grpc.NewServer()

	if err := grpcserver.Serve(lis); err != nil {
		log.Fatalf("Failed to listen grpc port 9000 :%v", err)
	}
}