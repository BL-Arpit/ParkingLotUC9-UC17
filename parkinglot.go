package main

import (
	"fmt"
)

// ParkingLot represents the parking lot with fixed size and parked vehicles.
type ParkingLot struct {
	ID              int
	Rows            int
	Columns         int
	AvailableSpaces int
	ParkedVehicles  []Vehicle
}

// NewParkingLot creates a new parking lot with the given number of rows and columns.
func NewParkingLot(id, rows, columns int) *ParkingLot {
	return &ParkingLot{
		ID:              id,
		Rows:            rows,
		Columns:         columns,
		AvailableSpaces: rows * columns,
		ParkedVehicles:  make([]Vehicle, 0, rows*columns),
	}
}

// ContainsSpot checks if the parking lot contains a specific parking spot.
func (p *ParkingLot) ContainsSpot(spotID string) bool {
	for _, vehicle := range p.ParkedVehicles {
		if vehicle.ParkingSpot == spotID {
			return true
		}
	}
	return false
}

// Status prints the current status of the parking lot.
func (p *ParkingLot) Status() {
	fmt.Printf("Parking Lot %d:\n", p.ID)
	fmt.Printf("Total Spaces: %d\n", p.Rows*p.Columns)
	fmt.Printf("Available Spaces: %d\n", p.AvailableSpaces)
	fmt.Printf("Parked Vehicles:\n")

	if len(p.ParkedVehicles) == 0 {
		fmt.Println("No vehicles parked.")
	} else {
		for _, vehicle := range p.ParkedVehicles {
			fmt.Printf("- License Plate: %s, Color: %s, Model: %s, Brand: %s, Parking Spot: %s\n",
				vehicle.LicensePlate, vehicle.Color, vehicle.Model, vehicle.Brand, vehicle.ParkingSpot)
		}
	}
}
