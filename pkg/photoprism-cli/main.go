package main

import (
	"fmt"
	"github.com/pniedzwiedzinski/photoprism-cli/internal/login"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "login",
				Action: login.Command,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("An error occurred:")
		log.Fatal(err)
	}
}
