package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey-lang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hello, %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
