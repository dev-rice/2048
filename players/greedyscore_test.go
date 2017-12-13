package players

import (
	"testing"

	"time"

	"github.com/donutmonger/2048/actions"
	"github.com/stretchr/testify/assert"
)

func TestGetActionSimpleHorizontal(t *testing.T) {
	b := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, 2},
	}

	p := NewGreedyScorePlayer(0 * time.Second)
	assert.Equal(t, actions.MoveLeft, p.GetAction(b))
}

func TestGetActionSimpleVertical(t *testing.T) {
	b := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 2},
		{0, 0, 0, 2},
	}

	p := NewGreedyScorePlayer(0 * time.Second)
	assert.Equal(t, actions.MoveUp, p.GetAction(b))
}

func TestGetActionPrioritizesHigherPairs(t *testing.T) {
	b := [][]int64{
		{0, 0, 8, 8},
		{0, 0, 0, 0},
		{0, 0, 0, 2},
		{0, 0, 0, 2},
	}

	p := NewGreedyScorePlayer(0 * time.Second)
	assert.Equal(t, actions.MoveLeft, p.GetAction(b))
}

func TestGetActionPrioritizesHigherPairsMore(t *testing.T) {
	b := [][]int64{
		{0, 8, 4, 8},
		{0, 0, 0, 0},
		{16, 0, 128, 2},
		{16, 0, 128, 2},
	}

	p := NewGreedyScorePlayer(0 * time.Second)
	assert.Equal(t, actions.MoveUp, p.GetAction(b))
}

func TestGetActionCanMoveDown(t *testing.T) {
	b := [][]int64{
		{4, 8, 4, 8},
		{8, 4, 8, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	p := NewGreedyScorePlayer(0 * time.Second)
	assert.Equal(t, actions.MoveDown, p.GetAction(b))
}

func TestGetActionCanMoveRight(t *testing.T) {
	b := [][]int64{
		{4, 8, 0, 0},
		{8, 4, 0, 0},
		{4, 8, 0, 0},
		{8, 4, 0, 0},
	}

	p := NewGreedyScorePlayer(0 * time.Second)
	assert.Equal(t, actions.MoveRight, p.GetAction(b))
}
