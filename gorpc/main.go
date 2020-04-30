package main

import (
	"google.golang.org/grpc"
	"gorpc/service"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	service.RegisterUserServiceServer(grpcServer, &UserServiceServerImpl{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve grpc: %s", err.Error())
	}
}
