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

func TestMoveRightWithEmptyReturnsEmptyAndError(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	board, err := MoveRight(initialBoard)
	assert.Equal(t, initialBoard, board)
	assert.NotNil(t, err)
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

	actualBoard, err := MoveRight(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveRight(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveRight(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveRight(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveRight(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
}

func TestMoveRightWithNoChangesReturnsError(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 2, 4},
		{0, 0, 4, 8},
		{0, 16, 2, 32},
	}

	actualBoard, err := MoveRight(initialBoard)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
}

func TestMoveLeftWithEmptyReturnsEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	actualBoard, err := MoveLeft(initialBoard)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
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

	actualBoard, err := MoveLeft(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveLeft(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveLeft(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveLeft(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveLeft(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
}

func TestMoveLeftWithNoChangesReturnsError(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{2, 4, 0, 0},
		{2, 8, 4, 8},
		{2, 16, 2, 32},
	}

	actualBoard, err := MoveLeft(initialBoard)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
}

func TestMoveDownWithEmptyReturnsEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	actualBoard, err := MoveDown(initialBoard)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
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

	actualBoard, err := MoveDown(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveDown(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveDown(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveDown(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveDown(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
}

func TestMoveDownWithNoChangesReturnsError(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{2, 4, 0, 0},
		{4, 8, 4, 8},
		{2, 16, 2, 32},
	}

	actualBoard, err := MoveDown(initialBoard)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
}

func TestMoveUpWithEmptyReturnsEmpty(t *testing.T) {
	initialBoard := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	actualBoard, err := MoveUp(initialBoard)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
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

	actualBoard, err := MoveUp(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveUp(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveUp(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveUp(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
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

	actualBoard, err := MoveUp(initialBoard)
	assert.Equal(t, expectedBoard, actualBoard)
	assert.Nil(t, err)
}

func TestMoveUpWithNoChangesReturnsError(t *testing.T) {
	initialBoard := [][]int64{
		{4, 16, 4, 2},
		{2, 8, 2, 4},
		{4, 4, 0, 8},
		{2, 0, 0, 32},
	}

	actualBoard, err := MoveUp(initialBoard)
	assert.Equal(t, initialBoard, actualBoard)
	assert.NotNil(t, err)
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
