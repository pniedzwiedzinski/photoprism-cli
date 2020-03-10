package utils

import (
	"os/user"
)

// GetTokenLocation - get path to `.photoprism` folder
func GetTokenLocation() string {
	myself, err := user.Current()
	if err != nil {
		panic(err)
	}
	homedir := myself.HomeDir
	return homedir + "/.photoprism/"
}
