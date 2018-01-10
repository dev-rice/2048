package game

import (
	"time"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/players"
)

// Ideas for more metrics:
// 		number of up, left, right, and down moves
//		longest time spent moving
type GameMetrics struct {
	MovesMade          int64
	Score              int64
	ElapsedTimeSeconds float64
	BiggestTile        int64
}

type printer interface {
	Printf(format string, v ...interface{})
	ClearScreen()
}

type Game struct {
	newBoardFunc     func() int64
	placeNewTileFunc func(board int64) int64
}

func New() Game {
	return Game{
		newBoardFunc:     board.NewEmptyBoard,
		placeNewTileFunc: board.PlaceRandomTile,
	}
}

func (g Game) Play(player players.Player, printer printer) (metrics GameMetrics) {
	gameBoard := g.newBoardFunc()
	gameBoard = g.placeNewTileFunc(gameBoard)
	gameBoard = g.placeNewTileFunc(gameBoard)

	start := time.Now()
	defer func() {
		metrics.ElapsedTimeSeconds = time.Since(start).Seconds()
		// should do this somewhere else
		metrics.BiggestTile = getBiggestTile(gameBoard)
	}()

	didMove := false
	for {
		if didMove {
			gameBoard = g.placeNewTileFunc(gameBoard)
			didMove = false
		}

		printer.ClearScreen()
		printer.Printf("Score: %v\n", metrics.Score)
		printer.Printf("%s\n\n", board.NewStringer(board.ExtractGridFromBoard(gameBoard)))

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
				printer.Printf("Quitting...\n")
				return metrics
			}
			if err == nil {
				metrics.MovesMade += 1
				metrics.Score += scoreAdd
				didMove = true
			}
		} else {
			printer.Printf("There are no moves left, you lose!\n")
			break
		}
	}
	return metrics
}

func getBiggestTile(compressed int64) int64 {
	b := board.ExtractGridFromBoard(compressed)
	biggest := int64(0)
	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b[0]); x++ {
			if b[y][x] > biggest {
				biggest = b[y][x]
			}
		}
	}
	return biggest
}
