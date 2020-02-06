package main

import (
	"fmt"
	"github.com/pniedzwiedzinski/photoprism-cli/internal/api"
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
				Action: loginCommand,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func loginCommand(c *cli.Context) error {
	if c.NArg() == 0 {
		fmt.Println(fmt.Errorf("You need to pass server IP: photoprism-cli login [IP]"))
		return nil
	}
	ip := c.Args().Get(0)
	password := login.GetPassword()
	body := fmt.Sprintf("{\"email\": \"admin\", \"password\": \"%s\"}", password)

	a := api.NewAPI(ip)
	resp := a.Post("session", body)
	fmt.Println(resp)
	return nil
}
