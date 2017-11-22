package main

import (
	"os"

	"github.com/donutmonger/2048/ai"
	"github.com/donutmonger/2048/game"
	"github.com/urfave/cli"
)

// Add a test color board (whenever cli is a thing)
func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "ai",
			Action: ai.Play,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name: "delay",
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
	g.Play()
}
