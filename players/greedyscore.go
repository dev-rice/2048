package players

import (
	"sort"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/stats"
)

type GreedyScorePlayer struct {
}

func NewGreedyScorePlayer() GreedyScorePlayer {
	return GreedyScorePlayer{}
}

func (g GreedyScorePlayer) GetAction(gameBoard [][]int64) actions.Action {
	upScore := stats.NewScore()
	_, err := board.MoveUp(gameBoard, upScore)
	if err != nil {
		upScore = stats.NewScore()
		upScore.Add(-1)
	}

	downScore := stats.NewScore()
	_, err = board.MoveDown(gameBoard, downScore)
	if err != nil {
		downScore = stats.NewScore()
		downScore.Add(-1)
	}

	leftScore := stats.NewScore()
	_, err = board.MoveLeft(gameBoard, leftScore)
	if err != nil {
		leftScore = stats.NewScore()
		leftScore.Add(-1)
	}

	rightScore := stats.NewScore()
	_, err = board.MoveRight(gameBoard, rightScore)
	if err != nil {
		rightScore = stats.NewScore()
		rightScore.Add(-1)
	}

	scores := []int{int(upScore.Get()), int(downScore.Get()), int(leftScore.Get()), int(rightScore.Get())}
	sort.Ints(scores)

	best := scores[len(scores)-1]
	if best == int(upScore.Get()) {
		return actions.MoveUp
	} else if best == int(downScore.Get()) {
		return actions.MoveDown
	} else if best == int(leftScore.Get()) {
		return actions.MoveLeft
	} else if best == int(rightScore.Get()) {
		return actions.MoveRight
	}

	return actions.MoveRight
}
