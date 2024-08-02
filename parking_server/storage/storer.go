package storage

import (
	"time"

	pb "github.com/aren55555/grpc-example/parking_server/grpc"
)

type CarDetails struct {
	Id        string
	Car       pb.Car
	EnteredAt time.Time
}

type Storer interface {
	AddCar(car *pb.Car) (CarDetails, error)
	RemoveCar(carId string) (CarDetails, error)
	GetMap() map[string]CarDetails
}
