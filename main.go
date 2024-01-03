package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	fmt.Printf("Monkey Programming Language\n")
	repl.Start(os.Stdin, os.Stdout)
}
