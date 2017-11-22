package game

import (
	"bufio"
	"fmt"
	"os"

	"errors"

	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/stats"
)

type Game struct {
}

func New() Game {
	return Game{}
}

func (g Game) Play() {
	gameBoard := board.NewEmptyBoard()
	board.PlaceRandomTile(gameBoard)
	board.PlaceRandomTile(gameBoard)

	score := stats.NewScore()
	player := newHumanPlayer()

	didMove := false
	for {
		if didMove {
			gameBoard = board.PlaceRandomTile(gameBoard)
			didMove = false
		}

		clearScreen()
		fmt.Printf("Score: %v\n", score.Get())
		fmt.Println(board.NewStringer(gameBoard).String() + "\n")

		if board.AreMovesLeft(gameBoard) {
			fmt.Print("Enter move (w,a,s,d): ")
			var err error
			gameBoard, err = player.makeMove(gameBoard, score)
			if err == nil {
				didMove = true
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println("There are no moves left, you lose!")
			break
		}
	}
}

type humanPlayer struct {
	scanner *bufio.Scanner
}

func newHumanPlayer() *humanPlayer {
	return &humanPlayer{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (h humanPlayer) makeMove(gameBoard [][]int64, score *stats.Score) ([][]int64, error) {
	h.scanner.Scan()
	switch h.scanner.Text() {
	case "d":
		return board.MoveRight(gameBoard, score)
		break
	case "a":
		return board.MoveLeft(gameBoard, score)
		break
	case "s":
		return board.MoveDown(gameBoard, score)
		break
	case "w":
		return board.MoveUp(gameBoard, score)
		break
	}
	return gameBoard, errors.New("Unknown move try w, a, s, or d")

}

func clearScreen() {
	print("\033[H\033[2J")
}
