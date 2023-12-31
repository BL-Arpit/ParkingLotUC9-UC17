package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// ParkingAttendant represents the attendant responsible for assigning parking spots.
type ParkingAttendant struct{}

// NewParkingAttendant creates a new parking attendant.
func NewParkingAttendant() *ParkingAttendant {
	return &ParkingAttendant{}
}

// AssignSpot assigns a parking spot to a vehicle gradually, filling all parking lots evenly in a round-robin fashion.
func (pa *ParkingAttendant) AssignSpot(parkingLots []*ParkingLot, parkingSpotLists [][][]ParkingSpot, vehicle *Vehicle) (string, error) {
	// Check if the vehicle is handicapped
	if vehicle.Handicapped {
		return pa.assignHandicappedSpot(parkingLots, parkingSpotLists, vehicle)
	} else if strings.ToLower(vehicle.Model) == "suv" {
		// If the vehicle is an SUV, find the parking lot with the highest available space
		highestAvailableLotIndex := -1
		highestAvailableSpace := 0

		for i, lot := range parkingLots {
			if lot.AvailableSpaces > highestAvailableSpace {
				highestAvailableLotIndex = i
				highestAvailableSpace = lot.AvailableSpaces
			}
		}

		if highestAvailableLotIndex == -1 {
			return "", fmt.Errorf("all parking lots are full")
		}

		// Assign a parking spot in the parking lot with the highest available space
		for i := 0; i < len(parkingSpotLists[highestAvailableLotIndex]); i++ {
			if !parkingSpotLists[highestAvailableLotIndex][i][0].Occupied {
				parkingSpot := string(rune('a'+i)) + fmt.Sprintf("%d", highestAvailableLotIndex+1)
				parkingSpotLists[highestAvailableLotIndex][i][0].Occupied = true
				vehicle.ParkingSpot = parkingSpot
				parkingLots[highestAvailableLotIndex].ParkedVehicles = append(parkingLots[highestAvailableLotIndex].ParkedVehicles, *vehicle)
				parkingLots[highestAvailableLotIndex].AvailableSpaces--
				return parkingSpot, nil
			}
		}
	} else {

		// Iterate over parking spots in a round-robin fashion
		for i := 0; i < len(parkingSpotLists[0]); i++ {
			// Iterate over parking lots in a round-robin fashion
			for j := 0; j < len(parkingLots); j++ {
				// Find the next available spot in the current parking lot
				if !parkingSpotLists[j][i][0].Occupied {
					// Convert row and column indices to parking spot identifier (e.g., a1, b2, etc.)
					parkingSpot := string(rune('a'+i)) + fmt.Sprintf("%d", j+1)
					parkingSpotLists[j][i][0].Occupied = true
					vehicle.ParkingSpot = parkingSpot
					parkingLots[j].ParkedVehicles = append(parkingLots[j].ParkedVehicles, *vehicle)
					parkingLots[j].AvailableSpaces--
					return parkingSpot, nil
				}
			}
		}
	}

	return "", fmt.Errorf("parking lots are full")
}

func (pa *ParkingAttendant) assignHandicappedSpot(parkingLots []*ParkingLot, parkingSpotLists [][][]ParkingSpot, vehicle *Vehicle) (string, error) {
	// Iterate over parking lots and find the nearest available spot
	for _, lot := range parkingLots {
		for i := 0; i < len(parkingSpotLists[0]); i++ {
			if !parkingSpotLists[lot.ID-1][i][0].Occupied {
				// Convert row and column indices to parking spot identifier (e.g., a1, b2, etc.)
				parkingSpot := string(rune('a'+i)) + fmt.Sprintf("%d", lot.ID)
				parkingSpotLists[lot.ID-1][i][0].Occupied = true
				vehicle.ParkingSpot = parkingSpot
				lot.ParkedVehicles = append(lot.ParkedVehicles, *vehicle)
				lot.AvailableSpaces--
				return parkingSpot, nil
			}
		}
	}

	return "", fmt.Errorf("handicapped parking is full")
}

// getRandomAvailableSpot returns an available parking spot randomly in the given parking lot.
func getRandomAvailableSpot(lotID int, parkingSpotLists [][][]ParkingSpot, rows, columns int) (string, error) {
	// Create a list of all possible parking spots
	allSpots := make([]string, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			parkingSpot := string(rune('a'+i)) + fmt.Sprintf("%d", j+1)
			allSpots = append(allSpots, parkingSpot)
		}
	}

	// Remove already occupied spots in the specified parking lot
	for i, row := range parkingSpotLists[lotID] {
		for j, spot := range row {
			if spot.Occupied {
				occupiedSpot := string(rune('a'+i)) + fmt.Sprintf("%d", j+1)
				for k, availableSpot := range allSpots {
					if availableSpot == occupiedSpot {
						allSpots = append(allSpots[:k], allSpots[k+1:]...)
						break
					}
				}
			}
		}
	}

	if len(allSpots) == 0 {
		return "", fmt.Errorf("no available spot")
	}

	// Choose a random available spot
	randomIndex := rand.Intn(len(allSpots))
	return allSpots[randomIndex], nil
}
func (pa *ParkingAttendant) IsLastSpace(parkingSpotLists [][][]ParkingSpot, lastSpot string) bool {
	for _, row := range parkingSpotLists {
		for _, spot := range row {
			for _, parkingSpot := range spot {
				if parkingSpot.ParkingSpot == lastSpot && !parkingSpot.Occupied {
					return true
				}
			}
		}
	}
	return false
}
