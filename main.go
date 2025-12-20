package main

import (
    "context"
    "log"
    pb "github.com/rifate-nur-shawn/gRpc-microservice/pb"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func main() {
    // Step 4.1: Connect to the server
    conn, err := grpc.Dial("localhost:50051", 
        grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Step 4.2: Create client using the GENERATED function
    client := pb.NewProductServiceClient(conn)

    // Step 4.3: Call CreateProduct (send a request)
    createResp, err := client.CreateProduct(context.Background(), &pb.CreateProductRequest{
        Product: &pb.Product{
            Id:       "123",
            Name:     "Gaming Laptop",
            Price:    2500.00,
            Category: "Electronics",
        },
    })
    if err != nil {
        log.Fatalf("Error creating product: %v", err)
    }
    log.Printf("Product created with ID: %s\n", createResp.Id)

    // Step 4.4: Call GetProduct (send another request)
    getResp, err := client.GetProduct(context.Background(), &pb.GetProductRequest{
        Id: "123",
    })
    if err != nil {
        log.Fatalf("Error getting product: %v", err)
    }
    log.Printf("Product: %s, Price: $%.2f\n", getResp.Product.Name, getResp.Product.Price)
}