package storage

import (
	"errors"
	"sync"
	"time"

	pb "github.com/aren55555/grpc-example/parking_server/grpc"
	"github.com/google/uuid"
)

var _ Storer = &storage{}

type storage struct {
	mu         sync.RWMutex
	cars       map[string]CarDetails
	maxNumCars uint
}

func New(maxNumCars uint) Storer {
	return &storage{
		mu:         sync.RWMutex{},
		cars:       make(map[string]CarDetails),
		maxNumCars: maxNumCars,
	}
}

func (s *storage) AddCar(car *pb.Car) (CarDetails, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if uint(len(s.cars)) >= s.maxNumCars {
		return CarDetails{}, errors.New("parking lot is full")
	}

	c := CarDetails{
		Id:        uuid.New().String(),
		Car:       *car,
		EnteredAt: time.Now(),
	}

	s.cars[c.Id] = c

	return CarDetails{
		Id:        uuid.New().String(),
		Car:       *car,
		EnteredAt: time.Now(),
	}, nil
}

func (s *storage) RemoveCar(carId string) (CarDetails, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, ok := s.cars[carId]
	if !ok {
		return CarDetails{}, errors.New("car not found")
	}

	return CarDetails{
		Id:        c.Id,
		Car:       c.Car,
		EnteredAt: c.EnteredAt,
	}, nil
}

func (s *storage) GetMap() map[string]CarDetails {
	s.mu.RLock()
	defer s.mu.RUnlock()

	m := make(map[string]CarDetails)
	for k, v := range s.cars {
		m[k] = v
	}

	return m
}
