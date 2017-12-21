package rating

func GetRatingMaximizeEmpty(board [][]int64) uint64 {
	emptyTiles := uint64(0)
	for x := 0; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			if board[x][y] == 0 {
				emptyTiles++
			}
		}
	}
	return emptyTiles
}
