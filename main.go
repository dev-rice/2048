package main

import (
	"os"

	"bufio"

	"time"

	"encoding/json"
	"fmt"

	"github.com/donutmonger/2048/ai"
	"github.com/donutmonger/2048/ai/rating"
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
			Action: aiPlay,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "delay",
					Value: 0,
				},
				cli.StringFlag{
					Name:  "strategy",
					Usage: "Changed the ai strategy. Choose from [maximizeEmpty, edgeLover, random]",
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
	stats := g.Play(players.NewHumanPlayer(bufio.NewScanner(os.Stdin)))

	statsJson, _ := json.MarshalIndent(stats, "", " ")
	fmt.Println(string(statsJson))
}

func aiPlay(ctx *cli.Context) error {
	g := game.New()

	strategy := ctx.String("strategy")

	var player players.Player
	if strategy == "maximizeEmpty" {
		delay := time.Duration(ctx.Int("delay")) * time.Millisecond
		t := ai.Traverser{
			GetRating: rating.GetRatingMaximizeEmpty,
			MaxDepth:  4,
		}
		player = players.NewAIPlayer(delay, t)
	} else if strategy == "edgeLover" {
		delay := time.Duration(ctx.Int("delay")) * time.Millisecond
		t := ai.Traverser{
			GetRating: rating.GetRatingEdgeLover,
			MaxDepth:  4,
		}
		player = players.NewAIPlayer(delay, t)
	} else if strategy == "random" {
		player = players.NewRandomPlayer()
	} else {
		return errors.Errorf("Unknown ai player type '%s'", strategy)
	}

	stats := g.Play(player)

	statsJson, _ := json.MarshalIndent(stats, "", " ")
	fmt.Println(string(statsJson))

	return nil
}
