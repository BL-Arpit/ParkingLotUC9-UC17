package main

import (
	"fmt"
	"testing"
)

func TestParkingAttendant(t *testing.T) {
	rows := 6
	columns := 6

	// Create parking lots
	//4 lots
	parkingLots := make([]*ParkingLot, 4)
	parkingSpotLists := make([][][]ParkingSpot, 4)

	for i := range parkingLots {
		parkingLots[i] = NewParkingLot(i, rows, columns)
		parkingSpotLists[i] = make([][]ParkingSpot, rows)

		for j := range parkingSpotLists[i] {
			parkingSpotLists[i][j] = make([]ParkingSpot, columns)
		}
	}

	// Create parking attendant
	attendant := NewParkingAttendant()

	// Simulate parking operations
	for j := 0; j < 12; j++ { // Assuming 4 parking lots, so 3 vehicles per lot
		vehicle := Vehicle{
			LicensePlate: "ABC123",
			Color:        "Red",
			Model:        "Sedan",
		}

		// Assign parking spot
		_, _ = attendant.AssignSpot(parkingLots, parkingSpotLists, &vehicle)
	}

	// Check if each parking lot is evenly filled
	for i, lot := range parkingLots {
		expectedCount := 3 // Assuming 3 vehicles per lot
		if len(lot.ParkedVehicles) != expectedCount {
			t.Errorf("Parking lot %d should have %d vehicles, got %d", i+1, expectedCount, len(lot.ParkedVehicles))
		}
	}
}

func TestParkingAttendant_HandicappedParking(t *testing.T) {
	rows := 6
	columns := 6

	// Create parking lots
	parkingLots := make([]*ParkingLot, 4)
	parkingSpotLists := make([][][]ParkingSpot, 4)

	for i := range parkingLots {
		parkingLots[i] = NewParkingLot(i+1, rows, columns)
		parkingSpotLists[i] = make([][]ParkingSpot, rows)

		for j := range parkingSpotLists[i] {
			parkingSpotLists[i][j] = make([]ParkingSpot, columns)
		}
	}

	// Create parking attendant
	attendant := NewParkingAttendant()

	//add 10 cars to parking lot
	for k := 0; k < 10; k++ {
		vehicle := Vehicle{
			LicensePlate: "ABC123",
			Color:        "Red",
			Model:        "Sedan",
		}

		_, err := attendant.AssignSpot(parkingLots, parkingSpotLists, &vehicle)
		if err != nil {
			t.Errorf("Error assigning parking spot for regular vehicle: %v", err)
		}
	}

	// Create handicapped vehicle
	handicappedVehicle := Vehicle{
		LicensePlate: "ABC123",
		Color:        "Blue",
		Model:        "SUV",
		Handicapped:  true,
	}

	// Assign parking spot for handicapped vehicle
	_, err := attendant.AssignSpot(parkingLots, parkingSpotLists, &handicappedVehicle)
	if err != nil {
		t.Errorf("Error assigning parking spot for handicapped vehicle: %v", err)
	}

	// Check if the handicapped vehicle is assigned to the nearest parking lot
	expectedLotID := 1 // The nearest parking lot for handicapped vehicles
	actualLotID := int(handicappedVehicle.ParkingSpot[1] - '0')
	if expectedLotID != actualLotID {
		t.Errorf("Expected handicapped vehicle to be assigned to parking lot %d, got %d", expectedLotID, actualLotID)
	}

	// Print parking status
	for i, lot := range parkingLots {
		fmt.Printf("Parking Lot %d Status:\n", i+1)
		//fmt.Printf("Total Spaces: %d\n", lot.TotalSpaces)
		fmt.Printf("Available Spaces: %d\n", lot.AvailableSpaces)
		fmt.Printf("Parked Vehicles:\n")
		for _, vehicle := range lot.ParkedVehicles {
			fmt.Printf("- License Plate: %s, Parking Spot: %s\n", vehicle.LicensePlate, vehicle.ParkingSpot)
		}
		fmt.Println()
	}
}
