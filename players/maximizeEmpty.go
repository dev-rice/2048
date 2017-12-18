package players

import (
	"time"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/boardtree"
)

type MaximizeEmptyPlayer struct {
	delay time.Duration
}

func NewMaximizeEmptyPlayer(delay time.Duration) MaximizeEmptyPlayer {
	return MaximizeEmptyPlayer{
		delay: delay,
	}
}

func (p MaximizeEmptyPlayer) GetAction(gameBoard [][]int64) actions.Action {
	if p.delay != 0*time.Second {
		defer func() { time.Sleep(p.delay) }()
	}

	t := boardtree.Traverser{
		GetScore: getNumEmptyTiles,
		MaxDepth: 3,
	}

	return t.GetBestMove(gameBoard)

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
