package bin

import "fmt"

func Fetch() {
	fmt.Println("   ______      __                      ____  _____")
	fmt.Println("  / ____/___  / /___ _____  ____ _    / __ \\/ ___/")
	fmt.Println(" / / __/ __ \\/ / __ `/ __ \\/ __ `/   / / / /\\__ \\ ")
	fmt.Println("/ /_/ / /_/ / / /_/ / / / / /_/ /   / /_/ /___/ /")
	fmt.Println("\\____/\\____/_/\\__,_/_| |_/\\__, /    \\____//____/  ")
	fmt.Println("                         /____/ ")
}

func About() {
	fmt.Println()
	fmt.Printf("%s@golangos\n", UserName)
	fmt.Println("---------------------")
	fmt.Println("System: GolangOS x86_64")
	fmt.Println("Github: github.com/o0travaa/GolangOS")
	fmt.Println("Shell: GoShell")
	fmt.Println()
}
