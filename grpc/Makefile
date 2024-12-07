generate:
	protoc -I ./proto ./proto/*.proto --go_out=./server --go-grpc_out=./server && protoc -I ./proto ./proto/*.proto --go_out=./client --go-grpc_out=./client

.PHONY: server
server:
	cd server && go run main.go
