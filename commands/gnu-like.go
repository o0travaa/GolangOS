package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Help() {
	fmt.Println("Available commands:")
	fmt.Println("help/h - show avialable commands")
	fmt.Println("shutdown/sd/exit - shut down GolangOS")
	fmt.Println("clear/c - clear the terminal")
	fmt.Println("echo - output a string (echo Hello! -> Hello!)")
}

func SysClear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
