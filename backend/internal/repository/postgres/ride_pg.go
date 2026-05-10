package postgres

import "rotasflex/internal/domain"

type RideRepository struct {}

func NewRideRepository() *RideRepository {
	return &RideRepository{}
}

func (r *RideRepository) Save(ride domain.Ride) error {
	// depois conecta no PostgreSQL
	return nil
}
