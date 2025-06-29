package main

import (
	"testing"
)

func TestNewParkingLot(t *testing.T) {
	capacity := 5
	pl := NewParkingLot(capacity)

	if pl.Capacity != capacity {
		t.Errorf("Expected capacity %d, got %d", capacity, pl.Capacity)
	}

	if len(pl.Slots) != capacity {
		t.Errorf("Expected %d slots, got %d", capacity, len(pl.Slots))
	}

	// Check all slots are initially empty
	for i, slot := range pl.Slots {
		if slot.IsOccupied {
			t.Errorf("Slot %d should be empty initially", i+1)
		}
		if slot.SlotNumber != i+1 {
			t.Errorf("Expected slot number %d, got %d", i+1, slot.SlotNumber)
		}
	}
}

func TestParkCar(t *testing.T) {
	pl := NewParkingLot(3)

	// Test parking first car
	slotNumber := pl.Park("CAR-1")
	if slotNumber != 1 {
		t.Errorf("Expected slot number 1, got %d", slotNumber)
	}

	// Test parking second car
	slotNumber = pl.Park("CAR-2")
	if slotNumber != 2 {
		t.Errorf("Expected slot number 2, got %d", slotNumber)
	}

	// Test parking third car
	slotNumber = pl.Park("CAR-3")
	if slotNumber != 3 {
		t.Errorf("Expected slot number 3, got %d", slotNumber)
	}

	// Test parking when lot is full
	slotNumber = pl.Park("CAR-4")
	if slotNumber != -1 {
		t.Errorf("Expected -1 for full parking lot, got %d", slotNumber)
	}
}

func TestLeaveCar(t *testing.T) {
	pl := NewParkingLot(3)

	// Park some cars first
	pl.Park("CAR-1")
	pl.Park("CAR-2")
	pl.Park("CAR-3")

	// Test leaving existing car
	slotNumber, charge, found := pl.Leave("CAR-2", 4)
	if !found {
		t.Error("Car should be found")
	}
	if slotNumber != 2 {
		t.Errorf("Expected slot number 2, got %d", slotNumber)
	}
	if charge != 30 {
		t.Errorf("Expected charge $30, got $%d", charge)
	}

	// Test leaving non-existing car
	_, _, found = pl.Leave("NON-EXISTING", 2)
	if found {
		t.Error("Non-existing car should not be found")
	}

	// Test parking in freed slot
	newSlotNumber := pl.Park("CAR-4")
	if newSlotNumber != 2 {
		t.Errorf("Expected to use freed slot 2, got %d", newSlotNumber)
	}
}

func TestChargeCalculation(t *testing.T) {
	pl := NewParkingLot(1)
	pl.Park("TEST-CAR")

	testCases := []struct {
		hours    int
		expected int
	}{
		{1, 10}, // First 2 hours = $10
		{2, 10}, // First 2 hours = $10
		{3, 20}, // $10 + $10 for 1 additional hour
		{4, 30}, // $10 + $20 for 2 additional hours
		{6, 50}, // $10 + $40 for 4 additional hours
	}

	for _, tc := range testCases {
		// Reset parking lot
		pl = NewParkingLot(1)
		pl.Park("TEST-CAR")

		_, charge, _ := pl.Leave("TEST-CAR", tc.hours)
		if charge != tc.expected {
			t.Errorf("For %d hours, expected charge $%d, got $%d",
				tc.hours, tc.expected, charge)
		}
	}
}

func TestParkingSlotAllocation(t *testing.T) {
	pl := NewParkingLot(5)

	// Park cars in slots 1, 2, 3
	pl.Park("CAR-1")
	pl.Park("CAR-2")
	pl.Park("CAR-3")

	// Remove car from slot 2
	pl.Leave("CAR-2", 2)

	// Next car should get slot 2 (nearest to entry)
	slotNumber := pl.Park("CAR-4")
	if slotNumber != 2 {
		t.Errorf("Expected nearest slot 2, got %d", slotNumber)
	}

	// Remove car from slot 1
	pl.Leave("CAR-1", 2)

	// Next car should get slot 1 (nearest to entry)
	slotNumber = pl.Park("CAR-5")
	if slotNumber != 1 {
		t.Errorf("Expected nearest slot 1, got %d", slotNumber)
	}
}

