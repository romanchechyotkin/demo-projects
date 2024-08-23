package main

import (
	"cars_searching_filtering/proto/pb"
	"cars_searching_filtering/server"
	"github.com/elastic/go-elasticsearch/v8"

	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("elasticsearch microservice")

	listen, err := net.Listen("tcp", ":5500")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	info, err := client.Info()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer info.Body.Close()
	fmt.Println(info)

	srv, err := server.NewServer(client)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCarsManagementServer(grpcServer, srv)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
