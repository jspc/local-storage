package main

import (
	"log"
	"net"
	"os"

	"github.com/j-and-j-global/storage-service"
	"google.golang.org/grpc"
)

var (
	Dir = os.Getenv("DIR")
	ID  = os.Getenv("ID")
)

func main() {
	log.Print("Starting")

	server := Server{
		ID:        ID,
		Directory: Dir,
		Mode:      0600,
	}

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Panic(err)
	}

	grpcServer := grpc.NewServer()
	storage.RegisterStorageServer(grpcServer, server)

	log.Print("Starting server")

	grpcServer.Serve(lis)
}
