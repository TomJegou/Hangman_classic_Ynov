package main

import (
	"hangman_classic/lib"
	"os"
	"strings"
)

/*
Function main that call the Engine function from the package lib
*/

func main() {
	executablepath, _ := os.Executable()
	t := strings.Split(executablepath, "/")
	srcpath := "/"
	for i := 0; i < len(t)-2; i++ {
		srcpath += t[i] + "/"
	}
	srcpath += "src"
	srcpath = srcpath[1:]
	os.Chdir(srcpath)
	_, err := os.ReadFile("saves/save.json")
	if err == nil {
		lib.MenuSave()
	} else {
		lib.Engine(false)
	}
}
