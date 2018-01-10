package players

import (
	"math/rand"

	"github.com/donutmonger/2048/actions"
)

type RandomPlayer struct {
}

func NewRandomPlayer() RandomPlayer {
	return RandomPlayer{}
}

func (r RandomPlayer) GetAction(int64) actions.Action {
	return actions.Action(rand.Intn(4))
}
