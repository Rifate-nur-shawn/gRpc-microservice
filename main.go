package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
    port := ":50051"
    lis , err := net.Listen("tcp",port )
    if err!= nil{
        log.Fatal("Faield to listen",err)
    }
    grpcServer := grpc.NewServer()


//skiping something here

log.Println("service is running on port", port)

 err = grpcServer.Serve(lis)
 if err!= nil{
    log.Fatal("Faield to serve",err)
 }
}