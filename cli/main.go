package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("There are no arguments")
	} else {
		for _, val := range os.Args[1:] {
			fmt.Println(val)
		}
	}
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			name := "Nefertiti"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	app.Run(os.Args)
}