func TestEdgeCases(t *testing.T) {
	// Test with capacity 0
	pl := NewParkingLot(0)
	slotNumber := pl.Park("TEST-CAR")
	if slotNumber != -1 {
		t.Errorf("Expected -1 for zero capacity lot, got %d", slotNumber)
	}

	// Test with capacity 1
	pl = NewParkingLot(1)
	slotNumber = pl.Park("TEST-CAR")
	if slotNumber != 1 {
		t.Errorf("Expected slot 1, got %d", slotNumber)
	}

	// Second car should fail
	slotNumber = pl.Park("TEST-CAR-2")
	if slotNumber != -1 {
		t.Errorf("Expected -1 for full single slot lot, got %d", slotNumber)
	}
}

func TestCarRegistrationStorage(t *testing.T) {
	pl := NewParkingLot(2)

	registrationNumber := "CAR"
	slotNumber := pl.Park(registrationNumber)

	// Verify car is stored correctly
	slot := pl.Slots[slotNumber-1]
	if !slot.IsOccupied {
		t.Error("Slot should be occupied")
	}
	if slot.Car == nil {
		t.Error("Car should not be nil")
	}
	if slot.Car.RegistrationNumber != registrationNumber {
		t.Errorf("Expected registration %s, got %s",
			registrationNumber, slot.Car.RegistrationNumber)
	}
}

func TestMultipleOperations(t *testing.T) {
	pl := NewParkingLot(3)

	// Simulate the example scenario
	cars := []string{"CAR-1", "CAR-2", "CAR-3"}

	// Park all cars
	for i, car := range cars {
		slot := pl.Park(car)
		if slot != i+1 {
			t.Errorf("Expected slot %d for car %s, got %d", i+1, car, slot)
		}
	}

	// Try to park when full
	slot := pl.Park("EXTRA-CAR")
	if slot != -1 {
		t.Errorf("Expected -1 for full lot, got %d", slot)
	}

	// Leave one car
	slotNumber, charge, found := pl.Leave("CAR-2", 4)
	if !found || slotNumber != 2 || charge != 30 {
		t.Errorf("Leave operation failed: found=%v, slot=%d, charge=%d",
			found, slotNumber, charge)
	}

	// Park new car in freed slot
	newSlot := pl.Park("NEW-CAR")
	if newSlot != 2 {
		t.Errorf("Expected freed slot 2, got %d", newSlot)
	}
}

func TestProcessCommand_CreateParkingLot(t *testing.T) {
	var pl ParkingLot

	// Test valid create command
	ProcessCommand(&pl, "create_parking_lot 5")
	if pl.Capacity != 5 {
		t.Errorf("Expected capacity 5, got %d", pl.Capacity)
	}

	// Test invalid format - missing capacity
	ProcessCommand(&pl, "create_parking_lot")
	// Should not crash and maintain previous state

	// Test invalid capacity - non-numeric
	ProcessCommand(&pl, "create_parking_lot abc")
	// Should not crash and maintain previous state
}

func TestProcessCommand_Park(t *testing.T) {
	var pl ParkingLot

	// Initialize parking lot first
	ProcessCommand(&pl, "create_parking_lot 3")

	// Test valid park command
	ProcessCommand(&pl, "park CAR-1")
	if !pl.Slots[0].IsOccupied {
		t.Error("First slot should be occupied")
	}
	if pl.Slots[0].Car.RegistrationNumber != "CAR-1" {
		t.Errorf("Expected CAR-1, got %s", pl.Slots[0].Car.RegistrationNumber)
	}

	// Test invalid format - missing registration number
	ProcessCommand(&pl, "park")
	// Should not crash

	// Fill up the parking lot
	ProcessCommand(&pl, "park CAR-2")
	ProcessCommand(&pl, "park CAR-3")

	// Test parking when lot is full
	ProcessCommand(&pl, "park CAR-4")
	// Should handle gracefully (covered by output test)
}

func TestProcessCommand_Leave(t *testing.T) {
	var pl ParkingLot

	// Initialize and park some cars
	ProcessCommand(&pl, "create_parking_lot 3")
	ProcessCommand(&pl, "park CAR-1")
	ProcessCommand(&pl, "park CAR-2")

	// Test valid leave command
	ProcessCommand(&pl, "leave CAR-1 4")
	if pl.Slots[0].IsOccupied {
		t.Error("First slot should be free after leaving")
	}

	// Test leave non-existing car
	ProcessCommand(&pl, "leave NON-EXISTING 2")
	// Should handle gracefully

	// Test invalid format - missing hours
	ProcessCommand(&pl, "leave CAR-2")
	// Should not crash

	// Test invalid format - non-numeric hours
	ProcessCommand(&pl, "leave CAR-3 abc")
	// Should not crash
}

