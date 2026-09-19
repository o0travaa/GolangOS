package main

import (
	"GolangOS/commands"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func sysClear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	shell := ShellSup()
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
			sysClear()
		case "echo":
			fmt.Println(flags)
		default:
			fmt.Println("Error: Command not found")
		}
	}
}
