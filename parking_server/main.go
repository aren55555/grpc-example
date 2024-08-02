package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/aren55555/grpc-example/parking_server/grpc"
	"github.com/aren55555/grpc-example/parking_server/parking"
	"github.com/aren55555/grpc-example/parking_server/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	httpServerPort = flag.Uint("port", 50051, "The port the HTTP server will listen on")
	maxNumCars     = flag.Uint("max_num_cars", 10, "The maximum number of cars allowed to be parked")
)

func main() {
	log.Printf("Parsing flags...")
	flag.Parse()

	log.Printf("Starting net listener on port %v...", *httpServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *httpServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Registering gRPC parking services")
	grpcServer := grpc.NewServer()
	pb.RegisterParkingServiceServer(grpcServer, parking.NewService(storage.New(*maxNumCars)))
	reflection.Register(grpcServer)

	log.Printf("Server starting on port %v", *httpServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
