package parking

import (
	"context"

	pb "github.com/aren55555/grpc-example/parking_server/grpc"
	"github.com/aren55555/grpc-example/parking_server/storage"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Our implementation of the parking service
type service struct {
	storage storage.Storer
	pb.UnimplementedParkingServiceServer
}

// Ensures that an instance implements the proto type
var _ pb.ParkingServiceServer = &service{}

func NewService(storage storage.Storer) pb.ParkingServiceServer {
	return &service{
		storage: storage,
	}
}

// Use to park the car in the parking lot
func (s *service) ParkCar(ctx context.Context, req *pb.ParkCarRequest) (*pb.ParkCarResponse, error) {
	cd, err := s.storage.AddCar(req.Car)
	if err != nil {
		return nil, err
	}

	return &pb.ParkCarResponse{
		Id:        cd.Id,
		EnteredAt: timestamppb.New(cd.EnteredAt),
	}, nil
}

// Used to pickup a car from the parking lot
func (s *service) CollectCar(ctx context.Context, req *pb.CollectCarRequest) (*pb.CollectCarResponse, error) {
	cd, err := s.storage.RemoveCar(req.Id)
	if err != nil {
		return nil, err
	}

	exitedAt := timestamppb.Now()

	return &pb.CollectCarResponse{
		Car:         &cd.Car,
		EnteredAt:   timestamppb.New(cd.EnteredAt),
		ExitedAt:    exitedAt,
		ChargeCents: 0, // TODO: compute based on time diff
	}, nil
}

func (s *service) GetStatus(ctx context.Context, req *pb.GetStatusRequest) (*pb.GetStatusResponse, error) {
	m := s.storage.GetMap()

	cars := make(map[string]*pb.StoredCar)

	for k, cd := range m {
		cars[k] = &pb.StoredCar{
			Car:       &cd.Car,
			EnteredAt: timestamppb.New(cd.EnteredAt),
		}
	}

	return &pb.GetStatusResponse{
		Cars: cars,
	}, nil
}
