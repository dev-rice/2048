package board

import (
	"testing"

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

func TestBoardToStringEmpty(t *testing.T) {
	expected :=
		`-----------------
|   |   |   |   |
-----------------
|   |   |   |   |
-----------------
|   |   |   |   |
-----------------
|   |   |   |   |
-----------------`

	board := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	actual := BoardToString(board)

	assert.Equal(t, expected, actual)
}

func TestBoardToString(t *testing.T) {
	expected :=
		`-------------------------
|     |     |     |   2 |
-------------------------
|     |     |   8 |   4 |
-------------------------
|     |  16 |  16 |  32 |
-------------------------
|     |  16 |  32 | 128 |
-------------------------`

	board := [][]int64{
		{0, 0, 0, 2},
		{0, 0, 8, 4},
		{0, 16, 16, 32},
		{0, 16, 32, 128},
	}
	actual := BoardToString(board)

	assert.Equal(t, expected, actual)
}

func TestMoveRightWithEmptyReturnsEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.Equal(t, initialBoard, MoveRight(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveRight(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveRight(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveRight(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveRight(initialBoard))
}

func TestMoveRightMultipleRows(t *testing.T) {
	initialBoard := [][]int64{
		{2, 128, 0, 0},
		{4, 0, 0, 4},
		{2, 0, 8, 8},
		{2, 2, 0, 16},
	}

	expectedBoard := [][]int64{
		{0, 0, 2, 128},
		{0, 0, 0, 8},
		{0, 0, 2, 16},
		{0, 0, 4, 16},
	}

	assert.Equal(t, expectedBoard, MoveRight(initialBoard))
}

func TestMoveLeftWithEmptyReturnsEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.Equal(t, initialBoard, MoveLeft(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveLeft(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveLeft(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveLeft(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveLeft(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveLeft(initialBoard))
}

func TestMoveDownWithEmptyReturnsEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.Equal(t, initialBoard, MoveDown(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveDown(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveDown(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveDown(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveDown(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveDown(initialBoard))
}

func TestMoveUpWithEmptyReturnsEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.Equal(t, initialBoard, MoveUp(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveUp(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveUp(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveUp(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveUp(initialBoard))
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

	assert.Equal(t, expectedBoard, MoveUp(initialBoard))
}

func TestBoardIsFullReturnsFalseForEmptyBoard(t *testing.T) {
	board := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.False(t, BoardIsFull(board))
}

func TestBoardIsFullReturnsTrueForFullBoard(t *testing.T) {
	board := [][]int64{
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
	}

	assert.True(t, BoardIsFull(board))
}

func TestBoardIsFullReturnsFalseForHalfFullBoard(t *testing.T) {
	board := [][]int64{
		{2, 0, 2, 2},
		{2, 0, 2, 2},
		{0, 0, 0, 2},
		{0, 0, 0, 2},
	}

	assert.False(t, BoardIsFull(board))
}

