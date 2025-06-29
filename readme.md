# Parking Lot Manager

A simple parking lot management system built with Go for automated parking operations.

## Features

- Create parking lot with specified capacity
- Park cars with nearest-to-entry slot allocation
- Remove cars with automatic parking charge calculation
- Display status of all occupied parking slots

## Data Structures

- **Car**: Stores vehicle registration number
- **ParkingSlot**: Represents parking slot with number and occupancy status
- **ParkingLot**: Main parking system with capacity and slot array

## Commands

1. `create_parking_lot {capacity}` - Create parking lot with capacity n
2. `park {car_number}` - Park a car with registration number
3. `leave {car_number} {hours}` - Remove car with parking duration (in hours)
4. `status` - Display status of all occupied slots

## Pricing Model

- $10 for the first 2 hours
- $10 for each additional hour

Examples:
- 2 hours = $10
- 4 hours = $30 ($10 + $20)
- 6 hours = $50 ($10 + $40)

## How to Run

1. Save the code as `main.go`
2. Create input file with commands (example: `input.txt`)
3. Run the program:
   ```bash
   go run main.go input.txt
   ```

## Running Tests

Run tests with coverage:
```bash
make test
```

## Sample Input

```
create_parking_lot 6
park KA-01-HH-1234
park KA-01-HH-9999
park KA-01-BB-0001
park KA-01-HH-7777
park KA-01-HH-2701
park KA-01-HH-3141
leave KA-01-HH-3141 4
status
park KA-01-P-333
park DL-12-AA-9999
leave KA-01-HH-1234 4
leave KA-01-BB-0001 6
leave DL-12-AA-9999 2
park KA-09-HH-0987
park CA-09-IO-1111
park KA-09-HH-0123
status
```

## Sample Output

```
Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 4
Allocated slot number: 5
Allocated slot number: 6
Registration number KA-01-HH-3141 with Slot Number 6 is free with Charge $30
Slot No.	Registration No.
1		KA-01-HH-1234
2		KA-01-HH-9999
3		KA-01-BB-0001
4		KA-01-HH-7777
5		KA-01-HH-2701
Allocated slot number: 6
Sorry, parking lot is full
Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30
Registration number KA-01-BB-0001 with Slot Number 3 is free with Charge $50
Registration number DL-12-AA-9999 not found
Allocated slot number: 1
Allocated slot number: 3
Sorry, parking lot is full
Slot No.	Registration No.
1		KA-09-HH-0987
2		KA-01-HH-9999
3		CA-09-IO-1111
4		KA-01-HH-7777
5		KA-01-HH-2701
6		KA-01-P-333
```

## Project Structure

```
parking-lot-manager/
├── main.go          # Main application code
├── parking.go       # Handler function code for each command
├── parking_test.go  # Unit tests
├── input.txt        # Sample input file
└── README.md        # Documentation
```

## Algorithm

1. **Parking**: Finds the lowest numbered available slot (nearest to entry)
2. **Leaving**: Searches for car by registration number and calculates charges
3. **Status**: Displays all occupied slots in sequential order

## Test Coverage

The application includes comprehensive unit tests covering:
- Parking lot creation and initialization
- Car parking with slot allocation
- Car removal with charge calculation
- Edge cases (full lot, non-existing cars)
- Slot reallocation after cars leave
- Pricing calculation for various durations

## Technical Details

- **Language**: Go 1.21+
- **Testing**: Built-in Go testing framework