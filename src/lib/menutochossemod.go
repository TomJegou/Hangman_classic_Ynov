package lib

import (
	"fmt"
	"time"
)

/*
Displays Menu to choose the diplay Mod and start the game
*/

func MenuMod(lists_words []string) {
	var input string
	loop := true
	invalid_ouput := false
	for loop {
		if invalid_ouput {
			PrintColor("Invalid input\n\n", "White")
			invalid_ouput = false
		}
		PrintColor("Choose your mod\n\n", "White")
		PrintColor("[0]Classique\n", "Blue")
		PrintColor("[1]Ascii\n", "White")
		PrintColor("[2]Stick\n", "Purple")
		PrintColor("[3]ShadowAscii\n", "White")
		PrintColor("[4]Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		fmt.Scanln(&input)
		if input == "4" {
			loop = false
			Clear()
		} else if input != "0" && input != "1" && input != "2" && input != "3" {
			invalid_ouput = true
			Clear()
		} else {
			Clear()
			PrintColor("Starting game...", "White")
			time.Sleep(1 * time.Second)
			Game(lists_words, input)
		}
	}
}
