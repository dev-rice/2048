package players

import (
	"time"

	"github.com/donutmonger/2048/actions"
)

const noDelay = 0 * time.Second

type bestMoveGetter interface {
	GetBestMove([][]int64) actions.Action
}

type AIPlayer struct {
	moveGetter bestMoveGetter
	delay      time.Duration
}

func NewAIPlayer(delay time.Duration, m bestMoveGetter) AIPlayer {
	return AIPlayer{
		moveGetter: m,
		delay:      delay,
	}
}

func (p AIPlayer) GetAction(gameBoard [][]int64) actions.Action {
	if p.delay != noDelay {
		defer func() { time.Sleep(p.delay) }()
	}
	return p.moveGetter.GetBestMove(gameBoard)
}
