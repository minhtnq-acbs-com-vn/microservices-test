package main

import (
	"google.golang.org/grpc"
	"log"
	"microservice-test/helper/service"
	"microservice-test/proto/helper"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const connectionString = "mongodb://localhost:27017"

func main() {
	lis, err := net.Listen("tcp", ":12000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	helperService := service.New(connectionString)

	helper.RegisterHelperServer(grpcServer, helperService)

	go func() {
		log.Println("Server Helper is listening on port: 12000")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("Server Helper is shutting down")
}
