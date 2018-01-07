package rating

import (
	"testing"

	"github.com/donutmonger/2048/board"
	"github.com/stretchr/testify/assert"
)

func TestGetRatingMaximizeEmptyReturns0ForFull(t *testing.T) {
	b := board.CompressBoardGrid([][]int64{
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
	})

	assert.Equal(t, uint64(0), GetRatingMaximizeEmpty(b))
}

func TestGetRatingMaximizeEmptyReturns16ForEmpty(t *testing.T) {
	b := board.CompressBoardGrid([][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	})

	assert.Equal(t, uint64(16), GetRatingMaximizeEmpty(b))
}

func TestGetRatingMaximizeEmptyReturns7ForPartiallyFull(t *testing.T) {
	b := board.CompressBoardGrid([][]int64{
		{2, 0, 0, 0},
		{2, 8, 4, 0},
		{2, 2, 2, 0},
		{0, 8, 2, 0},
	})

	assert.Equal(t, uint64(7), GetRatingMaximizeEmpty(b))
}
