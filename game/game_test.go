package game

import (
	"testing"

	"sync"

	"github.com/donutmonger/2048/actions"
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
	var stats GameStats

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stats = testGame.Play(player)
		wg.Done()
	}()

	player.executeAction(actions.Quit)

	wg.Wait()

	assert.Equal(t, int64(0), stats.MovesMade)
}

func TestPlayWithOneMove(t *testing.T) {
	player := newMockPlayer()

	testGame := Game{
		placeNewTileFunc: func([][]int64) [][]int64 {
			return [][]int64{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 2},
			}
		},
	}
	var stats GameStats

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stats = testGame.Play(player)
		wg.Done()
	}()

	player.executeAction(actions.MoveLeft)
	player.executeAction(actions.Quit)

	wg.Wait()

	assert.Equal(t, int64(1), stats.MovesMade)
	assert.Equal(t, int64(0), stats.Score)
}

func TestPlayWithOneMoveWithScore(t *testing.T) {
	player := newMockPlayer()

	testGame := Game{
		placeNewTileFunc: func([][]int64) [][]int64 {
			return [][]int64{
				{0, 0, 0, 0},
				{0, 32, 0, 0},
				{0, 32, 0, 0},
				{0, 0, 0, 0},
			}
		},
	}
	var stats GameStats

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stats = testGame.Play(player)
		wg.Done()
	}()

	player.executeAction(actions.MoveDown)
	player.executeAction(actions.Quit)

	wg.Wait()

	assert.Equal(t, int64(1), stats.MovesMade)
	assert.Equal(t, int64(64), stats.Score)
}
