package board

import (
	"math/rand"

	"errors"
)

type Board uint64

func NewEmptyBoard() int64 {
	return 0
}

func NewBoardFromGrid(board [][]int64) int64 {
	// compressedBoard board is int64 where each 4 bytes is a tile. Tile value is calculated as 2^(4 byte tile val)
	// it is filled horizontally then vertically strating from the top and moving right. Most significant 4 bits of compressedBoard are log(2,tile_0,0), second most significant 4 bits are log(2,tile_1,0), etc.

	size := len(board)

	var compressedBoard int64
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			tileValue := board[y][x]
			var compressedValue uint8
			if tileValue == 0 {
				compressedValue = 0
			} else {
				compressedValue = efficientLog2(tileValue)
			}
			shiftAmount := uint((size-1-x)*4 + (size-1-y)*16)
			compressedBoard = compressedBoard | (int64(compressedValue) << shiftAmount)
		}
	}
	return compressedBoard
}

// No tests for this OH GODDDD!
func PlaceRandomTile(compressed int64) int64 {

	if boardIsFull(compressed) {
		return compressed
	}

	tileNumber := int64(1)
	if rand.Intn(10) == 0 {
		tileNumber = 2
	}

	for {
		shiftAmount := uint(4 * rand.Intn(16))
		if (compressed>>shiftAmount)&0xf == 0 {
			mask := int64(0xf << shiftAmount)
			return compressed&^mask | tileNumber<<shiftAmount
		}
	}
}

func boardIsFull(board int64) bool {
	for i := 0; i < 16; i++ {
		shiftAmount := uint(4 * i)
		if (board>>shiftAmount)&0xf == 0 {
			return false
		}
	}
	return true
}

// Ugly but way faster than previous version
//pkg: github.com/donutmonger/2048/board
//BenchmarkAreMovesLeft-8             	50000000	        36.4 ns/op
//BenchmarkAreMovesLeftCompressed-8   	1000000000	         2.28 ns/op
func AreMovesLeft(board int64) bool {
	boardGrid := ExtractGridFromBoard(board)
	size := len(boardGrid)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			current := boardGrid[y][x]
			if current == 0 {
				return true
			}
			if y > 0 {
				above := boardGrid[y-1][x]
				if current == above {
					return true
				}
			}
			if y < size-1 {
				below := boardGrid[y+1][x]
				if current == below {
					return true
				}
			}
			if x > 0 {
				left := boardGrid[y][x-1]
				if current == left {
					return true
				}
			}
			if x < size-1 {
				right := boardGrid[y][x+1]
				if current == right {
					return true
				}
			}
		}
	}

	return false
	//for rowNum := 0; rowNum < 4; rowNum++ {
	//	rowShift := uint((3 - rowNum) * 16)
	//
	//	row := board >> uint(rowShift) & 0xffff
	//
	//	// Check for any zeros
	//	if (row&0xf000 == 0) || (row&0x0f00 == 0) || (row&0x00f0 == 0) || (row&0x000f == 0) {
	//		return true
	//	}
	//
	//	// Check if next tile in row is equal to current
	//	// (Checking if any tile pairs match)
	//	tilePair0 := row >> uint(8)
	//	tilePair1 := row >> uint(4) & 0xff
	//	tilePair2 := row & 0xff
	//
	//	if (tilePair0&0xf0 == tilePair0&0x0f) || (tilePair1&0xf0 == tilePair1&0x0f) || (tilePair2&0xf0 == tilePair2&0x0f) {
	//		return true
	//	}
	//
	//	// Check if tile in next row is equal (for all but last row)
	//	if rowNum < 3 {
	//		nextRowShift := uint((3 - (rowNum + 1)) * 16)
	//		nextRow := board << uint(nextRowShift) & 0xffff
	//		if (row&0xf000 == nextRow&0xf000) || (row&0x0f00 == nextRow&0x0f00) || (row&0x00f0 == nextRow&0x00f0) || (row&0x000f == nextRow&0x000f) {
	//			return true
	//		}
	//	}
	//}
	//
	//return false
}

// Only works for int64 that is a power of 2
func efficientLog2(v int64) (count uint8) {
	count = 1
	for v > 2 {
		v = v >> 1
		count++
	}
	return count
}

func ExtractGridFromBoard(compressedBoard int64) [][]int64 {
	boardGrid := [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			shiftAmount := uint((3-x)*4 + (3-y)*16)
			compressedValue := (compressedBoard >> shiftAmount) & 0xf
			if compressedValue == 0 {
				boardGrid[y][x] = 0
			} else {
				boardGrid[y][x] = int64(2 << uint(compressedValue-1))
			}
		}
	}

	return boardGrid
}

