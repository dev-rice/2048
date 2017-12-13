package board

import (
	"testing"

	"github.com/donutmonger/2048/stats"
	"github.com/stretchr/testify/assert"
)

func TestNewEmptyBoardReturns4x4OfZeros(t *testing.T) {
	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.Equal(t, expectedBoard, NewEmptyBoard())
}

func TestMoveRightWithEmptyReturnsEmptyAndError(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	s := stats.NewScore()
	board, err := MoveRight(initialBoard, s)
	assert.Equal(t, initialBoard, board)
	assert.NotNil(t, err)

	assert.Equal(t, int64(0), s.Get())
}

func TestMoveRightOneTwo(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 0, 0, 0},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 2},
	}

	s := stats.NewScore()
	actualBoard, err := MoveRight(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s.Get())
}

func TestMoveRightTwoTwos(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 2, 0, 0},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 4},
	}

	s := stats.NewScore()
	actualBoard, err := MoveRight(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s.Get())
}

func TestMoveRightThreeTwos(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 2, 2, 2},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, 4},
	}

	s := stats.NewScore()
	actualBoard, err := MoveRight(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s.Get())
}

func TestMoveRightFourTwos(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 2, 2, 2},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 4, 4},
	}

	s := stats.NewScore()
	actualBoard, err := MoveRight(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(8), s.Get())
}

func TestMoveRightMultipleRows(t *testing.T) {
	initialBoard := [][]int64{
		{2, 128, 0, 0},
		{4, 0, 0, 4},
		{2, 0, 8, 8},
		{2, 2, 2, 32},
	}

	expectedBoard := [][]int64{
		{0, 0, 2, 128},
		{0, 0, 0, 8},
		{0, 0, 2, 16},
		{0, 2, 4, 32},
	}

	s := stats.NewScore()
	actualBoard, err := MoveRight(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(28), s.Get())
}

func TestMoveRightWithNoChangesReturnsError(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 2, 4},
		{0, 0, 4, 8},
		{0, 16, 2, 32},
	}

	s := stats.NewScore()
	actualBoard, err := MoveRight(initialBoard, s)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), s.Get())
}

func TestMoveLeftWithEmptyReturnsEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveLeft(initialBoard, s)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), s.Get())
}

func TestMoveLeftOneTwo(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 2},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveLeft(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s.Get())
}

func TestMoveLeftTwoTwos(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, 2},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{4, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveLeft(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s.Get())
}

func TestMoveLeftThreeTwos(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 2, 2, 0},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{4, 2, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveLeft(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s.Get())
}

func TestMoveLeftFourTwos(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 2, 2, 2},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{4, 4, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveLeft(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(8), s.Get())
}

func TestMoveLeftMultipleRows(t *testing.T) {
	initialBoard := [][]int64{
		{2, 128, 0, 0},
		{4, 0, 0, 4},
		{2, 0, 8, 8},
		{2, 2, 0, 16},
	}

	expectedBoard := [][]int64{
		{2, 128, 0, 0},
		{8, 0, 0, 0},
		{2, 16, 0, 0},
		{4, 16, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveLeft(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(28), s.Get())
}

func TestMoveLeftWithNoChangesReturnsError(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{2, 4, 0, 0},
		{2, 8, 4, 8},
		{2, 16, 2, 32},
	}

	s := stats.NewScore()
	actualBoard, err := MoveLeft(initialBoard, s)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), s.Get())
}

func TestMoveDownWithEmptyReturnsEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveDown(initialBoard, s)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), s.Get())
}

func TestMoveDownOneTwo(t *testing.T) {
	initialBoard := [][]int64{
		{2, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveDown(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s.Get())
}

func TestMoveDownTwoTwos(t *testing.T) {
	initialBoard := [][]int64{
		{2, 0, 0, 0},
		{2, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{4, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveDown(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s.Get())
}

func TestMoveDownThreeTwos(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 2, 0},
		{0, 0, 2, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 0},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 4, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveDown(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s.Get())
}

func TestMoveDownFourTwos(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 2, 0},
		{0, 0, 2, 0},
		{0, 0, 2, 0},
		{0, 0, 2, 0},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 4, 0},
		{0, 0, 4, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveDown(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(8), s.Get())
}

func TestMoveDownMultipleRows(t *testing.T) {
	initialBoard := [][]int64{
		{2, 128, 0, 0},
		{4, 0, 8, 4},
		{2, 0, 8, 8},
		{2, 2, 0, 16},
	}

	expectedBoard := [][]int64{
		{0, 0, 0, 0},
		{2, 0, 0, 4},
		{4, 128, 0, 8},
		{4, 2, 16, 16},
	}

	s := stats.NewScore()
	actualBoard, err := MoveDown(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(20), s.Get())
}

func TestMoveDownWithNoChangesReturnsError(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{2, 4, 0, 0},
		{4, 8, 4, 8},
		{2, 16, 2, 32},
	}

	s := stats.NewScore()
	actualBoard, err := MoveDown(initialBoard, s)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), s.Get())
}

func TestMoveUpWithEmptyReturnsEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveUp(initialBoard, s)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), s.Get())
}

func TestMoveUpOneTwo(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 0, 0, 0},
	}

	expectedBoard := [][]int64{
		{2, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveUp(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), s.Get())
}

func TestMoveUpTwoTwos(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
	}

	expectedBoard := [][]int64{
		{4, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveUp(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s.Get())
}

func TestMoveUpThreeTwos(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 2, 0},
		{0, 0, 2, 0},
	}

	expectedBoard := [][]int64{
		{0, 0, 4, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveUp(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(4), s.Get())
}

func TestMoveUpFourTwos(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 2, 0},
		{0, 0, 2, 0},
		{0, 0, 2, 0},
		{0, 0, 2, 0},
	}

	expectedBoard := [][]int64{
		{0, 0, 4, 0},
		{0, 0, 4, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveUp(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(8), s.Get())
}

func TestMoveUpMultipleRows(t *testing.T) {
	initialBoard := [][]int64{
		{2, 128, 0, 0},
		{4, 0, 8, 4},
		{2, 0, 8, 8},
		{2, 2, 0, 16},
	}

	expectedBoard := [][]int64{
		{2, 128, 16, 4},
		{4, 2, 0, 8},
		{4, 0, 0, 16},
		{0, 0, 0, 0},
	}

	s := stats.NewScore()
	actualBoard, err := MoveUp(initialBoard, s)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
	assert.Equal(t, int64(20), s.Get())
}

func TestMoveUpWithNoChangesReturnsError(t *testing.T) {
	initialBoard := [][]int64{
		{4, 16, 4, 2},
		{2, 8, 2, 4},
		{4, 4, 0, 8},
		{2, 0, 0, 32},
	}

	s := stats.NewScore()
	actualBoard, err := MoveUp(initialBoard, s)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), s.Get())
}

func TestAreMovesLeftReturnsTrueForEmptyBoard(t *testing.T) {
	board := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.True(t, AreMovesLeft(board))
}

func TestAreMovesLeftReturnsFalseForFullStaggeredBoard(t *testing.T) {
	board := [][]int64{
		{2, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 4, 2, 4},
		{4, 2, 4, 2},
	}

	assert.False(t, AreMovesLeft(board))
}

func TestAreMovesLeftReturnsTrueForStaggeredBoardWithOneEmpty(t *testing.T) {
	board := [][]int64{
		{0, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 4, 2, 4},
		{4, 2, 4, 2},
	}

	assert.True(t, AreMovesLeft(board))
}
