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
	whiteCars := parkingService.FindAllWhiteCars()

	// Check if the white car is found
	if len(whiteCars) != 1 {
		t.Errorf("Expected 1 white car, found %d", len(whiteCars))
	}

	// Check the details of the found white car

}

func TestParkingService_FindByColorAndModel(t *testing.T) {
	// Create parking lots, attendant, police department, and parking service
	parkingLots := make([]*ParkingLot, 4)
	parkingSpotLists := make([][][]ParkingSpot, 4)

	for i := range parkingLots {
		parkingLots[i] = NewParkingLot(i+1, 6, 6)
		parkingSpotLists[i] = make([][]ParkingSpot, 6)

		for j := range parkingSpotLists[i] {
			parkingSpotLists[i][j] = make([]ParkingSpot, 6)
		}
	}

	attendant := NewParkingAttendant()
	securityStaff := &SecurityStaff{}
	parkingService := NewParkingService(parkingLots, attendant, securityStaff)

	// Park a vehicle
	vehicle := Vehicle{
		LicensePlate: "ABC123",
		Color:        "Red",
		Model:        "Sedan",
	}
	_, err := attendant.AssignSpot(parkingLots, parkingSpotLists, &vehicle)
	if err != nil {
		t.Errorf("Error parking vehicle: %v", err)
	}

	// Search for the vehicle by color and model
	result, err := parkingService.FindByColorAndModel("Red", "Sedan")
	if err != nil {
		t.Errorf("Error searching for vehicle: %v", err)
	}

	// Check the result
	expectedResult := "Vehicle with Color Red and Model Sedan is parked at spot a1 in Parking Lot 1"
	if result != expectedResult {
		t.Errorf("Expected result: %s, got: %s", expectedResult, result)
	}

	// Search for a nonexistent vehicle
	result, err = parkingService.FindByColorAndModel("Blue", "SUV")
	if err == nil {
		t.Errorf("Expected error for nonexistent vehicle, but got nil")
	}
}
