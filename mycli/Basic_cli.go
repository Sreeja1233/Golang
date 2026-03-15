package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage:mycli <command> [option]")
		return
	}
	command := args[1]

	switch command {
	case "hello":
		if len(args) >= 3 {
			fmt.Println("Hello,", args[2])
		} else {
			fmt.Println("Hello from mycli tool")
		}

	default:
		fmt.Println("Unknown command:command")
	}
}
