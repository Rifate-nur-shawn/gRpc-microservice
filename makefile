gen:
	protoc --go_out=. --go-grpc_out=. chat.proto

clean:
	rm -rf chat/*.pb.go
run:
	go run server.go