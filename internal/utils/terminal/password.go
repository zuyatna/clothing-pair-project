package terminal

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"syscall"
)

func HidePassword(message string) ([]byte, error) {
	if term.IsTerminal(int(syscall.Stdin)) {
		fmt.Printf(message)
		passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return nil, err
		}
		fmt.Println()
		return passwordBytes, nil
	} else {
		fmt.Print("Enter password (input will be visible): ")
		reader := bufio.NewReader(os.Stdin)
		passwordBytes, err := reader.ReadBytes('\n')
		if err != nil {
			return nil, err
		}
		return passwordBytes[:len(passwordBytes)-1], nil
	}
}
