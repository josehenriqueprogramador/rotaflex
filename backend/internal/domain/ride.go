package domain

type Ride struct {
	ID       int
	UserID   int
	DriverID int
	Status   string
	Price    float64
}
