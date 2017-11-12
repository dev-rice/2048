package board

import (
	"fmt"
	"strings"
	"github.com/fatih/color"
	"math"
)

type Stringer struct {
	board [][]int64
}

func NewStringer(board [][]int64) *Stringer {
	return &Stringer {
		board: board,
	}
}

func (s Stringer) String() string {
	longestNumDigits := 0
	for _, row := range s.board {
		for _, val := range row {
			length := len(fmt.Sprintf("%v", val))
			if length > longestNumDigits {
				longestNumDigits = length
			}
		}
	}

	output := ""

	lineLength := len(s.board[0])*(longestNumDigits+2) + len(s.board[0]) + 1
	for _, row := range s.board {
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