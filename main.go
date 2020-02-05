package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "login",
				Action: login,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getPassword() string {
	fmt.Println("Photoprism password: ")
	var password string
	fmt.Scanln(&password)
	return password
}

func login(c *cli.Context) error {
	if c.NArg() == 0 {
		fmt.Println(fmt.Errorf("You need to pass server IP: photoprism-cli login [IP]"))
		return nil
	}
	ip := c.Args().Get(0)
	password := getPassword()
	body, err := json.Marshal(map[string]string{
		"email":    "admin",
		"password": password,
	})

	apiURL := "http://" + ip + ":2342/api/v1/session"
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		fmt.Errorf(err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(b))
	return nil
}
