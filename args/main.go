package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "string",
				Value: "english",
				Usage: "language for the greeting",
			},
			&cli.StringFlag{
				Name:  "string1",
				Value: "english",
				Usage: "language for the greeting",
			},
			&cli.StringFlag{
				Name:  "string2",
				Value: "english",
				Usage: "language for the greeting",
			},
			&cli.BoolFlag{
				Name:  "bool",
				Value: false,
				Usage: "language for the greeting",
			},
			&cli.IntFlag{
				Name:  "int",
				Value: 0,
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Args len", c.NumFlags())
			fmt.Println(c.String("string"))
			fmt.Println(c.String("string1"))
			fmt.Println(c.String("string2"))
			fmt.Println(c.Bool("bool"))
			fmt.Println(c.Int("int"))
			return nil
		},
	}
	app.Run(os.Args)
}
