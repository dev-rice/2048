package players

import (
	"fmt"

	"github.com/donutmonger/2048/actions"
)

type scanner interface {
	Scan() bool
	Text() string
}

type HumanPlayer struct {
	s scanner
}

func NewHumanPlayer(scanner scanner) *HumanPlayer {
	return &HumanPlayer{
		s: scanner,
	}
}

func (h HumanPlayer) GetAction(gameBoard [][]int64) actions.Action {
	fmt.Print("Enter move (w,a,s,d): ")

	h.s.Scan()
	input := h.s.Text()
	if input == "w" {
		return actions.MoveUp
	} else if input == "a" {
		return actions.MoveLeft
	} else if input == "s" {
		return actions.MoveDown
	} else if input == "d" {
		return actions.MoveRight
	} else if input == "q" {
		return actions.Quit
	}

	return actions.None
}