// Before optimization:
//   BenchmarkMoveRight-8           	 1000000	      2321 ns/op
// After optimization:
//   BenchmarkMoveRight-8   	         50000000	      27.8 ns/op
func MoveRight(compressedBoard int64) (int64, int64, error) {

	if compressedBoard == 0 {
		return compressedBoard, 0, errors.New("no move was made")
	}

	var movedBoard int64
	var scoreAdd int64
	for rowNum := 0; rowNum < 4; rowNum++ {
		rowShift := uint((3 - rowNum) * 16)
		row := compressedBoard >> uint(rowShift) & 0xffff

		var scoreAddRow int64
		row, scoreAddRow = moveRowRightC(row)

		scoreAdd += scoreAddRow
		movedBoard = movedBoard | (row << uint(rowShift))
	}

	if movedBoard == compressedBoard {
		return movedBoard, 0, errors.New("no move was made")
	}

	return movedBoard, scoreAdd, nil
}

func moveRowRightC(row int64) (int64, int64) {
	var scoreAdd int64
	if row == 0x0000 {
		return row, scoreAdd
	}

	// Remove zeros
	for row&0x000f == 0 {
		row = row >> 4
	}

	if row&0xfff0 != 0 {
		for row&0x00f0 == 0 {
			row = ((row & 0xfff0) >> 4) | row&0x000f
		}
	}

	if row&0xff00 != 0 {
		for row&0x0f00 == 0 {
			row = ((row & 0xff00) >> 4) | row&0x00ff
		}
	}

	//Combine pairs (go from right to left)
	tilePairRight := row & 0xff
	if tilePairRight&0xf0>>4 == tilePairRight&0x0f {
		newTilePairRight := tilePairRight&0x0f + 1
		row = row&0xff00 | newTilePairRight
		scoreAdd += 2 << uint(newTilePairRight-1)
	}

	tilePairMiddle := (row >> 4) & 0xff
	if tilePairMiddle != 0 && tilePairMiddle&0xf0>>4 == tilePairMiddle&0x0f {
		newTilePairMiddle := tilePairMiddle&0x0f + 1
		row = row&0xf00f | (newTilePairMiddle << 4)
		scoreAdd += 2 << uint(newTilePairMiddle-1)
	}

	tilePairLeft := row >> 8
	if tilePairLeft != 0 && tilePairLeft&0xf0>>4 == tilePairLeft&0x0f {
		newTilePairLeft := tilePairLeft&0x0f + 1
		row = row&0x00ff | (newTilePairLeft << 8)
		scoreAdd += 2 << uint(newTilePairLeft-1)
	}

	// Remove zeros
	for row&0x000f == 0 {
		row = row >> 4
	}

	if row&0xfff0 != 0 {
		for row&0x00f0 == 0 {
			row = ((row & 0xfff0) >> 4) | row&0x000f
		}
	}

	if row&0xff00 != 0 {
		for row&0x0f00 == 0 {
			row = ((row & 0xff00) >> 4) | row&0x00ff
		}
	}

	return row, scoreAdd
}

// Before Optimization
// BenchmarkMoveLeft-8       	  500000	      2582 ns/op
// BenchmarkMoveLeft-8       	50000000	      30.4 ns/op
func MoveLeft(board int64) (int64, int64, error) {
	if board == 0 {
		return board, 0, errors.New("no move was made")
	}

	var movedBoard int64
	var scoreAdd int64
	for rowNum := 0; rowNum < 4; rowNum++ {
		rowShift := uint((3 - rowNum) * 16)
		row := board >> uint(rowShift) & 0xffff

		var scoreAddRow int64
		row, scoreAddRow = moveRowLeftC(row)

		scoreAdd += scoreAddRow
		movedBoard = movedBoard | (row << uint(rowShift))
	}

	if movedBoard == board {
		return board, 0, errors.New("no move was made")
	}

	return movedBoard, scoreAdd, nil
}

