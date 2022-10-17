package lib

/*
Function Engine is the main loop for the hangman game
*/

func Engine(issave bool) {
	if issave {
		issave = false
		DicMenu(LoadSave(), true)
	}
	var input string // Store the input player
	loop := true
	validOutput := true
	for loop {
		ClearConsole()
		if !validOutput { //display message in case of invalid output
			PrintColor("Invalid output !\n\n", "White")
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
		} else {
			ClearConsole()
			PrintColor("Thanks for Playing !\n", "White")
			loop = false
		}
	}
}
