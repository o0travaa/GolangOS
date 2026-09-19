package main

import "fmt"

func ShellSup() string {
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
		return ShellSup()
	}
	sysClear()
	return shell
}
