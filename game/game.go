package game

import (
	"fmt"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/players"
	"github.com/donutmonger/2048/stats"
)

type Game struct {
}

func New() Game {
	return Game{}
}

func (g Game) Play(player players.Player) {
	gameBoard := board.NewEmptyBoard()
	board.PlaceRandomTile(gameBoard)
	board.PlaceRandomTile(gameBoard)

	score := stats.NewScore()
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
			action := player.GetAction(gameBoard)

			var err error
			switch action {
			case actions.MoveUp:
				gameBoard, err = board.MoveUp(gameBoard, score)
				break
			case actions.MoveDown:
				gameBoard, err = board.MoveDown(gameBoard, score)
				break
			case actions.MoveLeft:
				gameBoard, err = board.MoveLeft(gameBoard, score)
				break
			case actions.MoveRight:
				gameBoard, err = board.MoveRight(gameBoard, score)
				break
			case actions.Quit:
				fmt.Println("Quitting...")
				return
			}
			if err == nil {
				didMove = true
			}
		} else {
			fmt.Println("There are no moves left, you lose!")
			break
		}
	}
}

func clearScreen() {
	print("\033[H\033[2J")
}
