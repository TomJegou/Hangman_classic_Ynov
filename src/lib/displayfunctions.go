package lib

import (
	"strconv"
)

/*
Function that displays the current state of the hangman according to the number of error
*/

func DisplayHangman(numberError int) {
	file := GetFileInLine("../Templates/hangman/hangman.txt")
	if numberError > 10 {
		DisplayHangman(10)
		return
	}
	for i := 0; i < 8; i++ {
		if i != 7 {
			PrintColor(file[i+(numberError-1)*8]+"\n", "White")
		} else {
			PrintColor(file[i+(numberError-1)*8], "White")
		}
	}
}

/*
Function that displays a message in case of incorrect letter
*/

func DisplayWrongLetter(numberError, maxError int) {
	PrintColor("Not present in the word, ", "White")
	PrintColor(strconv.Itoa(maxError-numberError), "Red")
	PrintColor(" attempts remaining\n", "White")
}

/*
function that displays the word according to the template mod
*/

func DisplayModLetter(save *Save, color string, withSpace bool) {
	pathDisplay := "../Templates/policies/" + save.TemplatesNames[save.DisplayMode] + ".txt"
	multiple := CalculateMultiple(pathDisplay)
	for i := 1; i <= multiple; i++ { // loop each line
		l := ""                                           // store line
		for j := 0; j < len(save.CurrentStateWord); j++ { // loop for each character
			l += GetFileInLine(pathDisplay)[33*multiple+i+Getpositioninalphabet(save.CurrentStateWord[j])*multiple]
			if j != len(save.CurrentStateWord)-1 && withSpace {
				l += " " // add a space after the letter except the last one
			}
		}
		PrintColor(l+"\n", "White")
	}
	PrintColor("Choose: ", "White") //Input display
}
