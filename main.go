package main

import (
	"fmt"

	"github.com/donutmonger/2048/board"
	"github.com/eiannone/keyboard"
)

func main() {
	gameBoard := board.NewEmptyBoard()
	board.PlaceRandomTile(gameBoard)
	board.PlaceRandomTile(gameBoard)

	didMove := false
	didQuit := false
	for !didQuit {
		clearScreen()

		if didMove {
			gameBoard = board.PlaceRandomTile(gameBoard)
			didMove = false
		}
		fmt.Println(board.BoardToString(gameBoard) + "\n")

		if !board.AreMovesLeft(gameBoard) {
			fmt.Println("You lose, no moves left")
			break
		}

		fmt.Print("Enter your move (w,a,s,d) ")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%q\n", char)
		switch char {
		case 'd':
			gameBoard = board.MoveRight(gameBoard)
			didMove = true
			break
		case 'a':
			gameBoard = board.MoveLeft(gameBoard)
			didMove = true
			break
		case 's':
			gameBoard = board.MoveDown(gameBoard)
			didMove = true
			break
		case 'w':
			gameBoard = board.MoveUp(gameBoard)
			didMove = true
			break
		case 'q':
			fmt.Println("Exiting game")
			didQuit = true
			break
		default:
			fmt.Println("Unknown move try w, a, s, or d")
		}
	}
}

func clearScreen() {
	print("\033[H\033[2J")
}
