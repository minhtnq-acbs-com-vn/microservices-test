package main

import (
	"google.golang.org/grpc"
	"log"
	"microservice-test/book/service"
	"microservice-test/proto/book"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const connectionString = "mongodb://host.docker.internal:27017"

func main() {
	lis, err := net.Listen("tcp", ":11000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	bookService := service.New(connectionString)

	book.RegisterBookServiceServer(grpcServer, bookService)

	go func() {
		log.Println("Server Book is listening on port: 11000")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("Server Book is shutting down")
}
