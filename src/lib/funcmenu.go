package lib

import (
	"fmt"
	"strings"
	"time"
)

/*
Displays Menu to choose the diplay Mod and start the game
*/

func MenuMode(save Save) {
	save.TemplatesNames = Scandir("../Templates/policies/") // Get the map of all templates policies as value and an index as key
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
		keys := Listmap(save.TemplatesNames) // Get the key list in order to check if the next input is valid
		PrintColor("[b]: Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		fmt.Scanln(&input) // Scan and store the user input into the variable input
		input = strings.ToLower(input)
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
			save.DisplayMod = save.TemplatesNames[input]
			Game(save) // Call the function Game
		}
	}
}

/*
Displays Menu to choose a dictionnary
*/

func MenuDic(save Save) {
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
		input = strings.ToLower(input)
		if input == "b" {
			loop = false
			Clear()
		} else if !IsIn(keys, input) {
			invalid_ouput = true
			Clear()
		} else {
			Clear()
			save.ListsWords = GetFileInLine("../dictionnaries/" + dictionnaries_names[input] + ".txt")
			MenuMode(save)
		}
	}
}

/*
 */

func MenuSave() {
	Clear()
	var input string // Store the input player
	keep_playing := true
	invalid_ouput := false
	for keep_playing {
		if invalid_ouput {
			Clear()
			PrintColor("Invalid output !\n\n", "White")
			invalid_ouput = false
		} else {
			PrintColor("Hum it seems that there is a save, would you like to load it? Y/n: ", "white")
		}
		fmt.Scanln(&input) // get the input player
		input = strings.ToLower(input)
		if input == "y" {
			Game(LoadSave())
		} else if input == "n" {
			Engine()
		} else {
			invalid_ouput = true
		}
	}
}
