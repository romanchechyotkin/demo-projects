generate_proto:
	protoc -I ./proto ./proto/proto.proto --go_out=./proto --go-grpc_out=./proto

build:
	go build -o bin/bin main.go && ./bin/bin