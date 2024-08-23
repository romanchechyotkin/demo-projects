package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"kv-store/logger"
	"kv-store/server"
	"kv-store/store"
)

func main() {
	log.Println("App is running")
	
	log.Println("Logger init")
	l, err := logger.NewFileTransactionLogger("transaction.log")

	log.Println("Create store")
	s := store.NewStore(l)	
	err = s.InitLog()
	if err != nil {
		log.Println(err)
	}

	log.Println("Server is running")
	
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	
	log.Fatal(server.RunServer(ctx, s))	
}

