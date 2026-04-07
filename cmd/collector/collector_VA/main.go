package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: myself <command> [options]")
		PrintAvailableCommands()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		RunAddCommand()
	case "list":
		RunListCommand()
	case "view":
		RunViewCommand()
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		PrintAvailableCommands()
	}
}
