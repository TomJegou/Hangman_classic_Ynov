package lib

import (
	"fmt"
	"time"
)

/*
Displays Menu to choose the diplay Mod and start the game
*/

func MenuMode(lists_words []string) {
	templates_names := Scandir("../Templates/policies/")
	var input string
	loop := true
	invalid_ouput := false
	for loop {
		if invalid_ouput {
			PrintColor("Invalid input\n\n", "White")
			invalid_ouput = false
		}
		PrintColor("Choose your mode\n\n", "White")
		PrintColor("[0]: Classique\n", "White")
		Listmap(templates_names)
		PrintColor("[b]: Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		fmt.Scanln(&input)
		if input == "b" {
			loop = false
			Clear()
		} else if input != "0" && input != "1" && input != "2" && input != "3" { // Todo: change
			invalid_ouput = true
			Clear()
		} else {
			Clear()
			PrintColor("Starting game...", "White")
			time.Sleep(1 * time.Second)
			Game(lists_words, input, templates_names)
		}
	}
}
