package main

import (
	"os"

	"bufio"

	"time"

	"github.com/donutmonger/2048/boardtree"
	"github.com/donutmonger/2048/game"
	"github.com/donutmonger/2048/players"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

// Add a test color board (whenever cli is a thing)
func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "ai",
			Action: ai,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "delay",
					Value: 0,
				},
				cli.StringFlag{
					Name:  "strategy",
					Usage: "Changed the ai strategy. Choose from [maximizeEmpty, random]",
					Value: "maximizeEmpty",
				},
			},
		},
		{
			Name:   "play",
			Action: play,
		},
	}
	app.Run(os.Args)
}

func play(_ *cli.Context) {
	g := game.New()
	g.Play(players.NewHumanPlayer(bufio.NewScanner(os.Stdin)))
}

func ai(ctx *cli.Context) error {
	g := game.New()

	aiType := ctx.String("strategy")

	var player players.Player
	if aiType == "maximizeEmpty" {
		delay := time.Duration(ctx.Int("delay")) * time.Millisecond
		t := boardtree.Traverser{
			GetScore: getNumEmptyTiles,
			MaxDepth: 4,
		}
		player = players.NewAIPlayer(delay, t)
	} else if aiType == "random" {
		player = players.NewRandomPlayer()
	} else {
		return errors.Errorf("Unknown ai player type '%s'", aiType)
	}

	g.Play(player)
	return nil
}

func getNumEmptyTiles(board [][]int64) uint64 {
	emptyTiles := uint64(0)
	for x := 0; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			if board[x][y] == 0 {
				emptyTiles++
			}
		}
	}
	return emptyTiles
}
