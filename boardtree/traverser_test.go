package boardtree

import (
	"testing"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/stats"
	"github.com/stretchr/testify/assert"
)

func areBoardsSame(a [][]int64, b [][]int64) bool {
	for x := 0; x < len(a[0]); x++ {
		for y := 0; y < len(a); y++ {
			if a[x][y] != b[x][y] {
				return false
			}
		}
	}
	return true
}

func TestGetBestMoveLeft(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	traverser := Traverser{
		GetScore: func(b [][]int64) uint64 {
			leftBoard, _ := board.MoveLeft(initialBoard, stats.NewScore())
			if areBoardsSame(b, leftBoard) {
				return 100
			} else {
				return 0
			}
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveLeft, traverser.GetBestMove(initialBoard))
}

func TestGetBestMoveRight(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	traverser := Traverser{
		GetScore: func(b [][]int64) uint64 {
			rightBoard, _ := board.MoveRight(initialBoard, stats.NewScore())
			if areBoardsSame(b, rightBoard) {
				return 100
			} else {
				return 0
			}
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveRight, traverser.GetBestMove(initialBoard))
}

func TestGetBestMoveUp(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, 2},
	}

	traverser := Traverser{
		GetScore: func(b [][]int64) uint64 {
			upBoard, _ := board.MoveUp(initialBoard, stats.NewScore())
			if areBoardsSame(b, upBoard) {
				return 100
			} else {
				return 0
			}
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveUp, traverser.GetBestMove(initialBoard))
}

func TestGetBestMoveDown(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	traverser := Traverser{
		GetScore: func(b [][]int64) uint64 {
			downBoard, _ := board.MoveDown(initialBoard, stats.NewScore())
			if areBoardsSame(b, downBoard) {
				return 100
			} else {
				return 0
			}
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveDown, traverser.GetBestMove(initialBoard))
}

func TestGetBestMoveAllMovesSameNoError(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 0},
	}

	traverser := Traverser{
		GetScore: func(b [][]int64) uint64 {
			return 0
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveRight, traverser.GetBestMove(initialBoard))
}

func TestGetBestMoveAllMovesSameScoreErrorUpAndLeftAndRightMovesDown(t *testing.T) {
	initialBoard := [][]int64{
		{2, 4, 2, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	_, err := board.MoveUp(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	_, err = board.MoveLeft(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	_, err = board.MoveRight(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	traverser := Traverser{
		GetScore: func(b [][]int64) uint64 {
			return 0
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveDown, traverser.GetBestMove(initialBoard))
}

func TestGetBestMoveAllMovesSameScoreErrorDownAndLeftAndRightMovesUp(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 4, 2, 4},
	}

	_, err := board.MoveDown(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	_, err = board.MoveLeft(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	_, err = board.MoveRight(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	traverser := Traverser{
		GetScore: func(b [][]int64) uint64 {
			return 0
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveUp, traverser.GetBestMove(initialBoard))
}

func TestGetBestMoveAllMovesSameScoreErrorUpAndDownAndRightMovesLeft(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 2},
		{0, 0, 0, 4},
		{0, 0, 0, 2},
		{0, 0, 0, 4},
	}

	_, err := board.MoveUp(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	_, err = board.MoveDown(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	_, err = board.MoveRight(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	traverser := Traverser{
		GetScore: func(b [][]int64) uint64 {
			return 0
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveLeft, traverser.GetBestMove(initialBoard))
}

func TestGetBestMoveAllMovesSameScoreErrorUpAndDownAndLeftMovesRight(t *testing.T) {
	initialBoard := [][]int64{
		{2, 0, 0, 0},
		{4, 0, 0, 0},
		{2, 0, 0, 0},
		{4, 0, 0, 0},
	}

	_, err := board.MoveUp(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	_, err = board.MoveDown(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	_, err = board.MoveLeft(initialBoard, stats.NewScore())
	assert.NotNil(t, err)

	traverser := Traverser{
		GetScore: func(b [][]int64) uint64 {
			return 0
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveRight, traverser.GetBestMove(initialBoard))
}

// Probably test error case when it somehow gets a board with no moves?

// These tests should probably be in a separate package
func getNumEmptyTiles(board [][]int64) uint64 {
	emptyTiles := uint64(0)
	for x := 0; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			if board[x][y] == 0 {
				emptyTiles++
			}
		}
	}
	return emptyTiles
}

func TestBuildRootDepth1MaximizeEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 8, 4},
		{8, 2, 2, 4},
		{2, 4, 2, 32},
	}

	r := buildRoot(initialBoard, getNumEmptyTiles, 1)
	assert.Equal(t, uint64(8), r.up.score)
	assert.Equal(t, uint64(8), r.down.score)
	assert.Equal(t, uint64(7), r.left.score)
	assert.Equal(t, uint64(7), r.right.score)
}

func TestBuildRootDepth2MaximizeEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 8, 4},
		{8, 2, 2, 4},
		{2, 4, 2, 32},
	}

	r := buildRoot(initialBoard, getNumEmptyTiles, 2)
	assert.Equal(t, uint64(10), r.up.score)
	assert.Equal(t, uint64(10), r.down.score)
	assert.Equal(t, uint64(9), r.left.score)
	assert.Equal(t, uint64(8), r.right.score)
}

func TestBuildRootDepth3MaximizeEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{2, 2, 8, 4},
		{8, 2, 2, 4},
		{2, 4, 2, 32},
	}

	r := buildRoot(initialBoard, getNumEmptyTiles, 3)
	assert.Equal(t, uint64(10), r.up.score)
	assert.Equal(t, uint64(9), r.down.score)
	assert.Equal(t, uint64(10), r.left.score)
	assert.Equal(t, uint64(9), r.right.score)
}
