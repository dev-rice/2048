package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/donutmonger/2048/board"
)

func main() {
	gameBoard := board.NewEmptyBoard()
	gameBoard[0][3] = 2
	gameBoard[0][1] = 2

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "q" { // break the loop if text == "q"
		fmt.Println("")
		fmt.Println(board.BoardToString(gameBoard))
		fmt.Println("")

		fmt.Print("Enter your move: ")
		scanner.Scan()
		text = scanner.Text()
		if text == "r" {
			gameBoard = board.MoveRight(gameBoard)
		}
	}

}
