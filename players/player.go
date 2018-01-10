package players

import "github.com/donutmonger/2048/actions"

type Player interface {
	GetAction(gameBoard int64) actions.Action
}
