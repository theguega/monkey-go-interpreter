// main.go

package main

import (
	"fmt"
	"monkey-interpreter/src/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Hello %s! This is the Monkey 🐒 programming language!\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
