package main

import (
	"os"

	"bufio"

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
			Name: "ai",
			Action: func(ctx *cli.Context) error {
				g := game.New()

				aiType := ctx.String("type")

				var player players.Player
				if aiType == "greedyscore" {
					player = players.NewGreedyScorePlayer()
				} else if aiType == "random" {
					player = players.NewRandomPlayer()
				} else {
					return errors.Errorf("Unknown ai player type '%s'", aiType)
				}

				g.Play(player)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "type",
					Value: "random",
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

func play(ctx *cli.Context) {
	g := game.New()
	g.Play(players.NewHumanPlayer(bufio.NewScanner(os.Stdin)))
}
