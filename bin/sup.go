package bin

import (
	"fmt"
)

var UserName string
var Shell string

func ShellSup() string {
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
		Shell = shell1
	case 2:
		Shell = shell2
	case 3:
		Shell = shell3
	default:
		fmt.Println("Error")
		return ShellSup()
	}
	SysClear()
	return Shell
}

func User() string {
	fmt.Println("Setting up your system:")
	fmt.Println("User")
	fmt.Print("Enter a name of user: ")
	fmt.Scanln(&UserName)
	fmt.Println("Successfully!")
	SysClear()
	return UserName
}
