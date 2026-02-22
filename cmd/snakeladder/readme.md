# Snake and Ladder Game in Golang (LLD Design)

## 1. Requirements

### Functional Requirements
- 2+ players play the game on a board of 100 cells (1 to 100).  
- Players take turns rolling a dice (1–6) and move accordingly.  
- If a player lands on the bottom of a ladder, they climb to the top.  
- If a player lands on the head of a snake, they slide down to the tail.  
- The first player to reach exactly 100 wins.  
- If the dice roll exceeds 100, the player doesn’t move.

### Non-Functional Requirements
- Modular, maintainable code.  
- Easy to extend (board size, snakes/ladders, dice variations).  

---

## 2. Core Entities

- **Player** – stores player name and current position.  
- **Board** – contains cells and mappings for snakes and ladders.  
- **Snake** – has head (start) and tail (end) positions.  
- **Ladder** – has bottom (start) and top (end) positions.  
- **Dice** – rolls a number between 1 and 6.  
- **Game** – controls game flow, players’ turns, dice rolls, and win conditions.  

---

## 3. Golang Structs

```go
type Player struct {
    Name     string
    Position int
}

type Snake struct {
    Head int
    Tail int
}

type Ladder struct {
    Bottom int
    Top    int
}

type Dice struct{}

type Board struct {
    Size    int
    Snakes  map[int]int // head -> tail
    Ladders map[int]int // bottom -> top
}

type Game struct {
    Board   *Board
    Players []*Player
    Dice    *Dice
    Turn    int
}
```

## Design
### Single Responsibility Principle (SRP)
- Each struct handles its own responsibility.
### Extensibility
- Board size, number of players, snakes/ladders, and dice rules can be modified easily.
### Separation of Concerns
- Dice rolling, board movement, and game flow are independent modules.
### Error Handling
- Player cannot move past the final cell; dice roll exceeding 100 is ignored.