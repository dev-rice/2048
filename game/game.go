package game

import (
	"fmt"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/players"
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

	var score int64
	didMove := false
	for {
		if didMove {
			gameBoard = board.PlaceRandomTile(gameBoard)
			didMove = false
		}

		clearScreen()
		fmt.Printf("Score: %v\n", score)
		fmt.Println(board.NewStringer(gameBoard).String() + "\n")

		if board.AreMovesLeft(gameBoard) {
			action := player.GetAction(gameBoard)

			var scoreAdd int64
			var err error
			switch action {
			case actions.MoveUp:
				gameBoard, scoreAdd, err = board.MoveUp(gameBoard)
				break
			case actions.MoveDown:
				gameBoard, scoreAdd, err = board.MoveDown(gameBoard)
				break
			case actions.MoveLeft:
				gameBoard, scoreAdd, err = board.MoveLeft(gameBoard)
				break
			case actions.MoveRight:
				gameBoard, scoreAdd, err = board.MoveRight(gameBoard)
				break
			case actions.Quit:
				fmt.Println("Quitting...")
				return
			}
			if err == nil {
				score += scoreAdd
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
