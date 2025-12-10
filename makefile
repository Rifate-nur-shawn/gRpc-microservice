gen:
	protoc --go_out=pb --go-grpc_out=pb proto/*.proto

clean:
	rm -rf pb/*.pb.go
run:
	go run main.go