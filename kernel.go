package main

import (
	"GolangOS/bin"
	"fmt"
)

func main() {
	Shell := bin.ShellSup()
	UserName := bin.User()
	var input string
	var flags string
	for {
		fmt.Printf("%s@%s", UserName, Shell)
		fmt.Scanf("%s %s\n", &input, &flags)
		switch input {
		case "help", "h":
			bin.Help()
		case "sd", "shutdown", "exit":
			fmt.Println("Shutting down...")
			return
		case "clear", "c":
			bin.SysClear()
		case "echo":
			fmt.Println(flags)
		case "fetch":
			bin.Fetch()
		case "about":
			bin.About()
		default:
			fmt.Println("Error: Command not found")
		}
	}
}
