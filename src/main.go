package main

import (
	"bufio"
	"fmt"
	"hangman_classic/lib"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		dic := os.Args[1:]
		file, err := os.Open("../dictionnaries/" + dic[0] + ".txt")
		if err != nil {
			lib.Clear()
			fmt.Println("\033[31m", "This file doesn't exist")
		} else {
			all_word := []string{}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				all_word = append(all_word, scanner.Text())
			}
			lib.Engine(all_word)
		}
	} else {
		fmt.Println("\033[31m", "No dictionnary given")
	}
}
