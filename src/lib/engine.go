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
		PrintColor("[q] quit\n[s] start new game\n", "White") // change print
		PrintColor("Choose: ", "White")
		fmt.Scanln(&input)
		if input == "s" {
			Clear()
			MenuMod(lists_words)
		} else if input == "q" {
			Clear()
			PrintColor("Thanks for Playing !\n", "White")
			keep_playing = false
		} else {
			invalid_ouput = true
		}
	}
}
