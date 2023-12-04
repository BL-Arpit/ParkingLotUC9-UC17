package main

import (
	"fmt"
	"os"
)

func main() {
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
	policeDepartment := NewPoliceDepartment(parkingService)
	parkingSpotLists := make([][][]ParkingSpot, len(parkingLots))

	for i := range parkingSpotLists {
		parkingSpotLists[i] = make([][]ParkingSpot, rows)
		for j := range parkingSpotLists[i] {
			parkingSpotLists[i][j] = make([]ParkingSpot, columns)
		}
	}

	var userType string
	fmt.Print("Are you a 'driver' or 'police'? ")
	fmt.Scan(&userType)

	switch userType {
	case "driver":
		driverMenu(parkingService, parkingSpotLists)
	case "police":
		policeMenu(policeDepartment)
	default:
		fmt.Println("Invalid user type. Exiting.")
		os.Exit(1)
	}
}

// driverMenu represents the menu for the driver.
func driverMenu(parkingService *ParkingService, parkingSpotLists [][][]ParkingSpot) {
	for {
		fmt.Println("1. Park a vehicle")
		fmt.Println("2. Unpark a vehicle")
		fmt.Println("3. Check parking lot status")
		fmt.Println("4. Find vehicle by License Plate")
		fmt.Println("5. Exit")

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

			var isHandicapped string
			fmt.Print("Is the driver handicapped? (yes/no): ")
			fmt.Scan(&isHandicapped)
			vehicle.Handicapped = isHandicapped == "yes" || isHandicapped == "y"

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
			var licensePlate string
			fmt.Print("Enter License Plate of the vehicle to find: ")
			fmt.Scan(&licensePlate)

			result, err := parkingService.FindByLicensePlate(licensePlate)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(result)
			}

		case 5:
			os.Exit(0)

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// policeMenu represents the menu for the police.
func policeMenu(policeDepartment *PoliceDepartment) {
	for {
		fmt.Println("1. Display Parking Lot Status")
		fmt.Println("2. Find and Display White Cars")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			policeDepartment.DisplayParkingLotStatus()
		case 2:
			policeDepartment.FindAndDisplayWhiteCars()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
