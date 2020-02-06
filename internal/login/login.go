package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pniedzwiedzinski/photoprism-cli/internal/api"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"os/user"
	"syscall"
)

func parseResponse(resp string) (string, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(resp), &result)
	if err != nil {
		panic(err)
	}
	if result["token"] == nil {
		error := errors.New("couldn't parse response from the server")
		if result["error"] != nil {
			error = fmt.Errorf("%s", result["error"])
		}
		return "", error
	}
	token := fmt.Sprintf("%s", result["token"])
	return token, nil
}

func saveToken(token string) {
	myself, err := user.Current()
	if err != nil {
		panic(err)
	}
	homedir := myself.HomeDir
	dir := homedir + "/.photoprism/"
	file := dir + "token.txt"
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(file, []byte(token), 0660)
	if err != nil {
		panic(err)
	}
}

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
	token, err := parseResponse(resp)
	if err != nil {
		return err
	}
	saveToken(token)
	fmt.Println("Logged in successfully. Token stored in ~/.photoprism/token.txt")
	return nil
}
