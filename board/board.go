package board

import (
	"container/list"
	"math/rand"

	"errors"
	"reflect"

	"time"
)

func NewEmptyBoard() [][]int64 {
	rand.Seed(time.Now().UnixNano())
	return [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
}

// No tests for this OH GODDDD!
func PlaceRandomTile(board [][]int64) [][]int64 {
	if boardIsFull(board) {
		return board
	}

	tileNumber := int64(2)
	if rand.Intn(10) == 0 {
		tileNumber = 4
	}

	for {
		size := len(board)
		x := rand.Intn(size)
		y := rand.Intn(size)
		if board[x][y] == 0 {
			board[x][y] = tileNumber
			return board
		}
	}
}

func boardIsFull(board [][]int64) bool {
	width := len(board)
	height := width

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if board[x][y] == 0 {
				return false
			}
		}
	}
	return true
}

// Ugly but way faster than previous version
//pkg: github.com/donutmonger/2048/board
//BenchmarkAreMovesLeft-8             	50000000	        36.4 ns/op
//BenchmarkAreMovesLeftCompressed-8   	1000000000	         2.28 ns/op
func AreMovesLeft(boardCompressed int64) bool {
	for rowNum := 0; rowNum < 4; rowNum++ {
		rowShift := uint((3 - rowNum) * 16)

		row := boardCompressed >> uint(rowShift) & 0xffff

		// Check for any zeros
		if (row&0xf000 == 0) || (row&0x0f00 == 0) || (row&0x00f0 == 0) || (row&0x000f == 0) {
			return true
		}

		// Check if next tile in row is equal to current
		// (Checking if any tile pairs match)
		tilePair0 := row >> uint(8)
		tilePair1 := row >> uint(4) & 0xff
		tilePair2 := row & 0xff

		if (tilePair0&0xf0 == tilePair0&0x0f) || (tilePair1&0xf0 == tilePair1&0x0f) || (tilePair2&0xf0 == tilePair2&0x0f) {
			return true
		}

		// Check if tile in next row is equal (for all but last row)
		if rowNum < 3 {
			nextRowShift := uint((3 - (rowNum + 1)) * 16)
			nextRow := boardCompressed << uint(nextRowShift) & 0xffff
			if (row&0xf000 == nextRow&0xf000) || (row&0x0f00 == nextRow&0x0f00) || (row&0x00f0 == nextRow&0x00f0) || (row&0x000f == nextRow&0x000f) {
				return true
			}
		}
	}

	return false
}

