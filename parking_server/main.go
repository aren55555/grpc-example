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
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *httpServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	storage := storage.New(*maxNumCars)
	parkingService := parking.NewService(storage)
	s := grpc.NewServer()
	pb.RegisterParkingServiceServer(s, parkingService)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("Server started on port %v", *httpServerPort)
}
