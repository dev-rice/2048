package players

import (
	"time"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/boardtree"
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

	return boardtree.GetBestMove(gameBoard, getNumEmptyTiles)

}

func getNumEmptyTiles(board [][]int64) uint64 {
	emptyTiles := uint64(0)
	for x := 0; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			if board[x][y] == 0 {
				emptyTiles++
			}
		}
	}
	return emptyTiles
}
