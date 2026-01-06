package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/rifate-nur-shawn/gRpc-microservice/pb/proto"
	fw "github.com/rifate-nur-shawn/gRpc-microservice/pb/proto/farewell"
	cal "github.com/rifate-nur-shawn/gRpc-microservice/pb/proto/stream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedGreeterServer
	fw.UnimplementedFarewellServer
}

type server1 struct {
	cal.UnimplementedCalculatorServer
}

func (s *server1) Adder(ctx context.Context, req *cal.AddRequest) (*cal.AddResponse, error) {
	result := req.A + req.B
	log.Printf("Adder called with A=%d, B=%d, Result=%d", req.A, req.B, result)
	return &cal.AddResponse{
		Result: result,
	}, nil
}

func (s *server1) GenarateFibonacci(req *cal.FibonacciRequest, stream cal.Calculator_GenarateFibonacciServer) error {
	n := req.N
	a, b := 0, 1
	for i := 0; i < int(n); i++ {
		err := stream.Send(&cal.FibonacciResponse{
			Value: int32(a),
		})

		if err != nil {
			return err
		}
		a, b = b, a+b
		time.Sleep(time.Second)
	}
	return nil
}

func (s *server1) BiStreaming(stream cal.Calculator_BiStreamingServer) error {
	log.Println("BiStreaming started")

	for {
		req, err := stream.Recv()
		if err != nil {
			log.Printf("BiStreaming finished: %v", err)
			return err
		}

		log.Printf("Received: message='%s', number=%d", req.Message, req.Number)

		// Process the request and send response
		response := &cal.BiResponse{
			Message: "Processed: " + req.Message,
			Result:  req.Number * 2,
		}

		if err := stream.Send(response); err != nil {
			log.Printf("Error sending response: %v", err)
			return err
		}

		log.Printf("Sent: message='%s', result=%d", response.Message, response.Result)
	}
}

func (f *server) Fare(ctx context.Context, req *fw.FarewellRequest) (*fw.FarewellResponse, error) {
	lastmessage := req.Name
	log.Println("last message", lastmessage)
	return &fw.FarewellResponse{
		Message: lastmessage,
	}, nil
}

func (c *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	result := req.Input
	year := req.Year
	log.Println("Name: ", result, "Year", year)
	return &pb.GreetResponse{
		Result: result,
		Year:   year,
	}, nil
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponce, error) {
	sum := req.A + req.B
	log.Println("sum: ", sum)
	return &pb.AddResponce{
		Sum: sum,
	}, nil
}

// func(s *server) Add(ctx context.Context , req *pb.AddRequest)(*pb.AddResponse,error){
// 	return &pb.AddResponse{
// 		Result: req.A+req.B,
// 	},nil
// }

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
	fw.RegisterFarewellServer(grpcServer, &server{})
	cal.RegisterCalculatorServer(grpcServer, &server1{})

	log.Println("server running on port :", port)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve ", err)
	}

}