func TestProcessCommand_Status(t *testing.T) {
	var pl ParkingLot

	// Test status on empty lot
	ProcessCommand(&pl, "status")
	// Should not crash

	// Initialize and park some cars
	ProcessCommand(&pl, "create_parking_lot 3")
	ProcessCommand(&pl, "park CAR-1")
	ProcessCommand(&pl, "park CAR-2")

	// Test status with occupied slots
	ProcessCommand(&pl, "status")
	// Should display occupied slots (covered by output verification)
}

func TestProcessCommand_InvalidCommands(t *testing.T) {
	var pl ParkingLot

	// Test empty command
	ProcessCommand(&pl, "")
	// Should not crash

	// Test unknown command
	ProcessCommand(&pl, "unknown_command")
	// Should handle gracefully

	// Test command with extra spaces
	ProcessCommand(&pl, "  create_parking_lot   5  ")
	if pl.Capacity != 5 {
		t.Errorf("Expected capacity 5, got %d", pl.Capacity)
	}
}

func TestProcessCommand_CommandSequence(t *testing.T) {
	var pl ParkingLot

	// Test the complete example sequence
	commands := []string{
		"create_parking_lot 6",
		"park CAR-1",
		"park CAR-2",
		"park CAR-3",
		"park CAR-4",
		"park CAR-5",
		"park CAR-6",
		"leave CAR-6 4",
		"park CAR-7",
		"park CAR-8", // Should fail - lot full
		"leave CAR-1 4",
		"leave CAR-3 6",
		"park CAR-9",
		"park CAR-10",
		"park CAR-11", // Should fail - lot full
	}

	for _, cmd := range commands {
		ProcessCommand(&pl, cmd)
	}

	// Verify final state
	expectedOccupiedSlots := 6 // Should have 5 cars parked
	occupiedCount := 0
	for _, slot := range pl.Slots {
		if slot.IsOccupied {
			occupiedCount++
		}
	}

	if occupiedCount != expectedOccupiedSlots {
		t.Errorf("Expected %d occupied slots, got %d", expectedOccupiedSlots, occupiedCount)
	}

	// Verify specific cars are in expected positions
	expectedCars := map[int]string{
		1: "CAR-9",  // Slot 1 freed by CAR-1, taken by CAR-9
		2: "CAR-2",  // Never left
		3: "CAR-10", // Slot 3 freed by CAR-3, taken by CAR-10
		4: "CAR-4",  // Never left
		5: "CAR-5",  // Never left
		6: "CAR-7",  // Slot 6 freed by CAR-6, taken by CAR-7
	}

	for slotNum, expectedCar := range expectedCars {
		slot := pl.Slots[slotNum-1]
		if !slot.IsOccupied {
			t.Errorf("Slot %d should be occupied", slotNum)
			continue
		}
		if slot.Car.RegistrationNumber != expectedCar {
			t.Errorf("Slot %d: expected %s, got %s",
				slotNum, expectedCar, slot.Car.RegistrationNumber)
		}
	}
}

func TestProcessCommand_EdgeCases(t *testing.T) {
	var pl ParkingLot

	// Test operations before creating parking lot
	ProcessCommand(&pl, "park KA-01-HH-1234")
	ProcessCommand(&pl, "leave KA-01-HH-1234 2")
	ProcessCommand(&pl, "status")
	// Should handle gracefully without crashing

	// Test with zero capacity
	ProcessCommand(&pl, "create_parking_lot 0")
	ProcessCommand(&pl, "park KA-01-HH-1234")
	// Should handle gracefully

	// Test multiple create commands
	ProcessCommand(&pl, "create_parking_lot 3")
	ProcessCommand(&pl, "park KA-01-HH-1234")
	ProcessCommand(&pl, "create_parking_lot 5") // Recreate with different capacity
	if pl.Capacity != 5 {
		t.Errorf("Expected capacity 5 after recreation, got %d", pl.Capacity)
	}
	// Previous car should be gone
	if pl.Slots[0].IsOccupied {
		t.Error("Slots should be reset after recreating parking lot")
	}
}

func TestProcessCommand_WhitespaceHandling(t *testing.T) {
	var pl ParkingLot

	// Test commands with various whitespace
	testCases := []struct {
		command          string
		expectedCapacity int
	}{
		{"create_parking_lot 3", 3},
		{" create_parking_lot 4 ", 4},
		{"  create_parking_lot   5  ", 5},
		{"\tcreate_parking_lot\t6\t", 6},
	}

	for _, tc := range testCases {
		ProcessCommand(&pl, tc.command)
		if pl.Capacity != tc.expectedCapacity {
			t.Errorf("For command '%s', expected capacity %d, got %d",
				tc.command, tc.expectedCapacity, pl.Capacity)
		}
	}
}
