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
					Usage: "Changes the ai strategy. Choose from [maximize_empty, maximize_score, edge_lover, monte_carlo, random]",
					Value: "maximize_empty",
				},
				cli.BoolFlag{
					Name:  "silent",
					Usage: "Will not print out the game state while playing. Only the metrics will be output after the game is done.",
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

type printer interface {
	Printf(format string, v ...interface{})
	ClearScreen()
}

type stdoutPrinter struct{}

func (p stdoutPrinter) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (p stdoutPrinter) ClearScreen() {
	print("\033[H\033[2J")
}

type silentPrinter struct{}

func (p silentPrinter) Printf(format string, v ...interface{}) {
}

func (p silentPrinter) ClearScreen() {
}

func play(_ *cli.Context) {
	g := game.New()
	metrics := g.Play(players.NewHumanPlayer(bufio.NewScanner(os.Stdin)), stdoutPrinter{})
	metricsJson, _ := json.MarshalIndent(metrics, "", " ")
	fmt.Println(string(metricsJson))
}

func aiPlay(ctx *cli.Context) error {
	g := game.New()

	strategy := ctx.String("strategy")

	var p printer = stdoutPrinter{}
	if ctx.Bool("silent") {
		p = silentPrinter{}
	}

	var player players.Player
	delay := time.Duration(ctx.Int("delay")) * time.Millisecond
	if strategy == "maximize_empty" {
		t := ai.Traverser{
			GetRating: rating.GetRatingMaximizeEmpty,
			MaxDepth:  4,
		}
		player = players.NewAIPlayer(delay, t)
	} else if strategy == "edge_lover" {
		t := ai.Traverser{
			GetRating: rating.GetRatingEdgeLover,
			MaxDepth:  4,
		}
		player = players.NewAIPlayer(delay, t)
	} else if strategy == "maximize_score" {
		t := ai.Traverser{
			GetRating: rating.GetRatingMaximizeScore,
			MaxDepth:  4,
		}
		player = players.NewAIPlayer(delay, t)
	} else if strategy == "monte_carlo" {
		player = players.NewMonteCarloPlayer()
	} else if strategy == "random" {
		player = players.NewRandomPlayer()
	} else {
		return errors.Errorf("Unknown ai player type '%s'", strategy)
	}

	metrics := g.Play(player, p)
	metricsJson, _ := json.MarshalIndent(metrics, "", " ")
	fmt.Println(string(metricsJson))

	return nil
}
