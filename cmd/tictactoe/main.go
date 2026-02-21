package main

import (
	"lld-examples/internal/tictactoe/handler"
)

func main() {
	cli := handler.NewCLIHandler()
	cli.Start()
}
