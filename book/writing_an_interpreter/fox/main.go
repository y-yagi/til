package main

import (
	"fmt"
	"fox/repl"
	"os"
)

func main() {
	fmt.Printf("This is fox\n")
	fmt.Printf("type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
