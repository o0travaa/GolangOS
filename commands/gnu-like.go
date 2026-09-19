package commands

import "fmt"

func Help() {
	fmt.Println("Avialable commands:")
	fmt.Println("help/h - show avialable commands")
	fmt.Println("shutdown/sd/exit - shut down GolangOS")
	fmt.Println("clear/c - clear the terminal")
}
