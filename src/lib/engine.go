package lib

import (
	"fmt"
	"time"
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
		fmt.Scan(&input)
		if input == "s" {
			Clear()
			PrintColor("Starting game...", "White")
			time.Sleep(1 * time.Second)
			Game(lists_words)
		} else if input == "q" {
			Clear()
			PrintColor("Thanks for Playing !", "White")
			keep_playing = false
		} else {
			invalid_ouput = true
		}
	}
}
