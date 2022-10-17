package main

import (
	"hangman_classic/lib"
	"os"
	"strings"
)

/*
Function main that check if a save exists and call the Engine function from the package lib or the load save Menu if there is a save
*/
func main() {
	/*
		We want that the user can start the executable file whatever his current working directory is
	*/
	executablepath, _ := os.Executable()    // get the absolute path of the executable
	t := strings.Split(executablepath, "/") //split into a slice all the dir's name
	srcpath := "/"                          // this will be the path where we will redirect the program
	for i := 0; i < len(t)-2; i++ {
		srcpath += t[i] + "/"
	}
	srcpath += "src"
	srcpath = srcpath[1:]
	os.Chdir(srcpath)                        // change the current working directory
	_, err := os.ReadFile("saves/save.json") //check if a save exists
	if err == nil {                          // call the Menu asking for loading the save
		lib.SaveMenu()
	} else {
		lib.Engine(false) // if there is no save start the game
	}
}
