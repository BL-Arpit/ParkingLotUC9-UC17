package main

import (
	"fmt"
)

// PoliceDepartment represents the police department managing parking lots.
type PoliceDepartment struct {
	parkingService *ParkingService
}

// NewPoliceDepartment creates a new police department.
func NewPoliceDepartment(parkingService *ParkingService) *PoliceDepartment {
	return &PoliceDepartment{
		parkingService: parkingService,
	}
}

// DisplayParkingLotStatus displays the status of all parking lots.
func (pd *PoliceDepartment) DisplayParkingLotStatus() {
	pd.parkingService.Status()
}

// FindAndDisplayWhiteCars finds and displays information about white cars.
func (pd *PoliceDepartment) FindAndDisplayWhiteCars() {
	whiteCars := pd.parkingService.FindAllWhiteCars()

	if len(whiteCars) == 0 {
		fmt.Println("No white cars found in the parking lots.")
	} else {
		fmt.Println("White cars found:")
		for _, vehicle := range whiteCars {
			fmt.Printf("License Plate: %s, Color: %s, Model: %s, Parking Spot: %s\n",
				vehicle.LicensePlate, vehicle.Color, vehicle.Model, vehicle.ParkingSpot)
		}
	}
}

func (pd *PoliceDepartment) SearchByColorAndModel() {
	var color, model string

	fmt.Print("Enter Color: ")
	fmt.Scan(&color)
	fmt.Print("Enter Model: ")
	fmt.Scan(&model)

	result, err := pd.parkingService.FindByColorAndModel(color, model)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
