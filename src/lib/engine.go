package lib

import (
	"fmt"
)

func Engine(lists_words []string) {
	var input string
	for input != "q" {
		fmt.Println("[q] quit\n[s] start new game")
		fmt.Printf("Choose: ")
		fmt.Scan(&input)
		if input == "s" {
			Game(lists_words)
		}
	}
}
