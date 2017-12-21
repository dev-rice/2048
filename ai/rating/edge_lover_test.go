package rating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRatingEdgeLoverLeft(t *testing.T) {
	bEdge := [][]int64{
		{0, 0, 0, 0},
		{8, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b := [][]int64{
		{0, 0, 0, 0},
		{0, 8, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.True(t, GetRatingEdgeLover(bEdge) > GetRatingEdgeLover(b))
}

func TestGetRatingEdgeLoverRight(t *testing.T) {
	bEdge := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 8},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b := [][]int64{
		{0, 0, 0, 0},
		{0, 8, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.True(t, GetRatingEdgeLover(bEdge) > GetRatingEdgeLover(b))
}

func TestGetRatingEdgeLoverBottom(t *testing.T) {
	bEdge := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 8, 0},
	}

	b := [][]int64{
		{0, 0, 0, 0},
		{0, 8, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.True(t, GetRatingEdgeLover(bEdge) > GetRatingEdgeLover(b))
}

func TestGetRatingEdgeLoverTop(t *testing.T) {
	bEdge := [][]int64{
		{0, 0, 8, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b := [][]int64{
		{0, 0, 0, 0},
		{0, 8, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.True(t, GetRatingEdgeLover(bEdge) > GetRatingEdgeLover(b))
}

func TestGetRatingEdgeLoverRatesHigherNumbersHigher(t *testing.T) {
	bBigEdge := [][]int64{
		{32, 0, 32, 0},
		{0, 0, 0, 0},
		{16, 0, 0, 0},
		{0, 8, 8, 0},
	}

	bLittleEdge := [][]int64{
		{2, 0, 2, 0},
		{0, 0, 64, 0},
		{8, 32, 0, 0},
		{0, 4, 4, 0},
	}

	assert.True(t, GetRatingEdgeLover(bBigEdge) > GetRatingEdgeLover(bLittleEdge))
}
