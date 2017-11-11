package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/donutmonger/2048/board"
	"github.com/fatih/color"
)

func main() {
	gameBoard := board.NewEmptyBoard()
	gameBoard[0][3] = 2

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "q" { // break the loop if text == "q"
		gameBoard = board.PlaceRandomTwo(gameBoard)
		fmt.Println(board.BoardToString(gameBoard))
		fmt.Println("")

		d := color.New(color.FgCyan, color.Bold)
		d.Println("hello world")

		fmt.Print("Enter your move: ")
		scanner.Scan()
		text = scanner.Text()

		switch text {
		case "r":
			gameBoard = board.MoveRight(gameBoard)
			break
		case "l":
			gameBoard = board.MoveLeft(gameBoard)
			break
		case "d":
			gameBoard = board.MoveDown(gameBoard)
			break
		case "u":
			gameBoard = board.MoveUp(gameBoard)
			break
		}

	}

}
