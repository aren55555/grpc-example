package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/aren55555/grpc-playground/garage_server/garage"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

var cars = []pb.Car{
	pb.Car{Make: "Volkswagen", Model: "R32", Year: 2008, Horsepower: 250},
	pb.Car{Make: "Honda", Model: "Civic Si", Year: 2010, Horsepower: 197},
	pb.Car{Make: "Ford", Model: "F-150 King Ranch", Year: 2002, Horsepower: 260},
}

// server is used to implement garage.GarageServer.
type server struct{}

// Random Car implements garage.GarageServer
func (s *server) RandomCar(ctx context.Context, cr *pb.CarRequest) (*pb.Car, error) {
	fmt.Printf("A request for a random car was recieved; it was number %v\n", cr.RequestNumber)
	seed := rand.NewSource(time.Now().Unix())
	r := rand.New(seed)
	car := &cars[r.Intn(len(cars))]

	return car, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGarageServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
