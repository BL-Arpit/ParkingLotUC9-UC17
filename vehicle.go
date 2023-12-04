package main

import "time"

// Vehicle represents a vehicle in the parking lot.
type Vehicle struct {
	LicensePlate string
	Color        string
	Model        string
	Brand        string
	ParkingSpot  string
	ParkingTime  time.Time
	Handicapped  bool
}
