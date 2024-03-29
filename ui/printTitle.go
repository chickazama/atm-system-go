package ui

import (
	"fmt"
	"os"
	"os/exec"
)

func printTitle() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("=== ATM Management System ===")
	return nil
}
