gen:
	protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto
	protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/farewell/*.proto
	protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/stream/*.proto

clean:
	rm -rf pb/proto/*.pb.go pb/proto/farewell/*.pb.go pb/proto/stream/*.pb.go

run:
	go run main.go