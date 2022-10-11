package lib

import (
	"fmt"
	"time"
)

/*
Displays Menu to choose the diplay Mod and start the game
*/

func MenuMode(lists_words []string) {
	templates_names := Scandir("../Templates/policies/") // Get the map of all templates policies as value and an index as key
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
		keys := Listmap(templates_names) // Get the key list in order to check if the next input is valid
		PrintColor("[b]: Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		fmt.Scanln(&input) // Scan and store the user input into the variable input
		if input == "b" {
			loop = false
			Clear()
		} else if !IsIn(keys, input) && input != "0" { // check the input validity
			invalid_ouput = true
			Clear()
		} else {
			Clear()
			PrintColor("Starting game...", "White")
			time.Sleep(1 * time.Second)
			Game(lists_words, input, templates_names) // Call the function Game
		}
	}
}

/*
Displays Menu to choose a dictionnary
*/

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
