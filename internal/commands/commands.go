package commands

import (
	"github.com/pniedzwiedzinski/photoprism-cli/internal/login"
	"github.com/pniedzwiedzinski/photoprism-cli/internal/upload"
	"github.com/urfave/cli/v2"
)

// Commands - All cli commands
var Commands = []*cli.Command{
	{
		Name:   "login",
		Action: login.Command,
	},
	{
		Name:   "upload",
		Action: upload.Command,
	},
}
