package game

import (
	"testing"

	"sync"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/stretchr/testify/assert"
)

type mockPrinter struct {
	PrintfCall struct {
		Times    int
		Receives struct {
			Format string
			Values []interface{}
		}
	}
	ClearScreenCall struct {
		Times int
	}
}

func (p *mockPrinter) Printf(format string, v ...interface{}) {
	p.PrintfCall.Receives.Format = format
	p.PrintfCall.Receives.Values = v
	p.PrintfCall.Times += 1
}

func (p *mockPrinter) ClearScreen() {
	p.ClearScreenCall.Times += 1
}

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
	printer := &mockPrinter{}

	testGame := New()
	var metrics GameMetrics

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		metrics = testGame.Play(player, printer)
		wg.Done()
	}()

	player.executeAction(actions.Quit)

	wg.Wait()

	assert.Equal(t, int64(0), metrics.MovesMade)
	assert.Equal(t, int64(0), metrics.Score)
	assert.Equal(t, 1, printer.ClearScreenCall.Times)
	assert.Equal(t, 3, printer.PrintfCall.Times)
}

func TestPlayWithOneMove(t *testing.T) {
	player := newMockPlayer()
	printer := &mockPrinter{}

	testGame := Game{
		newBoardFunc: board.NewEmptyBoard,
		placeNewTileFunc: func(int64) int64 {
			return board.CompressBoardGrid([][]int64{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 2},
			})
		},
	}
	var metrics GameMetrics

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		metrics = testGame.Play(player, printer)
		wg.Done()
	}()

	player.executeAction(actions.MoveLeft)
	player.executeAction(actions.Quit)

	wg.Wait()

	assert.Equal(t, int64(1), metrics.MovesMade)
	assert.Equal(t, int64(0), metrics.Score)
	assert.Equal(t, int64(2), metrics.BiggestTile)
	assert.Equal(t, 2, printer.ClearScreenCall.Times)
	assert.Equal(t, 5, printer.PrintfCall.Times)
}

func TestPlayWithOneMoveWithScore(t *testing.T) {
	player := newMockPlayer()
	printer := &mockPrinter{}

	testGame := Game{
		newBoardFunc: func() int64 {
			return board.CompressBoardGrid([][]int64{
				{0, 0, 0, 0},
				{0, 32, 0, 2},
				{0, 32, 0, 2},
				{0, 0, 0, 0},
			})
		},
		placeNewTileFunc: func(b int64) int64 {
			return b
		},
	}
	var metrics GameMetrics

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		metrics = testGame.Play(player, printer)
		wg.Done()
	}()

	player.executeAction(actions.MoveDown)
	player.executeAction(actions.Quit)

	wg.Wait()

	assert.Equal(t, int64(1), metrics.MovesMade)
	assert.Equal(t, int64(68), metrics.Score)
	assert.Equal(t, int64(64), metrics.BiggestTile)
	assert.Equal(t, 2, printer.ClearScreenCall.Times)
	assert.Equal(t, 5, printer.PrintfCall.Times)
}

func TestPlayWithFiveMoves(t *testing.T) {
	player := newMockPlayer()
	printer := &mockPrinter{}

	testGame := Game{
		newBoardFunc: func() int64 {
			return board.CompressBoardGrid([][]int64{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 2},
			})
		},
		placeNewTileFunc: func(b int64) int64 {
			return b
		},
	}
	var metrics GameMetrics

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		metrics = testGame.Play(player, printer)
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
	assert.Equal(t, int64(2), metrics.BiggestTile)
	assert.Equal(t, 6, printer.ClearScreenCall.Times)
	assert.Equal(t, 13, printer.PrintfCall.Times)
}
