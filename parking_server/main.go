package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/aren55555/grpc-example/parking_server/grpc"
	"github.com/aren55555/grpc-example/parking_server/handler"
	"github.com/aren55555/grpc-example/parking_server/parking"
	"github.com/aren55555/grpc-example/parking_server/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	httpServerPort = flag.Uint("httpServerPort", 80, "The port the HTTP server will listen on")
	grpcServerPort = flag.Uint("grpcServerPort", 50051, "The port the gRPC server will listen on")
	maxNumCars     = flag.Uint("maxNumCars", 10, "The maximum number of cars allowed to be parked")
)

func main() {
	log.Printf("Parsing flags...")
	flag.Parse()

	// Init the storage client
	s := storage.New(*maxNumCars)

	// Start HTTP Server
	go startHttp(s)

	// Start gRPC Server
	go startGRPC(s)

	// Block forever
	select {}
}

func startHttp(s storage.Storer) {
	handler := handler.NewHandler(s)

	log.Printf("Starting HTTP server listener on port %v...", *httpServerPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", *httpServerPort), handler); err != nil {
		log.Fatalf("failed to  http.ListenAndServe: %v", err)
	}
}

func startGRPC(s storage.Storer) {
	log.Printf("Starting net listener on port %v...", *grpcServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *grpcServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Registering gRPC parking services")
	grpcServer := grpc.NewServer()
	pb.RegisterParkingServiceServer(grpcServer, parking.NewService(s))
	reflection.Register(grpcServer)

	log.Printf("Server starting on port %v", *grpcServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to grpcServer.Serve: %v", err)
	}
}
