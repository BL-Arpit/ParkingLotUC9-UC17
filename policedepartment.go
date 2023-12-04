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

// CheckForBMW checks if there are any BMW brand cars parked and tightens security if found.
func (pd *PoliceDepartment) CheckForBMW() {
	bmwCars, err := pd.parkingService.FindByBrand("BMW")
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(bmwCars) > 0 {
		pd.parkingService.securityStaff.NotifyHighSecurity()
		fmt.Println("\nBMW Cars found. Security tightened.")
	} else {
		fmt.Println("\nNo BMW Cars found. No change in security.")
	}
}

func (pd *PoliceDepartment) DisplayRecentCars() {
	recentCars := pd.parkingService.FindCarsParkedLast30Mins()

	if len(recentCars) > 0 {
		fmt.Println("Cars parked in the last 30 minutes:")
		for _, car := range recentCars {
			fmt.Printf("License Plate: %s, Color: %s, Model: %s, Parking Spot: %s\n",
				car.LicensePlate, car.Color, car.Model, car.ParkingSpot)
		}
	} else {
		fmt.Println("No cars found that were parked in the last 30 minutes.")
	}
}
func (pd *PoliceDepartment) CheckIfNumberInHandicappedVehicles(userProvidedNumber string) bool {
	handicappedVehicles := pd.parkingService.FindHandicappedVehiclesAtRowsBAndD()

	for _, vehicle := range handicappedVehicles {
		// Extract the number from the license plate (assuming it's a single-digit number)
		vehicleNumber := string(vehicle.LicensePlate[len(vehicle.LicensePlate)-1])

		if vehicleNumber == userProvidedNumber {
			return true
		}
	}

	return false
}

func (pd *PoliceDepartment) CheckHandicapped() {
	handicappedVehicles := pd.parkingService.FindHandicappedVehiclesAtRowsBAndD()

	if len(handicappedVehicles) == 0 {
		fmt.Println("No handicapped vehicles found in Rows B and D.")
		return
	}

	fmt.Println("Handicapped vehicles found in Rows B and D:")
	for _, vehicle := range handicappedVehicles {
		fmt.Printf("License Plate: %s, Color: %s, Model: %s, Parking Spot: %s\n",
			vehicle.LicensePlate, vehicle.Color, vehicle.Model, vehicle.ParkingSpot)
	}
}
