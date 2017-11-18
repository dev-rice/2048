package main

import (
	"fmt"

	"bufio"
	"os"

	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/game"
)

// Add a test color board (whenever cli is a thing)
func main() {
	gameBoard := board.NewEmptyBoard()
	board.PlaceRandomTile(gameBoard)
	board.PlaceRandomTile(gameBoard)

	score := game.NewScore()

	scanner := bufio.NewScanner(os.Stdin)

	didMove := false
	didQuit := false
	for !didQuit {
		if didMove {
			gameBoard = board.PlaceRandomTile(gameBoard)
			didMove = false
		}

		clearScreen()
		fmt.Printf("Score: %v\n", score.Get())
		fmt.Println(board.NewStringer(gameBoard).String() + "\n")

		if board.AreMovesLeft(gameBoard) {
			fmt.Print("Enter move (w,a,s,d): ")
		} else {
			fmt.Println("There are no moves left, you lose!")
			break
		}

		var err error
		scanner.Scan()
		switch scanner.Text() {
		case "d":
			gameBoard, err = board.MoveRight(gameBoard, score)
			if err == nil {
				didMove = true
			}
			break
		case "a":
			gameBoard, err = board.MoveLeft(gameBoard, score)
			if err == nil {
				didMove = true
			}
			break
		case "s":
			gameBoard, err = board.MoveDown(gameBoard, score)
			if err == nil {
				didMove = true
			}
			break
		case "w":
			gameBoard, err = board.MoveUp(gameBoard, score)
			if err == nil {
				didMove = true
			}
			break
		case "q":
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
