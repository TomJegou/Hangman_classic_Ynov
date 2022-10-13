package lib

import (
	"fmt"
	"os"
	"strings"
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
	var input string
	loop := true
	invalid_ouput := false
	for loop {
		if invalid_ouput {
			PrintColor("Invalid input\n\n", "White")
			invalid_ouput = false
		}
		PrintColor("Choose your mode\n\n", "White")
		keys := Listmap(save.TemplatesNames) // Get the key list in order to check if the next input is valid
		PrintColor("[b]: Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		fmt.Scanln(&input) // Scan and store the user input into the variable input
		input = strings.ToLower(input)
		if len(input) > 1 {
			invalid_ouput = false
			continue
		}
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
	invalid_ouput := false
	for loop {
		if invalid_ouput {
			PrintColor("Invalid input\n\n", "White")
			invalid_ouput = false
		}
		PrintColor("Choose your dictionnaries\n\n", "White")
		keys := Listmap(save.DictionnaryNames)
		PrintColor("[b]: Back\n\n", "Red")
		PrintColor("Choose: ", "White")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if len(input) > 1 {
			invalid_ouput = false
			continue
		}
		if input == "b" {
			loop = false
			Clear()
		} else if !IsIn(keys, input) {
			invalid_ouput = true
			Clear()
		} else {
			Clear()
			save.ListsWords = GetFileInLine("../dictionnaries/" + save.DictionnaryNames[input] + ".txt")
			save.WordToGess = ChoseRandomWord(save.ListsWords)
			MenuMode(save, false)
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
		}
		PrintColor("Hum it seems that there is a save, would you like to load it? Y/n: ", "white")
		fmt.Scanln(&input) // get the input player
		input = strings.ToLower(input)
		if len(input) > 1 {
			invalid_ouput = false
			continue
		}
		if input == "y" {
			LoadingBar(25)
			Engine(true)
			return
		} else if input == "n" {
			Engine(false)
			return
		} else {
			invalid_ouput = true
		}
	}
}

func EndgameMenu(save *Save, found bool) {
	// Display endgame message
	if found {
		Clear()
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
		if !validOutput {
			PrintColor("Invalid output\n", "White")
		}
		PrintColor("[c]continue [q]quit [b]Back\n\n", "White")
		PrintColor("Choose: ", "White")
		input, validOutput = GetInputUser([]string{"c", "q", "b"})
		if !validOutput {
			continue
		}
		if input == "c" {
			Clear()
			PrintColor("Starting new game...", "White")
			time.Sleep(1 * time.Second)
			ResetWordFromSave(save)
			Game(save, true)
		} else if input == "q" {
			Clear()
			PrintColor("Thanks for playing !", "White")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		} else if input == "b" {
			Clear()
			loop = false
			ResetWordFromSave(save)
		}
	}
}
