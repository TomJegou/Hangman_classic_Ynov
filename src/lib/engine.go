package lib

import (
	"fmt"
	"strings"
)

/*
Function Engine is the main loop for the hangman game
*/

func Engine(issave bool) {
	Clear()
	var save *Save
	if issave {
		issave = false
		MenuDic(LoadSave(), true)
	} else {
		save = &Save{MaxError: 10, AttemptNumber: 0, Debug: false}
	}
	var input string // Store the input player
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
		PrintColor("[s]: Start new game", "Green")
		PrintColor("\n[q]: quit\n\n", "Red") // change print
		PrintColor("Choose: ", "White")
		fmt.Scanln(&input) // get the input player
		input = strings.ToLower(input)
		if input == "s" {
			Clear()
			MenuDic(save, false) // calls the menu function
		} else if input == "q" {
			Clear()
			PrintColor("Thanks for Playing !\n", "White")
			keep_playing = false
		} else {
			invalid_ouput = true
		}
	}
}
