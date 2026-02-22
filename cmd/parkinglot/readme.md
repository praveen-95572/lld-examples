# Parking Lot System – Low Level Design (LLD)

## 1. Problem Statement

Design a Parking Lot System that:

- Supports multiple floors
- Supports different vehicle types (Car, Bike, Truck)
- Allocates nearest available parking spot
- Generates ticket on entry
- Calculates parking fee on exit
- Frees spot after exit

---

## 2. Functional Requirements

- Vehicle enters parking lot → spot allocated → ticket generated
- Vehicle exits → fee calculated → spot freed
- Support multiple vehicle types
- Prevent parking if lot is full
- Track spot availability

---

## 3. Non-Functional Requirements

- Modular and extensible
- Easy to add more vehicle types
- Thread-safe (extendable)
- O(1) spot lookup

---

## 4. Core Entities

### 1. Vehicle
- License Number
- Vehicle Type

### 2. ParkingSpot
- Spot ID
- Spot Type
- IsOccupied
- Parked Vehicle

### 3. ParkingFloor
- Floor number
- List of parking spots

### 4. Ticket
- Ticket ID
- Vehicle
- Spot
- Entry Time

### 5. ParkingLot
- Floors
- Active Tickets

---

## 5. Flow

### Vehicle Entry Flow

1. Vehicle arrives
2. System finds nearest available spot
3. Spot is marked occupied
4. Ticket generated
5. Ticket returned to user

### Vehicle Exit Flow

1. Ticket scanned
2. Fee calculated
3. Spot freed
4. Ticket removed from active list

---

## 6. Design Principles

### Single Responsibility Principle
- Vehicle → only vehicle data
- Spot → parking state
- Ticket → transaction data
- ParkingLot → orchestrator

### Extensibility
- Add more vehicle types
- Add hourly pricing strategy
- Add multiple entry gates

---

## 7. Future Enhancements

- Add payment integration
- Add concurrency control (mutex)
- Add database persistence
- Add reservation system
- Add dynamic pricing
