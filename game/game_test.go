package game

import (
	"testing"

	"sync"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/stretchr/testify/assert"
)

type mockPlayer struct {
	moveSignal chan actions.Action
}

func newMockPlayer() mockPlayer {
	return mockPlayer{
		moveSignal: make(chan actions.Action, 0),
	}
}

func (m mockPlayer) GetAction(b [][]int64) actions.Action {
	var a actions.Action
	for {
		select {
		case a = <-m.moveSignal:
			return a
		}
	}
}

func (m mockPlayer) executeAction(a actions.Action) {
	m.moveSignal <- a
}

func TestPlayWithNoMoves(t *testing.T) {
	player := newMockPlayer()

	testGame := New()
	var metrics GameMetrics

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		metrics = testGame.Play(player)
		wg.Done()
	}()

	player.executeAction(actions.Quit)

	wg.Wait()

	assert.Equal(t, int64(0), metrics.MovesMade)
}

func TestPlayWithOneMove(t *testing.T) {
	player := newMockPlayer()

	testGame := Game{
		newBoardFunc: board.NewEmptyBoard,
		placeNewTileFunc: func([][]int64) [][]int64 {
			return [][]int64{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 2},
			}
		},
	}
	var metrics GameMetrics

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		metrics = testGame.Play(player)
		wg.Done()
	}()

	player.executeAction(actions.MoveLeft)
	player.executeAction(actions.Quit)

	wg.Wait()

	assert.Equal(t, int64(1), metrics.MovesMade)
	assert.Equal(t, int64(0), metrics.Score)
}

func TestPlayWithOneMoveWithScore(t *testing.T) {
	player := newMockPlayer()

	testGame := Game{
		newBoardFunc: board.NewEmptyBoard,
		placeNewTileFunc: func([][]int64) [][]int64 {
			return [][]int64{
				{0, 0, 0, 0},
				{0, 32, 0, 0},
				{0, 32, 0, 0},
				{0, 0, 0, 0},
			}
		},
	}
	var metrics GameMetrics

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		metrics = testGame.Play(player)
		wg.Done()
	}()

	player.executeAction(actions.MoveDown)
	player.executeAction(actions.Quit)

	wg.Wait()

	assert.Equal(t, int64(1), metrics.MovesMade)
	assert.Equal(t, int64(64), metrics.Score)
}

func TestPlayWithFiveMoves(t *testing.T) {
	player := newMockPlayer()

	testGame := Game{
		newBoardFunc: func() [][]int64 {
			return [][]int64{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 2},
			}
		},
		placeNewTileFunc: func(b [][]int64) [][]int64 {
			return b
		},
	}
	var metrics GameMetrics

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		metrics = testGame.Play(player)
		wg.Done()
	}()

	player.executeAction(actions.MoveLeft)
	player.executeAction(actions.MoveUp)
	player.executeAction(actions.MoveRight)
	player.executeAction(actions.MoveDown)
	player.executeAction(actions.MoveLeft)
	player.executeAction(actions.Quit)

	wg.Wait()

	assert.Equal(t, int64(5), metrics.MovesMade)
	assert.Equal(t, int64(0), metrics.Score)
}
