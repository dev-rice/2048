package players

import (
	"time"

	"sort"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/stats"
)

type MinimizePlayer struct {
	delay time.Duration
}

func NewMinimizePlayer(delay time.Duration) MinimizePlayer {
	return MinimizePlayer{
		delay: delay,
	}
}

func (p MinimizePlayer) GetAction(gameBoard [][]int64) actions.Action {
	if p.delay != 0*time.Second {
		defer func() { time.Sleep(p.delay) }()
	}

	nTilesBase := getNumTiles(gameBoard)
	moveToTilesMap := getMoveToTilesMap(gameBoard)

	var lowestMove actions.Action
	lowestNumTiles := nTilesBase
	for _, move := range []actions.Action{actions.MoveUp, actions.MoveDown, actions.MoveLeft, actions.MoveRight} {
		numTiles := moveToTilesMap[move]
		if numTiles <= lowestNumTiles {
			lowestNumTiles = numTiles
			lowestMove = move
		}
	}
	return lowestMove
}

func getMoveToTilesMap(gameBoard [][]int64) map[actions.Action]int {
	lookAhead := 2

	upBoard, err := board.MoveUp(gameBoard, stats.NewScore())
	var nTilesUp int
	if err != nil {
		nTilesUp = getNumTiles(gameBoard) + 1
	} else {
		nTilesUp = getLowestPossibleNumberOfTilesNMoves(upBoard, lookAhead)
	}

	leftBoard, err := board.MoveLeft(gameBoard, stats.NewScore())
	var nTilesLeft int
	if err != nil {
		nTilesLeft = getNumTiles(gameBoard) + 1
	} else {
		nTilesLeft = getLowestPossibleNumberOfTilesNMoves(leftBoard, lookAhead)
	}

	rightBoard, err := board.MoveRight(gameBoard, stats.NewScore())
	var nTilesRight int
	if err != nil {
		nTilesRight = getNumTiles(gameBoard) + 1
	} else {
		nTilesRight = getLowestPossibleNumberOfTilesNMoves(rightBoard, lookAhead)
	}

	downBoard, err := board.MoveDown(gameBoard, stats.NewScore())
	var nTilesDown int
	if err != nil {
		nTilesDown = getNumTiles(gameBoard) + 1
	} else {
		nTilesDown = getLowestPossibleNumberOfTilesNMoves(downBoard, lookAhead)
	}

	return map[actions.Action]int{
		actions.MoveUp:    nTilesUp,
		actions.MoveLeft:  nTilesLeft,
		actions.MoveRight: nTilesRight,
		actions.MoveDown:  nTilesDown,
	}
}

func getLowestPossibleNumberOfTilesNMoves(gameBoard [][]int64, movesToGo int) int {
	if movesToGo == 0 {
		return getNumTiles(gameBoard)
	}

	lowest := getNumTiles(gameBoard)

	upBoard, err := board.MoveUp(gameBoard, stats.NewScore())
	var upScore int
	if err != nil {
		upScore = lowest + 1
	} else {
		upScore = getLowestPossibleNumberOfTilesNMoves(upBoard, movesToGo-1)
	}

	downBoard, err := board.MoveDown(gameBoard, stats.NewScore())
	var downScore int
	if err != nil {
		downScore = lowest + 1
	} else {
		downScore = getLowestPossibleNumberOfTilesNMoves(downBoard, movesToGo-1)
	}

	leftBoard, err := board.MoveLeft(gameBoard, stats.NewScore())
	var leftScore int
	if err != nil {
		leftScore = lowest + 1
	} else {
		leftScore = getLowestPossibleNumberOfTilesNMoves(leftBoard, movesToGo-1)
	}

	rightBoard, err := board.MoveRight(gameBoard, stats.NewScore())
	var rightScore int
	if err != nil {
		rightScore = lowest + 1
	} else {
		rightScore = getLowestPossibleNumberOfTilesNMoves(rightBoard, movesToGo-1)
	}

	scores := []int{upScore, downScore, leftScore, rightScore}
	sort.Ints(scores)

	best := scores[0]
	if best == upScore {
		return upScore
	} else if best == downScore {
		return downScore
	} else if best == leftScore {
		return leftScore
	} else if best == rightScore {
		return rightScore
	} else {
		return rightScore
	}
}

func getNumTiles(board [][]int64) int {
	nTiles := 0
	for x := 0; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			if board[x][y] != 0 {
				nTiles++
			}
		}
	}
	return nTiles
}
