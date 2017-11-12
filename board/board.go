package board

import (
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"strings"

	"errors"
	"reflect"

	"github.com/fatih/color"
)

func NewEmptyBoard() [][]int64 {
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

func AreMovesLeft(board [][]int64) bool {
	size := len(board)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			current := board[y][x]
			if current == 0 {
				return true
			}
			if y > 0 {
				above := board[y-1][x]
				if current == above {
					return true
				}
			}
			if y < size-1 {
				below := board[y+1][x]
				if current == below {
					return true
				}
			}
			if x > 0 {
				left := board[y][x-1]
				if current == left {
					return true
				}
			}
			if x < size-1 {
				right := board[y][x+1]
				if current == right {
					return true
				}
			}
		}
	}

	return false
}

func BoardToString(board [][]int64) string {
	longestNumDigits := 0
	for _, row := range board {
		for _, val := range row {
			length := len(fmt.Sprintf("%v", val))
			if length > longestNumDigits {
				longestNumDigits = length
			}
		}
	}

	output := ""

	lineLength := len(board[0])*(longestNumDigits+2) + len(board[0]) + 1
	for _, row := range board {
		output += strings.Repeat("-", lineLength)
		output += "\n|"
		for _, val := range row {
			if val == 0 {
				output += fmt.Sprintf(" %s |", strings.Repeat(" ", longestNumDigits))
			} else {
				c := color.New(colorForNumber(val))
				formatString := fmt.Sprintf(" %%%dv", longestNumDigits)
				output += c.Sprintf(formatString, val)
				output += " |"
			}
		}
		output += "\n"
	}
	output += strings.Repeat("-", lineLength)

	return output
}

func colorForNumber(n int64) color.Attribute {
	exp := int(math.Log2(float64(n)))
	exp = exp % 16

	if exp <= 8 {
		return color.Attribute(exp + int(color.FgBlack) + 1)
	} else {
		return color.Attribute(exp - 8 + int(color.FgHiBlack) + 1)
	}
}

func MoveRight(board [][]int64) ([][]int64, error) {
	outputBoard := make([][]int64, len(board))
	for y := 0; y < len(board); y++ {
		outputBoard[y] = moveRowRight(board[y])
	}

	if reflect.DeepEqual(board, outputBoard) {
		return outputBoard, errors.New("No move was made")
	}

	return outputBoard, nil
}

func MoveLeft(board [][]int64) ([][]int64, error) {
	outputBoard := make([][]int64, len(board))
	for y := 0; y < len(board); y++ {
		outputBoard[y] = moveRowLeft(board[y])
	}

	if reflect.DeepEqual(board, outputBoard) {
		return outputBoard, errors.New("No move was made")
	}

	return outputBoard, nil
}

func MoveDown(board [][]int64) ([][]int64, error) {
	outputBoard := make([][]int64, len(board))
	for y := 0; y < len(board); y++ {
		outputBoard[y] = make([]int64, len(board))
	}

	for x := 0; x < len(board[0]); x++ {
		col := make([]int64, 0)
		for y := 0; y < len(board); y++ {
			col = append(col, board[y][x])
		}
		col = moveRowRight(col)
		for y := 0; y < len(board); y++ {
			outputBoard[y][x] = col[y]
		}
	}

	if reflect.DeepEqual(board, outputBoard) {
		return outputBoard, errors.New("No move was made")
	}

	return outputBoard, nil
}

func MoveUp(board [][]int64) ([][]int64, error) {
	outputBoard := make([][]int64, len(board))
	for y := 0; y < len(board); y++ {
		outputBoard[y] = make([]int64, len(board))
	}

	for x := 0; x < len(board[0]); x++ {
		col := make([]int64, 0)
		for y := 0; y < len(board); y++ {
			col = append(col, board[y][x])
		}
		col = moveRowLeft(col)
		for y := 0; y < len(board); y++ {
			outputBoard[y][x] = col[y]
		}
	}

	if reflect.DeepEqual(board, outputBoard) {
		return outputBoard, errors.New("No move was made")
	}

	return outputBoard, nil
}

func moveRowRight(row []int64) []int64 {
	rowList := sliceToList(row)

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

	return listToSlice(rowList)
}

func moveRowLeft(row []int64) []int64 {
	rowList := sliceToList(row)

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

	return listToSlice(rowList)
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
