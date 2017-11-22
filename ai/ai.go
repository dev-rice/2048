package ai

import (
	"fmt"

	"math/rand"

	"time"

	"sort"

	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/stats"
	"github.com/urfave/cli"
)

func Play(ctx *cli.Context) {
	rand.Seed(time.Now().UnixNano())

	shouldDelay := ctx.IsSet("delay")
	delay := ctx.Int("delay")

	gameBoard := board.NewEmptyBoard()
	board.PlaceRandomTile(gameBoard)
	board.PlaceRandomTile(gameBoard)

	score := stats.NewScore()

	didMove := false
	for board.AreMovesLeft(gameBoard) {
		if didMove {
			gameBoard = board.PlaceRandomTile(gameBoard)
			didMove = false
		}

		clearScreen()
		fmt.Printf("Score: %v\n", score.Get())
		fmt.Println(board.NewStringer(gameBoard).String() + "\n")

		var err error
		move := getMoveFuncGreedyScore(gameBoard)
		gameBoard, err = move(gameBoard, score)
		if err == nil {
			didMove = true
		}

		if shouldDelay {
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}

	}

	fmt.Println("There are no moves left, you lose!")
}

type moveFunc func([][]int64, *stats.Score) ([][]int64, error)

func getMoveFuncRandom() moveFunc {
	moveNum := rand.Intn(4)
	switch moveNum {
	case 0:
		return board.MoveUp
		break
	case 1:
		return board.MoveDown
		break
	case 2:
		return board.MoveLeft
		break
	case 3:
		return board.MoveRight
		break
	}

	return board.MoveUp
}

func getMoveFuncGreedyScore(gameBoard [][]int64) moveFunc {
	upScore := stats.NewScore()
	_, err := board.MoveUp(gameBoard, upScore)
	if err != nil {
		upScore = stats.NewScore()
		upScore.Add(-1)
	}

	downScore := stats.NewScore()
	_, err = board.MoveDown(gameBoard, downScore)
	if err != nil {
		downScore = stats.NewScore()
		downScore.Add(-1)
	}

	leftScore := stats.NewScore()
	_, err = board.MoveLeft(gameBoard, leftScore)
	if err != nil {
		leftScore = stats.NewScore()
		leftScore.Add(-1)
	}

	rightScore := stats.NewScore()
	_, err = board.MoveRight(gameBoard, rightScore)
	if err != nil {
		rightScore = stats.NewScore()
		rightScore.Add(-1)
	}

	scores := []int{int(upScore.Get()), int(downScore.Get()), int(leftScore.Get()), int(rightScore.Get())}
	sort.Ints(scores)

	best := scores[len(scores)-1]
	if best == int(upScore.Get()) {
		return board.MoveUp
	} else if best == int(downScore.Get()) {
		return board.MoveDown
	} else if best == int(leftScore.Get()) {
		return board.MoveLeft
	} else if best == int(rightScore.Get()) {
		return board.MoveRight
	}

	return board.MoveRight
}

func clearScreen() {
	print("\033[H\033[2J")
}
