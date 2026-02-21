package handler

import (
	"bufio"
	"fmt"
	"lld-examples/internal/tictactoe/entity"
	"lld-examples/internal/tictactoe/service"
	"os"
	"strings"
)

type CLIHandler struct {
	gameService *service.GameService
}

func NewCLIHandler() *CLIHandler {
	return &CLIHandler{gameService: service.NewGameService()}
}

func (c *CLIHandler) Start() {
	fmt.Println("Welcome to Tic-Tac-Toe!")
	reader := bufio.NewReader(os.Stdin)

	for i := 1; i <= 2; i++ {
		for {
			fmt.Printf("Enter player %d (name,symbol): ", i)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			parts := strings.Split(input, ",")
			if len(parts) != 2 {
				fmt.Println("Invalid input. Format: name,symbol")
				continue
			}
			name := strings.TrimSpace(parts[0])
			symbol := strings.TrimSpace(parts[1])
			if name == "" || symbol == "" {
				fmt.Println("Name and symbol cannot be empty")
				continue
			}
			player := entity.NewPlayer(name, symbol)
			err := c.gameService.AddPlayer(player)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			break
		}
	}
	for {
		fmt.Println(c.gameService.RenderBoard())
		fmt.Print("Enter move (row,col): ")
		var row, col int
		fmt.Fscanf(reader, "%d,%d\n", &row, &col)
		err := c.gameService.MakeMove(row, col)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if c.gameService.CheckWinner() {
			fmt.Println("Game Over!")
			fmt.Println("Winner is -> ", c.gameService.GetWinner().GetName())
			break
		} else if c.gameService.IsBoardFull() {
			fmt.Println("Game Over! DRAW")
			break
		}
	}
}
