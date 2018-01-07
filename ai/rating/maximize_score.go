package rating

import (
	"math"

	"github.com/donutmonger/2048/board"
)

func GetRatingMaximizeScore(compressed int64) (totalScore uint64) {
	b := board.UncompressBoard(compressed)
	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b[0]); x++ {
			totalScore += getScoreForTile(b[y][x])
		}
	}
	return totalScore
}

func getScoreForTile(t int64) (score uint64) {
	power := math.Log2(float64(t))
	return uint64(t) * uint64(power-1)
}
