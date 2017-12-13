package actions

type Action int

const (
	MoveUp Action = iota
	MoveDown
	MoveLeft
	MoveRight
	None
	Quit
)
