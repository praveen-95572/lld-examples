package service

import (
	"fmt"
	"lld-examples/internal/tictactoe/entity"
)

type GameService struct {
	board         entity.Board
	player1       entity.Player
	player2       entity.Player
	currentPlayer entity.Player

	movesMade int
	rowCount  map[entity.Player][]int
	colCount  map[entity.Player][]int
	diagCount map[entity.Player][2]int
	winner    entity.Player
}

func NewGameService() *GameService {
	return &GameService{
		board: entity.NewBoard(3)}
}

func (g *GameService) switchPlayer() {
	if g.currentPlayer == g.player1 {
		g.currentPlayer = g.player2
	} else {
		g.currentPlayer = g.player1
	}
}

func (g *GameService) AddPlayer(player entity.Player) error {
	if g.player1 == nil {
		g.player1 = player
		return nil
	} else if g.player2 == nil {
		if g.player1.GetName() == player.GetName() || g.player1.GetSymbol() == player.GetSymbol() {
			return fmt.Errorf("unique Name and symbol should be valid")
		}
		g.player2 = player
		g.currentPlayer = g.player1

		rowCount := map[entity.Player][]int{
			g.player1: make([]int, g.board.GetBoardSize()),
			g.player2: make([]int, g.board.GetBoardSize()),
		}
		colCount := map[entity.Player][]int{
			g.player1: make([]int, g.board.GetBoardSize()),
			g.player2: make([]int, g.board.GetBoardSize()),
		}
		diagCount := map[entity.Player][2]int{
			g.player1: {0, 0},
			g.player2: {0, 0},
		}
		g.rowCount = rowCount
		g.colCount = colCount
		g.diagCount = diagCount

		return nil
	}
	return fmt.Errorf("maximum players reached")
}

func (g *GameService) MakeMove(row, col int) error {
	if g.winner != nil {
		return nil
	}

	err := g.board.PlaceMark(row, col, g.currentPlayer.GetSymbol())
	if err != nil {
		return err
	}

	size := g.board.GetBoardSize()

	// Update counters
	g.rowCount[g.currentPlayer][row]++
	g.colCount[g.currentPlayer][col]++
	if row == col {
		diag := g.diagCount[g.currentPlayer]
		diag[0]++
		g.diagCount[g.currentPlayer] = diag
	}
	if row+col == size-1 {
		diag := g.diagCount[g.currentPlayer]
		diag[1]++
		g.diagCount[g.currentPlayer] = diag
	}

	// Check for winner
	if g.rowCount[g.currentPlayer][row] == size ||
		g.colCount[g.currentPlayer][col] == size ||
		g.diagCount[g.currentPlayer][0] == size ||
		g.diagCount[g.currentPlayer][1] == size {
		g.winner = g.currentPlayer
	}

	// Switch turn
	g.movesMade++
	g.switchPlayer()
	return nil
}

func (g *GameService) RenderBoard() string {
	return g.board.String()
}

func (g *GameService) CheckWinner() bool {
	return g.winner != nil
}

func (g *GameService) IsBoardFull() bool {
	return g.movesMade >= g.board.GetBoardSize()*g.board.GetBoardSize()
}

func (g *GameService) GetCurrentPlayer() entity.Player {
	return g.currentPlayer
}

func (g *GameService) GetWinner() entity.Player {
	return g.winner
}
