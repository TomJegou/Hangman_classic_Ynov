package main

import (
	"hangman_classic/lib"
	"os"
)

/*
Function main that call the Engine function from the package lib
*/

func main() {
	_, err := os.ReadFile("saves/save.json")
	if err == nil {
		lib.MenuSave()
	} else {
		lib.Engine()
	}
}
