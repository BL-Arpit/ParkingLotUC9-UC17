package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
	"time"
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

func TestPoliceDepartment_CheckForBMW(t *testing.T) {
	// Create parking lots, attendant, and police department
	rows := 6
	columns := 6
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
	securityStaff := &SecurityStaff{}
	parkingService := NewParkingService(parkingLots, attendant, securityStaff)
	policeDepartment := NewPoliceDepartment(parkingService)

	// Create a BMW car
	bmwCar := Vehicle{
		LicensePlate: "XYZ456",
		Color:        "Black",
		Model:        "Sedan",
		Brand:        "BMW",
	}

	// Assign a parking spot for the BMW car
	_, err := attendant.AssignSpot(parkingLots, parkingSpotLists, &bmwCar)
	assert.NoError(t, err, "Error assigning parking spot for BMW car")

	// Capture standard output and check printed messages
	expectedOutput := "Security staff notified: High-security measures activated.\n\nBMW Cars found. Security tightened.\n"
	actualOutput := captureOutput(func() {
		policeDepartment.CheckForBMW()
	})

	// Check if the expected output matches the actual output
	assert.Equal(t, expectedOutput, actualOutput, "Output mismatch")

}
func captureOutput(f func()) string {
	originalOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	capturedOutput, _ := io.ReadAll(r)
	os.Stdout = originalOutput

	return string(capturedOutput)
}
func TestParkingService_FindCarsParkedLast30Mins(t *testing.T) {
	rows := 6
	columns := 6

	// Create parking lots
	parkingLots := make([]*ParkingLot, 1)
	parkingSpotLists := make([][][]ParkingSpot, 1)

	for i := range parkingLots {
		parkingLots[i] = NewParkingLot(i+1, rows, columns)
		parkingSpotLists[i] = make([][]ParkingSpot, rows)

		for j := range parkingSpotLists[i] {
			parkingSpotLists[i][j] = make([]ParkingSpot, columns)
		}
	}

	// Create parking attendant
	attendant := NewParkingAttendant()

	// Create parking service
	parkingService := NewParkingService(parkingLots, attendant, securityStaffInstance)

	// Park a car
	car1 := Vehicle{
		LicensePlate: "ABC123",
		Color:        "Red",
		Model:        "Sedan",
	}
	_, err := attendant.AssignSpot(parkingLots, parkingSpotLists, &car1)
	if err != nil {
		t.Fatalf("Error parking car: %v", err)
	}

	// Park another car after 15 minutes
	time.Sleep(15 * time.Minute)
	car2 := Vehicle{
		LicensePlate: "XYZ789",
		Color:        "Blue",
		Model:        "SUV",
	}
	_, err = attendant.AssignSpot(parkingLots, parkingSpotLists, &car2)
	if err != nil {
		t.Fatalf("Error parking car: %v", err)
	}

	// Park another car after 45 minutes
	time.Sleep(30 * time.Minute)
	car3 := Vehicle{
		LicensePlate: "DEF456",
		Color:        "Green",
		Model:        "Compact",
	}
	_, err = attendant.AssignSpot(parkingLots, parkingSpotLists, &car3)
	if err != nil {
		t.Fatalf("Error parking car: %v", err)
	}

	// Find cars parked in the last 30 minutes
	recentCars := parkingService.FindCarsParkedLast30Mins()

	// Check if the correct number of cars is found
	if len(recentCars) != 2 {
		t.Errorf("Expected 2 cars parked in the last 30 minutes, got %d", len(recentCars))
	}

	// Check if the correct cars are found
	expectedLicensePlates := []string{"XYZ789", "DEF456"}
	for _, expectedLicensePlate := range expectedLicensePlates {
		found := false
		for _, car := range recentCars {
			if car.LicensePlate == expectedLicensePlate {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected to find car with license plate %s in recent cars, but it was not found", expectedLicensePlate)
		}
	}
}
