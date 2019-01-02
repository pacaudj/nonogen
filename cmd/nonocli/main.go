package nonocli

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	Flags = []cli.Flag{
		cli.Int64Flag{
			Name:  "size, s",
			Usage: "Size of the nonogram",
		},
		cli.Int64Flag{
			Name:  "brightness, b",
			Usage: "Brightness of the nonogram",
		},
		cli.StringFlag{
			Name:  "path, p",
			Usage: "Path to the input image",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "Path to the output image",
		},
	}
	Commands = []cli.Command{
		{
			Name:  "generate",
			Usage: "Generates a nonogram",
			Action: func(c *cli.Context) error {
				fmt.Println("Generate")
				return nil
			},
		},
		{
			Name:  "solve",
			Usage: "Solves a given nonogram.",
			Action: func(c *cli.Context) error {
				fmt.Println("Solve")
				return nil
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Flags = Flags
	app.Commands = Commands
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
