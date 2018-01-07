package players

import (
	"testing"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/stretchr/testify/assert"
)

type mockScanner struct {
	text string
}

func (m mockScanner) Scan() bool {
	return true
}

func (m mockScanner) Text() string {
	return m.text
}

func TestGetActionPressWReturnsMoveUp(t *testing.T) {
	m := mockScanner{
		text: "w",
	}
	h := NewHumanPlayer(m)

	assert.Equal(t, actions.MoveUp, h.GetAction(board.NewEmptyBoard()))
}

func TestGetActionPressAReturnsMoveLeft(t *testing.T) {
	m := mockScanner{
		text: "a",
	}
	h := NewHumanPlayer(m)

	assert.Equal(t, actions.MoveLeft, h.GetAction(board.NewEmptyBoard()))

}

func TestGetActionWithSReturnsMoveDown(t *testing.T) {
	m := mockScanner{
		text: "s",
	}
	h := NewHumanPlayer(m)

	assert.Equal(t, actions.MoveDown, h.GetAction(board.NewEmptyBoard()))

}

func TestGetActionWithDReturnsMoveRight(t *testing.T) {
	m := mockScanner{
		text: "d",
	}
	h := NewHumanPlayer(m)

	assert.Equal(t, actions.MoveRight, h.GetAction(board.NewEmptyBoard()))

}

func TestGetActionWithQReturnsQuit(t *testing.T) {
	m := mockScanner{
		text: "q",
	}
	h := NewHumanPlayer(m)

	assert.Equal(t, actions.Quit, h.GetAction(board.NewEmptyBoard()))

}

func TestGetActionWithLReturnsNone(t *testing.T) {
	m := mockScanner{
		text: "l",
	}
	h := NewHumanPlayer(m)

	assert.Equal(t, actions.None, h.GetAction(board.NewEmptyBoard()))

}
