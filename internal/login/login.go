package login

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

// GetPassword
func GetPassword() string {
	fmt.Print("Photoprism password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println("\n")
	if err != nil {
		fmt.Println("Password typed: " + string(bytePassword))
	}
	return string(bytePassword)
}
