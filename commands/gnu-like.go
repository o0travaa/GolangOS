package commands

import "fmt"

func help() {
	fmt.Println("Avialable commands:")
	fmt.Println("help - show avialable commands")
	fmt.Println("shutdown/sd - power off GolangOS")
}
