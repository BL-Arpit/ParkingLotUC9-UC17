package main

import (
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
