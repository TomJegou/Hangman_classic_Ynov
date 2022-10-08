package main

import (
	"hangman_classic/lib"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		dic := os.Args[1:]
		lib.Engine(lib.GetFileInLine("../dictionnaries/" + dic[0] + ".txt"))
	} else {
		lib.PrintColor("No dictionnary given\n", "Red")
	}
}
