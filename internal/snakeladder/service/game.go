package service

import (
	"fmt"
	"lld-examples/internal/snakeladder/entity"
)

type Game struct {
	Board   *entity.Board
	Players []*entity.Player
	Dice    *entity.Dice
	Turn    int
	Winner  *entity.Player
}

func NewGame(players []*entity.Player, board *entity.Board, dice *entity.Dice) *Game {
	return &Game{
		Board:   board,
		Players: players,
		Dice:    dice,
		Turn:    0,
		Winner:  nil,
	}
}

// Play one turn
func (g *Game) PlayTurn() string {
	if g.Winner != nil {
		return fmt.Sprintf("Game over! Winner is %s", g.Winner.Name)
	}

	player := g.Players[g.Turn]
	roll := g.Dice.Roll()
	nextPos := player.Position + roll

	if nextPos > g.Board.Size {
		g.NextTurn()
		return fmt.Sprintf("%s rolled %d and stays at %d", player.Name, roll, player.Position)
	}

	player.Position = g.Board.GetNextPosition(nextPos)

	message := fmt.Sprintf("%s rolled %d and moved to %d", player.Name, roll, player.Position)

	if player.Position == g.Board.Size {
		g.Winner = player
		message += fmt.Sprintf(" -- %s wins!", player.Name)
	}

	g.NextTurn()
	return message
}

func (g *Game) NextTurn() {
	g.Turn = (g.Turn + 1) % len(g.Players)
}

// Check if game is over
func (g *Game) IsOver() bool {
	return g.Winner != nil
}
