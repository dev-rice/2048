package players

import (
	"time"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/stats"
)

type GreedyMinimizePlayer struct {
	delay time.Duration
}

func NewGreedyMinimizePlayer(delay time.Duration) GreedyMinimizePlayer {
	return GreedyMinimizePlayer{
		delay: delay,
	}
}

func (p GreedyMinimizePlayer) GetAction(gameBoard [][]int64) actions.Action {
	if p.delay != 0*time.Second {
		defer func() { time.Sleep(p.delay) }()
	}

	nTilesBase := getNumTiles(gameBoard)

	upBoard, err := board.MoveUp(gameBoard, stats.NewScore())
	nTilesUp := getNumTiles(upBoard)
	if err != nil {
		nTilesUp += 1
	}

	leftBoard, err := board.MoveLeft(gameBoard, stats.NewScore())
	nTilesLeft := getNumTiles(leftBoard)
	if err != nil {
		nTilesLeft += 1
	}

	rightBoard, err := board.MoveRight(gameBoard, stats.NewScore())
	nTilesRight := getNumTiles(rightBoard)
	if err != nil {
		nTilesRight += 1
	}

	if nTilesUp < nTilesLeft && nTilesUp <= nTilesBase {
		return actions.MoveUp
	} else if nTilesLeft <= nTilesUp && nTilesLeft <= nTilesBase {
		return actions.MoveLeft
	} else if nTilesRight <= nTilesBase {
		return actions.MoveRight
	} else {
		return actions.MoveDown
	}
}

func getNumTiles(board [][]int64) int {
	nTiles := 0
	for x := 0; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			if board[x][y] != 0 {
				nTiles++
			}
		}
	}
	return nTiles
}
