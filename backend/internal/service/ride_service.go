package service

import "rotasflex/internal/domain"

type RideService struct {}

func NewRideService() *RideService {
	return &RideService{}
}

func (s *RideService) CreateRide(userID int) domain.Ride {
	return domain.Ride{
		UserID: userID,
		Status: "created",
	}
}
