package main

import (
	"fmt"
	"fox/repl"
	"os"
)

func main() {
	fmt.Printf("This is the Fox programming language!\n")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
