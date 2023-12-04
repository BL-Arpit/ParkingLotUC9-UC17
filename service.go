package main

import (
	"fmt"
	"time"
)

// ParkingService represents the service for managing parking operations.
type ParkingService struct {
	parkingLots   []*ParkingLot
	attendant     *ParkingAttendant
	securityStaff *SecurityStaff
	parkingSpots  []string
}

// NewParkingService creates a new parking service with the given parking lots, attendant, and security staff.
func NewParkingService(parkingLots []*ParkingLot, attendant *ParkingAttendant, securityStaff *SecurityStaff) *ParkingService {
	return &ParkingService{
		parkingLots:   parkingLots,
		attendant:     attendant,
		securityStaff: securityStaff,
	}
}

// Park parks a vehicle in the parking lot.
func (s *ParkingService) Park(vehicle Vehicle, parkingSpotLists [][][]ParkingSpot) error {
	// Check if all parking lots are full
	if s.AllLotsFull() {
		return fmt.Errorf("all parking lots are full")
	}

	// Assign a parking spot using the attendant
	lastSpot, err := s.attendant.AssignSpot(s.parkingLots, parkingSpotLists, &vehicle)
	if err != nil {
		return err
	}

	// Record the parking time
	vehicle.ParkingTime = time.Now()

	// Check if parking lot is full and notify only when the last space is occupied
	if s.attendant.IsLastSpace(parkingSpotLists, lastSpot) {
		s.NotifyFull()
		s.securityStaff.Notify()
	}

	return nil
}

// AllLotsFull checks if all parking lots are full.
func (s *ParkingService) AllLotsFull() bool {
	for _, parkingLot := range s.parkingLots {
		if parkingLot.AvailableSpaces > 0 {
			return false
		}
	}
	return true
}

// Unpark removes a parked vehicle from one of the parking lots.
func (s *ParkingService) Unpark(licensePlate string, parkingSpotLists [][][]ParkingSpot) error {
	// Check if the parking lot was full before the unpark operation
	wasFullBeforeUnpark := false
	for _, parkingLot := range s.parkingLots {
		if parkingLot.AvailableSpaces == 0 {
			wasFullBeforeUnpark = true
			break
		}
	}

	for _, parkingLot := range s.parkingLots {
		for i, vehicle := range parkingLot.ParkedVehicles {
			if vehicle.LicensePlate == licensePlate {
				// Remove the vehicle from the slice
				parkingLot.ParkedVehicles = append(parkingLot.ParkedVehicles[:i], parkingLot.ParkedVehicles[i+1:]...)
				parkingLot.AvailableSpaces++

				// Mark the corresponding parking spot as unoccupied
				row := int(vehicle.ParkingSpot[0] - 'a')
				col := int(vehicle.ParkingSpot[1] - '0' - 1)

				// Assume the third dimension in parkingSpotLists is for parking lots
				for _, parkingSpotList := range parkingSpotLists {
					if row < len(parkingSpotList) && col < len(parkingSpotList[row]) {
						parkingSpotList[row][col].Occupied = false // Ensure correct casing
					}
				}

				// Check if parking lot is now available and notify
				if wasFullBeforeUnpark && parkingLot.AvailableSpaces == 1 {
					s.NotifyAvailable()
				}

				return nil
			}
		}
	}

	return fmt.Errorf("vehicle with license plate %s is not parked", licensePlate)
}

// NotifyFull is a notification function when any of the parking lots is full.
func (s *ParkingService) NotifyFull() {
	fmt.Println("Parking lot is full. No more vehicles can be parked.")
}

// NotifyAvailable is a notification function when any of the parking lots is available again.
func (s *ParkingService) NotifyAvailable() {
	fmt.Println("Parking lot is available. You can now park vehicles.")
}

// Status prints the current status of all parking lots, including parked vehicles.
func (s *ParkingService) Status() {
	for _, parkingLot := range s.parkingLots {
		parkingLot.Status()
		fmt.Println()
	}
}
func (s *ParkingService) FindByLicensePlate(licensePlate string) (string, error) {
	for _, parkingLot := range s.parkingLots {
		for _, vehicle := range parkingLot.ParkedVehicles {
			if vehicle.LicensePlate == licensePlate {
				return fmt.Sprintf("Vehicle with License Plate %s is parked at spot %s in Parking Lot %d",
					licensePlate, vehicle.ParkingSpot, parkingLot.ID), nil
			}
		}
	}
	return "", fmt.Errorf("vehicle with license plate %s not found", licensePlate)
}

func (s *ParkingService) FindAllWhiteCars() []Vehicle {
	whiteCars := make([]Vehicle, 0)

	for _, parkingLot := range s.parkingLots {
		for _, vehicle := range parkingLot.ParkedVehicles {
			if vehicle.Color == "White" {
				whiteCars = append(whiteCars, vehicle)
			}
		}
	}

	return whiteCars
}

func (s *ParkingService) FindByColorAndModel(color, model string) (string, error) {
	for _, parkingLot := range s.parkingLots {
		for _, vehicle := range parkingLot.ParkedVehicles {
			if vehicle.Color == color && vehicle.Model == model {
				return fmt.Sprintf("Vehicle with Color %s and Model %s is parked at spot %s in Parking Lot %d",
					color, model, vehicle.ParkingSpot, parkingLot.ID), nil
			}
		}
	}
	return "", fmt.Errorf("vehicle with color %s and model %s not found", color, model)
}
func (s *ParkingService) FindByBrand(brand string) ([]Vehicle, error) {
	foundVehicles := make([]Vehicle, 0)

	for _, parkingLot := range s.parkingLots {
		for _, vehicle := range parkingLot.ParkedVehicles {
			if vehicle.Brand == brand {
				foundVehicles = append(foundVehicles, vehicle)
			}
		}
	}

	if len(foundVehicles) == 0 {
		return nil, fmt.Errorf("no vehicles found with brand: %s", brand)
	}

	return foundVehicles, nil
}

func (s *ParkingService) FindCarsParkedLast30Mins() []Vehicle {
	var recentCars []Vehicle
	currentTime := time.Now()

	for _, parkingLot := range s.parkingLots {
		for _, vehicle := range parkingLot.ParkedVehicles {
			elapsedTime := currentTime.Sub(vehicle.ParkingTime)
			if elapsedTime <= 30*time.Minute {
				recentCars = append(recentCars, vehicle)
			}
		}
	}

	return recentCars
}

func (s *ParkingService) FindHandicappedVehiclesAtRowsBAndD() []Vehicle {
	handicappedVehicles := make([]Vehicle, 0)

	for _, parkingLot := range s.parkingLots {
		for row := 1; row <= parkingLot.Rows; row++ {
			if row == 2 || row == 4 { // Rows B and D
				for col := 1; col <= parkingLot.Columns; col++ {
					// Iterate over all vehicles in the parking lot
					for _, vehicle := range parkingLot.ParkedVehicles {
						// Check if the vehicle is in the specified row, col, and handicapped
						if vehicle.Handicapped && string(rune('A'+row-1)) == string(vehicle.ParkingSpot[0]) && col == int(vehicle.ParkingSpot[1]-'0') {
							handicappedVehicles = append(handicappedVehicles, vehicle)
						}
					}
				}
			}
		}
	}

	return handicappedVehicles
}
