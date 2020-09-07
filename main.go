package main

import (
	"log" // imports as package cli)
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	//here we start our CLI
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			//this is where we call our search engine to return results
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
