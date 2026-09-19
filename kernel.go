package main

import (
	"GolangOS/commands"
	"fmt"
)

func main() {
	shell := shellSup()
	var input string
	var flags string
	for {
		fmt.Print(shell)
		fmt.Scanf("%s %s\n", &input, &flags)
		switch input {
		case "help", "h":
			commands.Help()
		case "sd", "shutdown", "exit":
			fmt.Println("Shutting down...")
			return
		case "clear", "c":
			commands.SysClear()
		case "echo":
			fmt.Println(flags)
		default:
			fmt.Println("Error: Command not found")
		}
	}
}
