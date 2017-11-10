package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
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

func TestBoardToString(t *testing.T) {
	expected :=
		`-------------------------
|   0 |   0 |   0 |   2 |
-------------------------
|   0 |   0 |   8 |   4 |
-------------------------
|   0 |  16 |  16 |  32 |
-------------------------
|   0 |  16 |  32 | 128 |
-------------------------`

	board := [][]int64{
		{0, 0, 0, 2},
		{0, 0, 8, 4},
		{0, 16, 16, 32},
		{0, 16, 32, 128},
	}

	actual :=  BoardToString(board)

	fmt.Println(expected)
	fmt.Println(actual)

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
