package main

import (
	"fmt"
	"lld-examples/internal/snakeladder/entity"
	"lld-examples/internal/snakeladder/service"
)

func main() {
	// Define snakes and ladders
	snakes := []entity.Snake{
		{Head: 16, Tail: 6},
		{Head: 48, Tail: 26},
		{Head: 49, Tail: 11},
		{Head: 56, Tail: 53},
		{Head: 62, Tail: 19},
		{Head: 64, Tail: 60},
		{Head: 87, Tail: 24},
		{Head: 93, Tail: 73},
		{Head: 95, Tail: 75},
		{Head: 98, Tail: 78},
	}

	ladders := []entity.Ladder{
		{Bottom: 1, Top: 38},
		{Bottom: 4, Top: 14},
		{Bottom: 9, Top: 31},
		{Bottom: 21, Top: 42},
		{Bottom: 28, Top: 84},
		{Bottom: 36, Top: 44},
		{Bottom: 51, Top: 67},
		{Bottom: 71, Top: 91},
		{Bottom: 80, Top: 100},
	}

	board := entity.NewBoard(100, snakes, ladders)

	players := []*entity.Player{
		entity.NewPlayer("Nobita"),
		entity.NewPlayer("Gian"),
		entity.NewPlayer("Shizuka"),
		entity.NewPlayer("Doraemon"),
	}

	dice := entity.NewDice(6)

	game := service.NewGame(players, board, dice)

	fmt.Println("Starting Snake and Ladder Game!")

	for !game.IsOver() {
		msg := game.PlayTurn()
		fmt.Println(msg)
	}
}
