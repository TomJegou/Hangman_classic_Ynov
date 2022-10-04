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
			PrintColor("Invalid output !\n\n", "White")
			invalid_ouput = false
		} else {
			PrintColor("Welcome to Classic_Hangman made by jtom and rlouis !\n\n", "white")
		}
		fmt.Println("[q] quit\n[s] start new game") // change print
		fmt.Print("Choose: ")
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
