# Tic Tac Toe Game in Golang (LLD Design)

## 1. Requirements

### Functional Requirements
- Two players (X and O) play on a 3x3 board.
- Players take turns placing their mark.
- The game ends with a win or a draw.
- Must handle invalid moves (cell already occupied or out-of-bounds).

### Non-Functional Requirements
- Easy to extend to an NxN board.
- Clear, modular, and maintainable code.

---

## 2. Core Entities

From a low-level design (LLD) perspective, the main entities are:

- **Player** – stores a name and a symbol (X or O).  
- **Board** – stores the board state and manages move validations.  
- **Game** – manages game flow, turn switching, and win/draw checks.  
- **Cell** – optional; part of the board to store cell values (empty, X, or O).  

---

## 3. Golang Structs

## 4. Board Methods

## 5. Game Methods

## 6. Main Function

## 7. Design 

### SRP
- Player → stores player info.
- Board → manages board state and validations.
- Game → controls game flow and turns.

### Extensibility
-Board size can be NxN.
-Easy to extend for AI, network multiplayer, or GUI.

### Error Handling
- Invalid moves are handled in Place() method.