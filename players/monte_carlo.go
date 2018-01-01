package players

import (
	"sort"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/game"
)

type actionTuple struct {
	score  float32
	action actions.Action
}
type actionTuples []actionTuple

func (s actionTuples) Len() int {
	return len(s)
}
func (s actionTuples) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s actionTuples) Less(i, j int) bool {
	return s[i].score < s[j].score
}

type silentPrinter struct{}

func (p silentPrinter) Printf(format string, v ...interface{}) {
}

func (p silentPrinter) ClearScreen() {
}

type MonteCarloPlayer struct {
}

func NewMonteCarloPlayer() MonteCarloPlayer {
	return MonteCarloPlayer{}
}

func (m MonteCarloPlayer) GetAction(b [][]int64) actions.Action {
	a := make([]actionTuple, 0)

	timesPerMove := 100

	upBoard, upScore, err := board.MoveUp(b)
	if err == nil {
		a = append(a, actionTuple{action: actions.MoveUp, score: getAverageScoreForNRandomGamesForBoard(timesPerMove, upBoard) + float32(upScore)})
	}

	downBoard, downScore, err := board.MoveDown(b)
	if err == nil {
		a = append(a, actionTuple{action: actions.MoveDown, score: getAverageScoreForNRandomGamesForBoard(timesPerMove, downBoard) + float32(downScore)})
	}

	leftBoard, leftScore, err := board.MoveLeft(b)
	if err == nil {
		a = append(a, actionTuple{action: actions.MoveLeft, score: getAverageScoreForNRandomGamesForBoard(timesPerMove, leftBoard) + float32(leftScore)})
	}

	rightBoard, rightScore, err := board.MoveRight(b)
	if err == nil {
		a = append(a, actionTuple{action: actions.MoveRight, score: getAverageScoreForNRandomGamesForBoard(timesPerMove, rightBoard) + float32(rightScore)})
	}

	sort.Sort(actionTuples(a))
	return a[len(a)-1].action
}

func getAverageScoreForNRandomGamesForBoard(n int, board [][]int64) float32 {
	sum := float32(0)
	for i := 0; i < n; i++ {
		g := game.New()
		g.NewBoardFunc = func() [][]int64 {
			return board
		}
		sum += float32(g.Play(NewRandomPlayer(), silentPrinter{}).Score)
	}

	return sum / float32(n)
}
