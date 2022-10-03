package main

import (
	"bufio"
	"fmt"
	"hangman_classic/lib"
	"os"
)

func main() {
	dic := os.Args[1:]
	file, err := os.Open("../dictionnaries/" + dic[0] + ".txt")
	if err != nil {
		lib.Clear()
		fmt.Println(err)
	} else {
		all_word := []string{}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			all_word = append(all_word, scanner.Text())
		}
		lib.Engine(all_word)
	}
}
