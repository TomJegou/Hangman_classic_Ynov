package lib

import (
	"os"
	"time"
)

/*
Displays Menu to choose the diplay Mod and start the game
*/
func ModeMenu(save *Save, issave bool) {
	if issave {
		issave = false
		Game(save, false)
	}
	save.TemplatesNames = ScanDir("../Templates/policies/") // Get the map of all templates policies as value and an index as key
	var input string
	loop := true
	validOutput := true
	for loop {
		ClearConsole()
		if !validOutput {
			PrintColor("Invalid input\n\n", "White")
		}
		PrintColor("Choose your mode\n\n", "White")
		keys := ListMap(save.TemplatesNames) // Get the key list in order to check if the next input is valid
		PrintColor("[b]: Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		keys = append(keys, []string{"b"}...)
		input, validOutput = GetInputUser(keys)
		if !validOutput {
			continue
		} else if input == "b" {
			loop = false
		} else {
			ClearConsole()
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
func DicMenu(save *Save, issave bool) {
	if issave {
		issave = false
		ModeMenu(save, true)
	}
	save.DictionnaryNames = ScanDir("../dictionnaries/")
	var input string
	loop := true
	validOutput := true
	for loop {
		ClearConsole()
		if !validOutput {
			PrintColor("Invalid input\n\n", "White")
		}
		PrintColor("Choose your dictionnaries\n\n", "White")
		keys := ListMap(save.DictionnaryNames)
		PrintColor("[b]: Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		keys = append(keys, []string{"b"}...)
		input, validOutput = GetInputUser(keys)
		if !validOutput {
			continue
		} else if input == "b" {
			loop = false
		} else {
			save.ListsWords = GetFileLineInSlice("../dictionnaries/" + save.DictionnaryNames[input] + ".txt")
			save.WordToGess = ChoseRandomWord(save.ListsWords)
			ModeMenu(save, false)
		}
	}
}

/*
Menu to choose if it loads the save
*/
func SaveMenu() {
	var input string // Store the input player
	loop := true
	validOutput := true
	for loop {
		ClearConsole()
		if !validOutput {
			PrintColor("Invalid input !\n\n", "White")
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

/*
Menu to at the endgame to choose to continue or not
*/
func EndgameMenu(save *Save, found bool) {
	// Display endgame message
	if found {
		ClearConsole()
		if save.NumberError > 0 {
			DisplayHangman(save.NumberError)
		}
		DisplayModLetter(save, "White", true, true)
		PrintColor("\n\nCongrat !\nYou've found the word\nThe word was: "+save.WordToGess+"\n\n", "Green")
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
		if !validOutput {
			ClearConsole()
			PrintColor("Invalid input\n", "White")
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
			ClearConsole()
			PrintColor("Thanks for playing !", "White")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		} else {
			ResetWordFromSave(save)
			loop = false
			ModeMenu(save, false)
		}
	}
}
