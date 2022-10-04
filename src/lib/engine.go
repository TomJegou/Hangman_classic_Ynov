package lib

import (
	"fmt"
)

func Engine(lists_words []string) {
	Clear()
	var input string
	keep_playing := true
	invalid_ouput := false
	for keep_playing {
		if invalid_ouput {
			Clear()
			fmt.Print("Invalid output !\n\n")
			invalid_ouput = false
		} else {
			fmt.Print("Welcome to Classic_Hangman made by jtom and rlouis !\n\n")
		}
		fmt.Println("[q] quit\n[s] start new game")
		fmt.Printf("Choose: ")
		fmt.Scan(&input)
		if input == "s" {
			Game(lists_words)
		} else if input == "q" {
			Clear()
			fmt.Println("Thanks for Playing !")
			keep_playing = false
		} else {
			invalid_ouput = true
		}
	}
}
