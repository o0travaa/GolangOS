package main

import (
	"GolangOS/commands"
	"fmt"
)

func main() {
	shell := shellSup()
	userName := user()
	var input string
	var flags string
	for {
		fmt.Printf("%s@%s", userName, shell)
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
