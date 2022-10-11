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
		keys := Listmap(templates_names)
		PrintColor("[b]: Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		fmt.Scanln(&input)
		if input == "b" {
			loop = false
			Clear()
		} else if !IsIn(keys, input) && input != "0" {
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

func MenuDic() {
	dictionnaries_names := Scandir("../dictionnaries/")
	var input string
	loop := true
	invalid_ouput := false
	for loop {
		if invalid_ouput {
			PrintColor("Invalid input\n\n", "White")
			invalid_ouput = false
		}
		PrintColor("Choose your dictionnaries\n\n", "White")
		keys := Listmap(dictionnaries_names)
		PrintColor("[b]: Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		fmt.Scanln(&input)
		if input == "b" {
			loop = false
			Clear()
		} else if !IsIn(keys, input) {
			invalid_ouput = true
			Clear()
		} else {
			Clear()
			MenuMode(GetFileInLine("../dictionnaries/" + dictionnaries_names[input] + ".txt"))
		}
	}
}
