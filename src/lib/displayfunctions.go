package lib

import (
	"strconv"
)

/*
Function that displays the word in normal mod
it takes as parameter a table of bytes
*/

func DisplayClassic(t []byte) {
	currentStateWord := ""
	for i := 0; i < len(t); i++ {
		if i == len(t)-1 {
			currentStateWord += string(t[i])
		} else {
			currentStateWord += string(t[i]) + " "
		}
	}
	PrintColor(currentStateWord+"\n", "White")
	PrintColor("Choose: ", "White")
}

/*
Function that displays the current state of the hangman according to the number of error
*/

func DisplayHangman(numberError int) {
	file := GetFileInLine("../Templates/hangman/hangman.txt")
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
* 0: classic mod
* 1: standard mod
* 2: thinkertoy mod
*/

func DisplayModLetter(t []byte, template_mod string) {
	if template_mod == "0" { //if it's classic mod, calls the classicdisplay function
		DisplayClassic(t)
		return
	}
	template_name := Scandir("../Templates/policies/")
	for i := 1; i <= 9; i++ { // loop each line
		l := ""                       // store line
		for j := 0; j < len(t); j++ { // loop for each character
			l += GetFileInLine("../Templates/policies/" + template_name[template_mod] + ".txt")[297+i+Getpositioninalphabet(t[j])*9]
			if j != len(t)-1 {
				l += GetFileInLine("../Templates/policies/" + template_name[template_mod] + ".txt")[i] // add a space after the letter except the last one
			}
		}
		PrintColor(l+"\n", "White")
	}
	PrintColor("Choose: ", "White") //Input display
}
