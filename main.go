package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed for random number generation

	rows := 6
	columns := 6

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

	for {
		fmt.Println("1. Park a vehicle")
		fmt.Println("2. Unpark a vehicle")
		fmt.Println("3. Check parking lot status")
		fmt.Println("4. Exit")
		fmt.Println("5. Find by License Plate")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var vehicle Vehicle
			fmt.Print("Enter License Plate: ")
			fmt.Scan(&vehicle.LicensePlate)
			fmt.Print("Enter Color: ")
			fmt.Scan(&vehicle.Color)
			fmt.Print("Enter Model: ")
			fmt.Scan(&vehicle.Model)
			fmt.Print("Is the driver handicapped? (true/false): ")
			fmt.Scan(&vehicle.Handicapped)

			err := parkingService.Park(vehicle, parkingSpotLists)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Vehicle parked successfully.")
			}

		case 2:
			var licensePlate string
			fmt.Print("Enter License Plate of the vehicle to unpark: ")
			fmt.Scan(&licensePlate)

			err := parkingService.Unpark(licensePlate, parkingSpotLists)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Vehicle unparked successfully.")
			}

		case 3:
			parkingService.Status()

		case 4:
			os.Exit(0)

		case 5:
			var licensePlateToFind string
			fmt.Print("Enter License Plate to find the parked car: ")
			fmt.Scan(&licensePlateToFind)

			result, err := parkingService.FindByLicensePlate(licensePlateToFind)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(result)
			}

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
