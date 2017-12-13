package players

import (
	"math/rand"
	"time"

	"github.com/donutmonger/2048/actions"
)

type RandomPlayer struct {
}

func NewRandomPlayer() RandomPlayer {
	rand.Seed(time.Now().UnixNano())
	return RandomPlayer{}
}

func (r RandomPlayer) GetAction(gameBoard [][]int64) actions.Action {
	actionMap := map[int]actions.Action{
		0: actions.MoveUp,
		1: actions.MoveDown,
		2: actions.MoveLeft,
		3: actions.MoveRight,
	}
	return actionMap[rand.Intn(4)]
}
