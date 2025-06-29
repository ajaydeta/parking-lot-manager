package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Car represents a car with registration number
type Car struct {
	RegistrationNumber string
}

// ParkingSlot represents a parking slot
type ParkingSlot struct {
	SlotNumber int
	Car        *Car
	IsOccupied bool
}

// ParkingLot represents the parking lot system
type ParkingLot struct {
	Capacity int
	Slots    []*ParkingSlot
}

// NewParkingLot creates a new parking lot with given capacity
func NewParkingLot(capacity int) *ParkingLot {
	slots := make([]*ParkingSlot, capacity)
	for i := 0; i < capacity; i++ {
		slots[i] = &ParkingSlot{
			SlotNumber: i + 1,
			Car:        nil,
			IsOccupied: false,
		}
	}
	return &ParkingLot{
		Capacity: capacity,
		Slots:    slots,
	}
}

// Park parks a car and returns the allocated slot number
func (pl *ParkingLot) Park(registrationNumber string) int {
	// Find the nearest available slot
	for i := 0; i < pl.Capacity; i++ {
		if !pl.Slots[i].IsOccupied {
			pl.Slots[i].Car = &Car{RegistrationNumber: registrationNumber}
			pl.Slots[i].IsOccupied = true
			return pl.Slots[i].SlotNumber
		}
	}
	return -1 // Parking lot is full
}

// Leave removes a car from parking lot and calculates charge
func (pl *ParkingLot) Leave(registrationNumber string, hours int) (int, int, bool) {
	for i := 0; i < pl.Capacity; i++ {
		if pl.Slots[i].IsOccupied && pl.Slots[i].Car.RegistrationNumber == registrationNumber {
			slotNumber := pl.Slots[i].SlotNumber
			pl.Slots[i].Car = nil
			pl.Slots[i].IsOccupied = false

			// Calculate charge: $10 for first 2 hours, $10 for every additional hour
			charge := 10 // Base charge for first 2 hours
			if hours > 2 {
				charge += (hours - 2) * 10
			}

			return slotNumber, charge, true
		}
	}
	return -1, 0, false // Car not found
}

// Status prints the current status of all occupied slots
func (pl *ParkingLot) Status() {
	fmt.Println("Slot No.\tRegistration No.")
	for i := 0; i < pl.Capacity; i++ {
		if pl.Slots[i].IsOccupied {
			fmt.Printf("%d\t\t%s\n", pl.Slots[i].SlotNumber, pl.Slots[i].Car.RegistrationNumber)
		}
	}
}

// ProcessCommand processes a single command
func ProcessCommand(pl *ParkingLot, command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "create_parking_lot":
		if len(parts) != 2 {
			fmt.Println("Invalid command format")
			return
		}
		capacity, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Invalid capacity")
			return
		}
		*pl = *NewParkingLot(capacity)
		fmt.Printf("Created a parking lot with %d slots\n", capacity)

	case "park":
		if len(parts) != 2 {
			fmt.Println("Invalid command format")
			return
		}
		registrationNumber := parts[1]
		slotNumber := pl.Park(registrationNumber)
		if slotNumber == -1 {
			fmt.Println("Sorry, parking lot is full")
		} else {
			fmt.Printf("Allocated slot number: %d\n", slotNumber)
		}

	case "leave":
		if len(parts) != 3 {
			fmt.Println("Invalid command format")
			return
		}
		registrationNumber := parts[1]
		hours, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Invalid hours")
			return
		}
		slotNumber, charge, found := pl.Leave(registrationNumber, hours)
		if found {
			fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n",
				registrationNumber, slotNumber, charge)
		} else {
			fmt.Printf("Registration number %s not found\n", registrationNumber)
		}

	case "status":
		pl.Status()

	default:
		fmt.Println("Unknown command")
	}
}
