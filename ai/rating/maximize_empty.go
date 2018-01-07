package rating

import "github.com/donutmonger/2048/board"

func GetRatingMaximizeEmpty(compressed int64) uint64 {
	b := board.UncompressBoard(compressed)
	emptyTiles := uint64(0)
	for x := 0; x < len(b[0]); x++ {
		for y := 0; y < len(b); y++ {
			if b[x][y] == 0 {
				emptyTiles++
			}
		}
	}
	return emptyTiles
}
