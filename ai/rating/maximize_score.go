package rating

import "math"

func GetRatingMaximizeScore(b [][]int64) (totalScore uint64) {
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
