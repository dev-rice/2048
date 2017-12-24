package rating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRatingMaximizeScoreSimple(t *testing.T) {
	b := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 2, 4},
	}

	assert.Equal(t, uint64(4), GetRatingMaximizeScore(b))
}

func TestGetRatingMaximizeScore(t *testing.T) {
	b := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 8, 16},
	}

	assert.Equal(t, uint64(64), GetRatingMaximizeScore(b))
}

func TestGetRatingMaximizeScore2(t *testing.T) {
	b := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 2},
		{2, 0, 0, 8},
		{0, 4, 4, 16},
	}

	assert.Equal(t, uint64(72), GetRatingMaximizeScore(b))
}

func TestGetRatingMaximizeScore3(t *testing.T) {
	b := [][]int64{
		{2, 4, 16, 4},
		{32, 128, 32, 64},
		{4, 32, 256, 4},
		{2, 8, 16, 2},
	}

	assert.Equal(t, uint64(3392), GetRatingMaximizeScore(b))
}
