package ai

import (
	"testing"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/stretchr/testify/assert"
)

func areBoardsSame(a int64, b int64) bool {
	return a == b
}

func TestGetBestMoveLeft(t *testing.T) {
	initialBoard := board.CompressBoardGrid([][]int64{
		{0, 0, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	})

	traverser := Traverser{
		GetRating: func(b [][]int64) uint64 {
			leftBoard, _, _ := board.MoveLeft(initialBoard)
			if areBoardsSame(board.CompressBoardGrid(b), leftBoard) {
				return 100
			} else {
				return 0
			}
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveLeft, traverser.GetBestMove(board.UncompressBoard(initialBoard)))
}

func TestGetBestMoveRight(t *testing.T) {
	initialBoard := board.CompressBoardGrid([][]int64{
		{0, 0, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	})

	traverser := Traverser{
		GetRating: func(b [][]int64) uint64 {
			rightBoard, _, _ := board.MoveRight(initialBoard)
			if areBoardsSame(board.CompressBoardGrid(b), rightBoard) {
				return 100
			} else {
				return 0
			}
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveRight, traverser.GetBestMove(board.UncompressBoard(initialBoard)))
}

func TestGetBestMoveUp(t *testing.T) {
	initialBoard := board.CompressBoardGrid([][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, 2},
	})

	traverser := Traverser{
		GetRating: func(b [][]int64) uint64 {
			upBoard, _, _ := board.MoveUp(initialBoard)
			if areBoardsSame(board.CompressBoardGrid(b), upBoard) {
				return 100
			} else {
				return 0
			}
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveUp, traverser.GetBestMove(board.UncompressBoard(initialBoard)))
}

func TestGetBestMoveDown(t *testing.T) {
	initialBoard := board.CompressBoardGrid([][]int64{
		{0, 0, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	})

	traverser := Traverser{
		GetRating: func(b [][]int64) uint64 {
			downBoard, _, _ := board.MoveDown(initialBoard)
			if areBoardsSame(board.CompressBoardGrid(b), downBoard) {
				return 100
			} else {
				return 0
			}
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveDown, traverser.GetBestMove(board.UncompressBoard(initialBoard)))
}

func TestGetBestMoveAllMovesSameNoError(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 0},
	}

	traverser := Traverser{
		GetRating: func(b [][]int64) uint64 {
			return 0
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveRight, traverser.GetBestMove(initialBoard))
}

func TestGetBestMoveAllMovesSameScoreErrorUpAndLeftAndRightMovesDown(t *testing.T) {
	initialBoard := board.CompressBoardGrid([][]int64{
		{2, 4, 2, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	})

	_, _, err := board.MoveUp(initialBoard)
	assert.NotNil(t, err)

	_, _, err = board.MoveLeft(initialBoard)
	assert.NotNil(t, err)

	_, _, err = board.MoveRight(initialBoard)
	assert.NotNil(t, err)

	traverser := Traverser{
		GetRating: func(b [][]int64) uint64 {
			return 0
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveDown, traverser.GetBestMove(board.UncompressBoard(initialBoard)))
}

func TestGetBestMoveAllMovesSameScoreErrorDownAndLeftAndRightMovesUp(t *testing.T) {
	initialBoard := board.CompressBoardGrid([][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 4, 2, 4},
	})

	_, _, err := board.MoveDown(initialBoard)
	assert.NotNil(t, err)

	_, _, err = board.MoveLeft(initialBoard)
	assert.NotNil(t, err)

	_, _, err = board.MoveRight(initialBoard)
	assert.NotNil(t, err)

	traverser := Traverser{
		GetRating: func(b [][]int64) uint64 {
			return 0
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveUp, traverser.GetBestMove(board.UncompressBoard(initialBoard)))
}

func TestGetBestMoveAllMovesSameScoreErrorUpAndDownAndRightMovesLeft(t *testing.T) {
	initialBoard := board.CompressBoardGrid([][]int64{
		{0, 0, 0, 2},
		{0, 0, 0, 4},
		{0, 0, 0, 2},
		{0, 0, 0, 4},
	})

	_, _, err := board.MoveUp(initialBoard)
	assert.NotNil(t, err)

	_, _, err = board.MoveDown(initialBoard)
	assert.NotNil(t, err)

	_, _, err = board.MoveRight(initialBoard)
	assert.NotNil(t, err)

	traverser := Traverser{
		GetRating: func(b [][]int64) uint64 {
			return 0
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveLeft, traverser.GetBestMove(board.UncompressBoard(initialBoard)))
}

func TestGetBestMoveAllMovesSameScoreErrorUpAndDownAndLeftMovesRight(t *testing.T) {
	initialBoard := board.CompressBoardGrid([][]int64{
		{2, 0, 0, 0},
		{4, 0, 0, 0},
		{2, 0, 0, 0},
		{4, 0, 0, 0},
	})

	_, _, err := board.MoveUp(initialBoard)
	assert.NotNil(t, err)

	_, _, err = board.MoveDown(initialBoard)
	assert.NotNil(t, err)

	_, _, err = board.MoveLeft(initialBoard)
	assert.NotNil(t, err)

	traverser := Traverser{
		GetRating: func(b [][]int64) uint64 {
			return 0
		},
		MaxDepth: 1,
	}

	assert.Equal(t, actions.MoveRight, traverser.GetBestMove(board.UncompressBoard(initialBoard)))
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
	initialBoard := board.CompressBoardGrid([][]int64{
		{0, 0, 0, 0},
		{0, 0, 8, 4},
		{8, 2, 2, 4},
		{2, 4, 2, 32},
	})

	r := buildRoot(initialBoard, getNumEmptyTiles, 1)
	assert.Equal(t, uint64(8), r.up.rating)
	assert.Equal(t, uint64(8), r.down.rating)
	assert.Equal(t, uint64(7), r.left.rating)
	assert.Equal(t, uint64(7), r.right.rating)
}

func TestBuildRootDepth2MaximizeEmpty(t *testing.T) {
	initialBoard := board.CompressBoardGrid([][]int64{
		{0, 0, 0, 0},
		{0, 0, 8, 4},
		{8, 2, 2, 4},
		{2, 4, 2, 32},
	})

	r := buildRoot(initialBoard, getNumEmptyTiles, 2)
	assert.Equal(t, uint64(10), r.up.rating)
	assert.Equal(t, uint64(10), r.down.rating)
	assert.Equal(t, uint64(9), r.left.rating)
	assert.Equal(t, uint64(8), r.right.rating)
}

func TestBuildRootDepth3MaximizeEmpty(t *testing.T) {
	initialBoard := board.CompressBoardGrid([][]int64{
		{0, 0, 0, 0},
		{2, 2, 8, 4},
		{8, 2, 2, 4},
		{2, 4, 2, 32},
	})

	r := buildRoot(initialBoard, getNumEmptyTiles, 3)
	assert.Equal(t, uint64(10), r.up.rating)
	assert.Equal(t, uint64(9), r.down.rating)
	assert.Equal(t, uint64(10), r.left.rating)
	assert.Equal(t, uint64(9), r.right.rating)
}
