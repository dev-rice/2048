package board

import (
	"fmt"
	"reflect"
	"strings"
)

func NewEmptyBoard() [][]int64 {
	return [][]int64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
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
			formatString := fmt.Sprintf(" %%%dv |", longestNumDigits)
			output += fmt.Sprintf(formatString, val)
		}
		output += "\n"
	}
	output += strings.Repeat("-", lineLength)

	return output
}

func MoveRight(board [][]int64) [][]int64 {
	for y := 0; y < len(board); y++ {
		board[y] = moveRowRight(board[y])
	}
	return board
}

func moveRowRight(row []int64) []int64 {
	rowCopy := make([]int64, len(row))
	copy(rowCopy, row)

	for i := 0; i < len(rowCopy)-1; i++ {
		if rowCopy[i] != 0 && rowCopy[i+1] == 0 {
			rowCopy[i+1] = rowCopy[i]
			rowCopy[i] = 0
		} else if rowCopy[i] == rowCopy[i+1] {
			rowCopy[i+1] = rowCopy[i] + rowCopy[i+1]
			rowCopy[i] = 0
		}
	}

	if reflect.DeepEqual(row, rowCopy) {
		return row
	} else {
		return moveRowRight(rowCopy)
	}
}