func moveRowLeftC(row int64) (int64, int64) {
	var scoreAdd int64
	if row == 0x0000 {
		return row, 0
	}

	// Remove zeros
	for row&0xf000 == 0 {
		row = row << 4
	}

	if row&0x0fff != 0 {
		for row&0x0f00 == 0 {
			row = ((row & 0x0fff) << 4) | row&0xf000
		}
	}

	if row&0x00ff != 0 {
		for row&0x00f0 == 0 {
			row = ((row & 0x00ff) << 4) | row&0xff00
		}
	}

	//Combine pairs (go from left to right)
	tilePairLeft := row >> 8
	if tilePairLeft&0xf0>>4 == tilePairLeft&0x0f {
		newTilePairLeft := (tilePairLeft&0x0f + 1) << 4
		row = row&0x00ff | newTilePairLeft<<8
		scoreAdd += 2 << uint((newTilePairLeft>>4)-1)
	}

	tilePairMiddle := (row >> 4) & 0xff
	if tilePairMiddle != 0 && tilePairMiddle&0xf0>>4 == tilePairMiddle&0x0f {
		newTilePairMiddle := (tilePairMiddle&0x0f + 1) << 4
		row = row&0xf00f | (newTilePairMiddle << 4)
		scoreAdd += 2 << uint((newTilePairMiddle>>4)-1)
	}

	tilePairRight := row & 0xff
	if tilePairRight != 0 && tilePairRight&0xf0>>4 == tilePairRight&0x0f {
		newTilePairRight := (tilePairRight&0x0f + 1) << 4
		row = row&0xff00 | newTilePairRight
		scoreAdd += 2 << uint((newTilePairRight>>4)-1)
	}

	// Remove zeros
	for row&0xf000 == 0 {
		row = row << 4
	}

	if row&0x0fff != 0 {
		for row&0x0f00 == 0 {
			row = ((row & 0x0fff) << 4) | row&0xf000
		}
	}

	if row&0x00ff != 0 {
		for row&0x00f0 == 0 {
			row = ((row & 0x00ff) << 4) | row&0xff00
		}
	}

	return row, scoreAdd
}

// Before Optimization
// BenchmarkMoveDown-8       	 1000000	      2640 ns/op
// BenchmarkMoveDown-8       	50000000	        39.6 ns/op
func MoveDown(board int64) (int64, int64, error) {
	if board == 0 {
		return board, 0, errors.New("no move was made")
	}

	// transpose compressed board
	transposed := transposeCompressedBoard(board)

	var movedTransposed int64
	var scoreAdd int64
	for rowNum := 0; rowNum < 4; rowNum++ {
		rowShift := uint((3 - rowNum) * 16)
		row := transposed >> uint(rowShift) & 0xffff

		var scoreAddRow int64
		row, scoreAddRow = moveRowRightC(row)

		scoreAdd += scoreAddRow
		movedTransposed = movedTransposed | (row << uint(rowShift))
	}

	movedBoard := transposeCompressedBoard(movedTransposed)

	if movedBoard == board {
		return board, 0, errors.New("no move was made")
	}

	return movedBoard, scoreAdd, nil
}

func transposeCompressedBoard(b int64) (t int64) {
	t = (((b >> 60) & 0xf) << 60) | t
	t = (((b >> 56) & 0xf) << 44) | t
	t = (((b >> 52) & 0xf) << 28) | t
	t = (((b >> 48) & 0xf) << 12) | t

	t = (((b >> 44) & 0xf) << 56) | t
	t = (((b >> 40) & 0xf) << 40) | t
	t = (((b >> 36) & 0xf) << 24) | t
	t = (((b >> 32) & 0xf) << 8) | t

	t = (((b >> 28) & 0xf) << 52) | t
	t = (((b >> 24) & 0xf) << 36) | t
	t = (((b >> 20) & 0xf) << 20) | t
	t = (((b >> 16) & 0xf) << 4) | t

	t = (((b >> 12) & 0xf) << 48) | t
	t = (((b >> 8) & 0xf) << 32) | t
	t = (((b >> 4) & 0xf) << 16) | t
	t = (((b >> 0) & 0xf) << 0) | t

	return t
}

// Before Optimization
// BenchmarkMoveUp-8         	  500000	      2848 ns/op
// BenchmarkMoveUp-8         	30000000	        41.6 ns/op
func MoveUp(board int64) (int64, int64, error) {
	if board == 0 {
		return board, 0, errors.New("no move was made")
	}

	// transpose compressed board
	transposed := transposeCompressedBoard(board)

	var movedTransposed int64
	var scoreAdd int64
	for rowNum := 0; rowNum < 4; rowNum++ {
		rowShift := uint((3 - rowNum) * 16)
		row := transposed >> uint(rowShift) & 0xffff

		var scoreAddRow int64
		row, scoreAddRow = moveRowLeftC(row)

		scoreAdd += scoreAddRow
		movedTransposed = movedTransposed | (row << uint(rowShift))
	}

	movedBoard := transposeCompressedBoard(movedTransposed)

	if movedBoard == board {
		return board, 0, errors.New("no move was made")
	}

	return movedBoard, scoreAdd, nil
}
