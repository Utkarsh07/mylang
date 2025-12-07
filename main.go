package main

import (
	"fmt"
	"mylang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n")
	fmt.Printf("╔═══════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║                                                           ║\n")
	fmt.Printf("║          Welcome to MyLang Programming Language!          ║\n")
	fmt.Printf("║                                                           ║\n")
	fmt.Printf("║         Feel free to type in commands and explore!        ║\n")
	fmt.Printf("║           Type 'exit' or press Ctrl+C to quit.            ║\n")
	fmt.Printf("║                                                           ║\n")
	fmt.Printf("╚═══════════════════════════════════════════════════════════╝\n")
	fmt.Printf("\n")
	fmt.Printf("Hello, %s", user.Username)
	fmt.Printf("\n")
	repl.Start(os.Stdin, os.Stdout)
}
