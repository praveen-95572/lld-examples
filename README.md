# LLD Examples

This repository contains **Low-Level Design (LLD) examples** implemented in Go. It is intended to demonstrate clean architecture, SOLID principles, and best practices for designing systems and small projects.

## Purpose

- Provide **reference implementations** for LLD problems.
- Showcase **Go project structure** for CMD, internal, and pkg.
- Demonstrate **domain-driven design, hexagonal architecture**, and other design patterns.
- Serve as a learning resource for **interview preparation and system design practice**.

## Project Structure

```
lld-examples/
│
├── cmd/ # Entry points for different examples
│ └── main.go # CLI or main application for a specific example
│
├── internal/ # Core business logic
│ ├── tictactoe/ # Tic-Tac-Toe game example
│ ├── parking/ # Parking Lot design example
│ └── ... # Other LLD examples
│
├── pkg/utils/ # Reusable utility packages
│
├── go.mod # Go module file
├── go.sum
└── README.md # Project documentation
```


## Examples Included

- **Tic-Tac-Toe**: Turn-based game demonstrating interfaces, services, board management, and winner tracking.
- **Parking Lot**: Object-oriented design with vehicles, slots, and services, demonstrating extensibility and O(1) operations.
- **Others**: LRU Cache, Snake and Ladder, and additional practice examples.

## How to Run

1. Clone the repository:

```bash
git clone https://github.com/praveen-95572/lld-examples.git
cd lld-examples
```
