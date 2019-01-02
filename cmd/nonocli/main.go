package main

import (
	"fmt"
	"github.com/shadonovitch/nonogen/pkg"
	"github.com/shadonovitch/nonogen/pkg/draw"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	Commands = []cli.Command{
		{
			Name:  "generate",
			Usage: "Generates a nonogram",
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "size, s",
					Usage: "Size of the nonogram",
				},
				cli.Int64Flag{
					Name:  "brightness, b",
					Usage: "Brightness of the nonogram",
				},
				cli.StringFlag{
					Name:  "input, i",
					Usage: "Path to the input image",
				},
				cli.StringFlag{
					Name:  "output, o",
					Usage: "Path to the output image",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("Generate")
				nono, err := pkg.Nonogen(c.Int("size"), c.Int("brightness"), c.String("input"))
				if err != nil {
					return err
				}
				if c.String("output") != "" {
					draw.NonoToJPG(nono, c.String("output"))
				}
				return nil
			},
		},
		{
			Name:  "solve",
			Usage: "Solves a given nonogram.",
			Flags: []cli.Flag {
				cli.StringFlag {
					Name: "input, i",
					Usage: "Path to the input nonogram.",
				},
				cli.StringFlag{
					Name:  "output, o",
					Usage: "Path to the output image",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("Solve")
				return nil
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Commands = Commands
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
