package rating

func GetRatingEdgeLover(b [][]int64) uint64 {
	rating := uint64(0)
	for x := 0; x < len(b); x++ {
		// Left Edge
		rating += uint64(b[x][0])

		// Right Edge
		rating += uint64(b[x][len(b)-1])
	}

	for y := 0; y < len(b[0]); y++ {
		// Top Edge
		rating += uint64(b[0][y])

		// Bottom Edge
		rating += uint64(b[len(b[0])-1][y])
	}

	return rating
}
