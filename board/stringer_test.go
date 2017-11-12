package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringerStringForEmpty(t *testing.T) {
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
	s := NewStringer(board)

	assert.Equal(t, expected, s.String())
}

func TestStringerStringForPopulatedBoard(t *testing.T) {
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
	s := NewStringer(board)

	assert.Equal(t, expected, s.String())
}
