package main

import (
	"GolangOS/commands"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func CallClear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func shellSup() string {
	var shell string

	shell1 := "sh > "
	shell2 := "[shell]$ > "
	shell3 := "sh ~> "

	fmt.Println("Setting up your system:")

	fmt.Println("Shell")
	fmt.Printf("1) %s\n", shell1)
	fmt.Printf("2) %s\n", shell2)
	fmt.Printf("3) %s\n", shell3)
	fmt.Print("> ")

	var shInput int
	fmt.Scanln(&shInput)
	switch shInput {
	case 1:
		shell = shell1
	case 2:
		shell = shell2
	case 3:
		shell = shell3
	default:
		fmt.Println("Error")
		return shellSup()
	}
	CallClear()
	return shell
}

func main() {
	shell := shellSup()
	var input string
	for {
		fmt.Print(shell)
		fmt.Scanln(&input)
		switch input {
		case "help":
			commands.Help()
		}
	}
}
