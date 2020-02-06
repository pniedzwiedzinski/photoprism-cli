package login

import (
	"fmt"
	"github.com/pniedzwiedzinski/photoprism-cli/internal/api"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

func getPassword() string {
	fmt.Print("Photoprism password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println("\n")
	if err != nil {
		fmt.Println("Password typed: " + string(bytePassword))
	}
	return string(bytePassword)
}

// LoginCommand
func Command(c *cli.Context) error {
	if c.NArg() == 0 {
		fmt.Println(fmt.Errorf("You need to pass server IP: photoprism-cli login [IP]"))
		return nil
	}
	ip := c.Args().Get(0)
	password := getPassword()
	body := fmt.Sprintf("{\"email\": \"admin\", \"password\": \"%s\"}", password)

	a := api.NewAPI(ip)
	resp := a.Post("session", body)
	fmt.Println(resp)
	return nil
}
