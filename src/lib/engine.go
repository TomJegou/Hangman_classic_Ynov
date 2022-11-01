package lib

import "os"

/*
Function Engine is the main loop for the hangman game
*/
func Engine(isSave bool) {
	if isSave {
		isSave = false
		DicMenu(LoadSave(), true)
		return
	}
	var input string // Store the input player
	loop := true
	validOutput := true
	for loop {
		ClearConsole()
		if !validOutput { //display message in case of invalid output
			PrintColor("Invalid Input !\n\n", "White")
		} else {
			PrintColor("Welcome to Classic Hangman made by rLouis and jTom\n\n", "White")
		}
		PrintColor("[s]: Start new game", "Green")
		PrintColor("\n[q]: quit\n\n", "Red") // change print
		PrintColor("Choose: ", "White")
		input, validOutput = GetInputUser([]string{"s", "q"})
		if !validOutput {
			continue
		}
		if input == "s" {
			DicMenu(&Save{MaxError: 10, AttemptNumber: 0, Debug: false}, false) // calls the menu function
			loop = false
		} else {
			ClearConsole()
			PrintColor("Thanks for Playing !\n", "White")
			os.Exit(0)
		}
	}
}
