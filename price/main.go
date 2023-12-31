package main

import (
	"google.golang.org/grpc"
	"log"
	"microservice-test/price/service"
	"microservice-test/proto/price"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	priceService := service.New()

	price.RegisterPriceServiceServer(grpcServer, priceService)

	go func() {
		log.Println("Server Price is listening on port: 10000")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("Server Price is shutting down")
}
