package main

// ParkingSpot represents a parking spot in the parking lot.
type ParkingSpot struct {
	ParkingSpot string
	Occupied    bool
}

// getRandomAvailableSpot returns an available parking spot randomly.
//func getRandomAvailableSpot(parkingSpotList [][]ParkingSpot) (string, error) {
//	availableSpots := make([]string, 0)
//	for i, row := range parkingSpotList {
//		for j, spot := range row {
//			if !spot.Occupied {
//				// Convert row and column indices to parking spot identifier (e.g., a1, b2, etc.)
//				parkingSpot := string(rune('a'+i)) + fmt.Sprintf("%d", j+1)
//				availableSpots = append(availableSpots, parkingSpot)
//			}
//		}
//	}
//
//	if len(availableSpots) == 0 {
//		return "", fmt.Errorf("no available spot")
//	}
//
//	// Choose a random available spot
//	randomIndex := rand.Intn(len(availableSpots))
//	return availableSpots[randomIndex], nil
//}
