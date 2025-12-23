package main

import (
	"context"
	"log"
	"net"

	pb "github.com/rifate-nur-shawn/gRpc-microservice/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedGreeterServer
}

func (c *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Result: req.Input,
		Year:   req.Year,
	}, nil
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponce, error) {
	sum := req.A + req.B
	log.Println("sum: ", sum)
	return &pb.AddResponce{
		Sum: sum,
	}, nil
}

func main() {

	cert := "cert.pem"
	key := "key.pem"

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen ", err)

	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)

	if err != nil {
		log.Fatalln("Failed to load credentials", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterCalculateServer(grpcServer, &server{})
	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("server running on port :", port)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve ", err)
	}

}