func CompressBoardGrid(board [][]int64) int64 {
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

// Only works for int64 that is a power of 2
func efficientLog2(v int64) (count uint8) {
	count = 1
	for v > 2 {
		v = v >> 1
		count++
	}
	return count
}

func UncompressBoard(compressedBoard int64) [][]int64 {
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
		return compressedBoard, 0, errors.New("No move was made")
	}

	// push right (don't merge) step
	var movedBoard int64
	var scoreAdd int64
	for rowNum := 0; rowNum < 4; rowNum++ {
		rowShift := uint((3 - rowNum) * 16)
		row := compressedBoard >> uint(rowShift) & 0xffff

		if row == 0x0000 {
			continue
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

		//Combine pairs
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

		movedBoard = movedBoard | (row << uint(rowShift))
	}

	if movedBoard == compressedBoard {
		return movedBoard, 0, errors.New("No move was made")
	}

	return movedBoard, scoreAdd, nil
}

func MoveLeft(board [][]int64) ([][]int64, int64, error) {
	outputBoard := make([][]int64, len(board))
	var score int64
	for y := 0; y < len(board); y++ {
		var scoreAdd int64
		outputBoard[y], scoreAdd = moveRowLeft(board[y])
		score += scoreAdd
	}

	if reflect.DeepEqual(board, outputBoard) {
		return outputBoard, 0, errors.New("No move was made")
	}

	return outputBoard, score, nil
}

func MoveDown(board [][]int64) ([][]int64, int64, error) {
	outputBoard := make([][]int64, len(board))
	var score int64
	for y := 0; y < len(board); y++ {
		outputBoard[y] = make([]int64, len(board))
	}

	for x := 0; x < len(board[0]); x++ {
		col := make([]int64, 0)
		for y := 0; y < len(board); y++ {
			col = append(col, board[y][x])
		}
		col, scoreAdd := moveRowRight(col)
		score += scoreAdd
		for y := 0; y < len(board); y++ {
			outputBoard[y][x] = col[y]
		}
	}

	if reflect.DeepEqual(board, outputBoard) {
		return outputBoard, 0, errors.New("No move was made")
	}

	return outputBoard, score, nil
}

func MoveUp(board [][]int64) ([][]int64, int64, error) {
	outputBoard := make([][]int64, len(board))
	var score int64
	for y := 0; y < len(board); y++ {
		outputBoard[y] = make([]int64, len(board))
	}

	for x := 0; x < len(board[0]); x++ {
		col := make([]int64, 0)
		for y := 0; y < len(board); y++ {
			col = append(col, board[y][x])
		}
		col, scoreAdd := moveRowLeft(col)
		score += scoreAdd
		for y := 0; y < len(board); y++ {
			outputBoard[y][x] = col[y]
		}
	}

	if reflect.DeepEqual(board, outputBoard) {
		return outputBoard, 0, errors.New("No move was made")
	}

	return outputBoard, score, nil
}

func moveRowRight(row []int64) ([]int64, int64) {
	rowList := sliceToList(row)

	score := int64(0)

	// Remove all zeros and put them at the front
	current := rowList.Front()
	for current != nil {
		next := current.Next()
		if current.Value == int64(0) {
			rowList.Remove(current)
			rowList.PushFront(int64(0))
		}
		current = next
	}

	// Merge non-zero pairs together
	current = rowList.Back()
	for current != nil && current.Prev() != nil {
		prev := current.Prev()
		prevPrev := prev.Prev()
		if current.Value == prev.Value && current.Value != 0 {
			// prev.value becomes current.value + prev.value
			prev.Value = current.Value.(int64) + prev.Value.(int64)

			// current.value becomes 0
			current.Value = int64(0)

			// current becomes prevPrev
			current = prevPrev

			// add to score
			score += prev.Value.(int64)
		} else {
			current = prev
		}
	}

	// Remove all zeros and put them at the front
	current = rowList.Front()
	for current != nil {
		next := current.Next()
		if current.Value == int64(0) {
			rowList.Remove(current)
			rowList.PushFront(int64(0))
		}
		current = next
	}

	return listToSlice(rowList), score
}

func moveRowLeft(row []int64) ([]int64, int64) {
	rowList := sliceToList(row)

	score := int64(0)

	// Remove all zeros and put them at the back
	current := rowList.Back()
	for current != nil {
		prev := current.Prev()
		if current.Value == int64(0) {
			rowList.Remove(current)
			rowList.PushBack(int64(0))
		}
		current = prev
	}

	// Merge non-zero pairs together
	current = rowList.Front()
	for current != nil && current.Next() != nil {
		next := current.Next()
		nextNext := next.Next()
		if current.Value == next.Value && current.Value != 0 {
			// next.value becomes current.value + next.value
			next.Value = current.Value.(int64) + next.Value.(int64)

			// current.value becomes 0
			current.Value = int64(0)

			// current becomes nextNext
			current = nextNext

			// add to score
			score += next.Value.(int64)
		} else {
			current = next
		}
	}

	// Remove all zeros and put them at the back
	current = rowList.Back()
	for current != nil {
		prev := current.Prev()
		if current.Value == int64(0) {
			rowList.Remove(current)
			rowList.PushBack(int64(0))
		}
		current = prev
	}

	return listToSlice(rowList), score
}

func sliceToList(slice []int64) *list.List {
	l := list.New()

	for _, elem := range slice {
		l.PushBack(elem)
	}

	return l
}

func listToSlice(l *list.List) []int64 {
	slice := make([]int64, l.Len())
	current := l.Front()
	i := 0
	for current != nil {
		slice[i] = current.Value.(int64)
		i++
		current = current.Next()
	}
	return slice
}
