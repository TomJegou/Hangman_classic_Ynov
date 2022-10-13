package lib

import (
	"os"
	"time"
)

/*
Displays Menu to choose the diplay Mod and start the game
*/

func MenuMode(save *Save, issave bool) {
	if issave {
		issave = false
		Game(save, false)
	}
	save.TemplatesNames = Scandir("../Templates/policies/") // Get the map of all templates policies as value and an index as key
	keys := Listmap(save.TemplatesNames)                    // Get the key list in order to check if the next input is valid
	var input string
	loop := true
	validOutput := true
	for loop {
		Clear()
		if !validOutput {
			PrintColor("Invalid input\n\n", "White")
		}
		PrintColor("Choose your mode\n\n", "White")
		PrintColor("[b]: Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		keys = append(keys, []string{"b"}...)
		input, validOutput = GetInputUser(keys)
		if !validOutput {
			continue
		} else if input == "b" {
			loop = false
		} else {
			Clear()
			PrintColor("Starting game...", "White")
			time.Sleep(1 * time.Second)
			save.DisplayMode = input
			Game(save, true) // Call the function Game
		}
	}
}

/*
Displays Menu to choose a dictionnary
*/

func MenuDic(save *Save, issave bool) {
	if issave {
		issave = false
		MenuMode(save, true)
	}
	save.DictionnaryNames = Scandir("../dictionnaries/")
	var input string
	loop := true
	validOutput := true
	for loop {
		Clear()
		if !validOutput {
			PrintColor("Invalid input\n\n", "White")
		}
		PrintColor("Choose your dictionnaries\n\n", "White")
		keys := Listmap(save.DictionnaryNames)
		PrintColor("[b]: Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		keys = append(keys, []string{"b"}...)
		input, validOutput = GetInputUser(keys)
		if !validOutput {
			continue
		} else if input == "b" {
			loop = false
		} else {
			save.ListsWords = GetFileInLine("../dictionnaries/" + save.DictionnaryNames[input] + ".txt")
			save.WordToGess = ChoseRandomWord(save.ListsWords)
			MenuMode(save, false)
		}
	}
}

/*
Menu to choose if it loads the save
*/

func MenuSave() {
	var input string // Store the input player
	loop := true
	validOutput := true
	for loop {
		Clear()
		if !validOutput {
			PrintColor("Invalid output !\n\n", "White")
		}
		PrintColor("Hum it seems that there is a save, would you like to load it? Y/n: ", "white")
		input, validOutput = GetInputUser([]string{"y", "n"})
		if !validOutput {
			continue
		} else if input == "y" {
			LoadingBar(25)
			Engine(true)
			return
		} else {
			Engine(false)
			return
		}
	}
}

/*Menu to at the endgame to choose to continue or not*/

func EndgameMenu(save *Save, found bool) {
	// Display endgame message
	if found {
		if save.NumberError > 0 {
			DisplayHangman(save.NumberError)
		}
		DisplayModLetter(save, "White", true, true)
		PrintColor("\nCongrat !\nYou've found the word\nThe word was: "+save.WordToGess+"\n\n", "Green")
	} else {
		DisplayWrongLetter(save.NumberError, save.MaxError)
		DisplayHangman(save.NumberError)
		DisplayModLetter(save, "White", true, true)
		PrintColor("\nYou didn't find the word !\nThe word was: "+save.WordToGess+"\n\n", "Red")
	}
	loop := true
	var input string
	var validOutput = true
	for loop {
		Clear()
		if !validOutput {
			PrintColor("Invalid output\n", "White")
		}
		PrintColor("[c]continue [q]quit [b]Back\n\n", "White")
		PrintColor("Choose: ", "White")
		input, validOutput = GetInputUser([]string{"c", "q", "b"})
		if !validOutput {
			continue
		} else if input == "c" {
			PrintColor("Starting new game...", "White")
			time.Sleep(1 * time.Second)
			ResetWordFromSave(save)
			Game(save, true)
		} else if input == "q" {
			Clear()
			PrintColor("Thanks for playing !", "White")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		} else {
			ResetWordFromSave(save)
			loop = false
		}
	}
}
