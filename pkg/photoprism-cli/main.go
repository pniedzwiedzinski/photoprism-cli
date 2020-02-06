package main

import (
	"fmt"
	"github.com/pniedzwiedzinski/photoprism-cli/internal/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: commands.Commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("An error occurred:")
		log.Fatal(err)
	}
}
