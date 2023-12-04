package main

import (
	"testing"
)

func TestPoliceDepartment_FindWhiteCars(t *testing.T) {
	rows := 6
	columns := 6

	// Create parking lots, attendant, police department, and parking service
	parkingLots := make([]*ParkingLot, 4)
	parkingSpotLists := make([][][]ParkingSpot, 4)

	for i := range parkingLots {
		parkingLots[i] = NewParkingLot(i+1, rows, columns)
		parkingSpotLists[i] = make([][]ParkingSpot, rows)

		for j := range parkingSpotLists[i] {
			parkingSpotLists[i][j] = make([]ParkingSpot, columns)
		}
	}

	attendant := NewParkingAttendant()
	parkingService := NewParkingService(parkingLots, attendant, &SecurityStaff{})
	//policeDepartment := NewPoliceDepartment(parkingService)

	// Park a white car
	whiteCar := Vehicle{
		LicensePlate: "WHT123",
		Color:        "White",
		Model:        "Sedan",
	}

	_, err := attendant.AssignSpot(parkingLots, parkingSpotLists, &whiteCar)
	if err != nil {
		t.Errorf("Error parking white car: %v", err)
	}

	// Call the function that finds white cars
	whiteCars := parkingService.FindAllWhiteCars();

	// Check if the white car is found
	if len(whiteCars) != 1 {
		t.Errorf("Expected 1 white car, found %d", len(whiteCars))
	}

	// Check the details of the found white car

}