package commands

import (
	"github.com/pniedzwiedzinski/photoprism-cli/internal/login"
	"github.com/urfave/cli/v2"
)

// Commands - All cli commands
var Commands = []*cli.Command{
	{
		Name:   "login",
		Action: login.Command,
	},
}
