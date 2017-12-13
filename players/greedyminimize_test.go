package players

import (
	"testing"

	"time"

	"github.com/donutmonger/2048/actions"
	"github.com/stretchr/testify/assert"
)

func TestSimpleMoveUp(t *testing.T) {
	b := [][]int64{
		{8, 4, 0, 0},
		{8, 4, 4, 2},
		{2, 2, 4, 2},
		{2, 2, 32, 32},
	}

	p := NewGreedyMinimizePlayer(0 * time.Second)
	assert.Equal(t, actions.MoveUp, p.GetAction(b))
}

func TestSimpleMoveLeft(t *testing.T) {
	b := [][]int64{
		{64, 4, 4, 2},
		{64, 2, 2, 0},
		{32, 4, 4, 4},
		{32, 2, 2, 2},
	}

	p := NewGreedyMinimizePlayer(0 * time.Second)
	assert.Equal(t, actions.MoveLeft, p.GetAction(b))
}

func TestCanMoveRight(t *testing.T) {
	b := [][]int64{
		{64, 4, 0, 0},
		{32, 2, 0, 0},
		{64, 4, 0, 0},
		{32, 2, 0, 0},
	}

	p := NewGreedyMinimizePlayer(0 * time.Second)
	assert.Equal(t, actions.MoveRight, p.GetAction(b))
}

func TestCanMoveDown(t *testing.T) {
	b := [][]int64{
		{64, 16, 32, 8},
		{32, 8, 16, 64},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	p := NewGreedyMinimizePlayer(0 * time.Second)
	assert.Equal(t, actions.MoveDown, p.GetAction(b))
}

func TestCanMoveUp(t *testing.T) {
	b := [][]int64{
		{0, 0, 2, 16},
		{0, 0, 128, 4},
		{0, 64, 16, 2},
		{0, 2, 8, 4},
	}

	p := NewGreedyMinimizePlayer(0 * time.Second)
	assert.Equal(t, actions.MoveLeft, p.GetAction(b))
}
