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
func TestAssignSpotForSUV(t *testing.T) {
	rows := 6
	columns := 6

	// Create parking lots
	parkingLot1 := NewParkingLot(1, rows, columns)
	parkingLot2 := NewParkingLot(2, rows, columns)
	parkingLot3 := NewParkingLot(3, rows, columns)
	parkingLot4 := NewParkingLot(4, rows, columns)

	parkingLots := []*ParkingLot{parkingLot1, parkingLot2, parkingLot3, parkingLot4}

	attendant := NewParkingAttendant()
	securityStaff := &SecurityStaff{}
	parkingService := NewParkingService(parkingLots, attendant, securityStaff)

	parkingSpotLists := make([][][]ParkingSpot, len(parkingLots))

	for i := range parkingSpotLists {
		parkingSpotLists[i] = make([][]ParkingSpot, rows)
		for j := range parkingSpotLists[i] {
			parkingSpotLists[i][j] = make([]ParkingSpot, columns)
		}
	}

	// Park 9 normal vehicles
	for j := 0; j < 9; j++ {
		vehicle := Vehicle{
			LicensePlate: fmt.Sprintf("ABC%d", j+1),
			Color:        "Red",
			Model:        "Sedan",
		}

		err := parkingService.Park(vehicle, parkingSpotLists)
		if err != nil {
			t.Fatalf("Error parking vehicle: %v", err)
		}
	}

	// Park 3 SUVs and check if they are parked in the lot with the highest available space
	for j := 0; j < 2; j++ {
		suvVehicle := Vehicle{
			LicensePlate: fmt.Sprintf("SUV%d", j+1),
			Color:        "Black",
			Model:        "SUV",
		}

		err := parkingService.Park(suvVehicle, parkingSpotLists)
		if err != nil {
			t.Fatalf("Error parking SUV: %v", err)
		}

		// Print the parking status after parking each SUV
		parkingService.Status()
	}
	//unpark from parking lot 1
	for i := 0; i < 3; i++ {
		err := parkingService.Unpark(fmt.Sprintf("ABC%d", i+1), parkingSpotLists)
		if err != nil {
			t.Fatalf("Error unparking vehicle: %v", err)
		}
	}

	for j := 0; j < 2; j++ {
		suvVehicle := Vehicle{
			LicensePlate: fmt.Sprintf("SUV%d", j+1),
			Color:        "Black",
			Model:        "SUV",
		}

		err := parkingService.Park(suvVehicle, parkingSpotLists)
		if err != nil {
			t.Fatalf("Error parking SUV: %v", err)
		}

		// Print the parking status after parking each SUV
		parkingService.Status()
	}
}
func (pd *PoliceDepartment) DisplayAllParkedVehicles() {
	fmt.Println("Parked Vehicles:")

	for _, lot := range pd.parkingService.parkingLots {
		for _, vehicle := range lot.ParkedVehicles {
			fmt.Printf("License Plate: %s, Color: %s, Model: %s, Brand: %s, Parking Spot: %s\n",
				vehicle.LicensePlate, vehicle.Color, vehicle.Model, vehicle.Brand, vehicle.ParkingSpot)
		}
	}

	if len(pd.parkingService.parkingLots) == 0 {
		fmt.Println("No vehicles are currently parked.")
	}
}
