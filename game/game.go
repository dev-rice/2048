package game

import (
	"time"

	"fmt"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/players"
)

type GameStats struct {
	MovesMade          int64
	Score              int64
	ElapsedTimeSeconds float64
}

type Game struct {
	placeNewTileFunc func(board [][]int64) [][]int64
	newBoardFunc     func() [][]int64
}

func New() Game {
	return Game{
		placeNewTileFunc: board.PlaceRandomTile,
	}
}

func (g Game) Play(player players.Player) (stats GameStats) {
	gameBoard := board.NewEmptyBoard()
	gameBoard = g.placeNewTileFunc(gameBoard)
	gameBoard = g.placeNewTileFunc(gameBoard)

	start := time.Now()
	defer func() {
		stats.ElapsedTimeSeconds = time.Since(start).Seconds()
	}()

	didMove := false
	for {
		if didMove {
			gameBoard = g.placeNewTileFunc(gameBoard)
			didMove = false
		}

		clearScreen()
		fmt.Printf("Score: %v\n", stats.Score)
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
				return stats
			}
			if err == nil {
				stats.MovesMade += 1
				stats.Score += scoreAdd
				didMove = true
			}
		} else {
			fmt.Println("There are no moves left, you lose!")
			break
		}
	}
	return stats
}

func clearScreen() {
	print("\033[H\033[2J")
}
