package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// https://github.com/urfave/cli

func main() {
	app := &cli.App{
		Name:  "hi-cli",
		Usage: "try urfave-cli",
		Action: func(c *cli.Context) error {
			fmt.Println("Hi https://github.com/urfave/cli")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
