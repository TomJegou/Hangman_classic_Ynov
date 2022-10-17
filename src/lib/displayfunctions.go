package lib

import (
	"fmt"
	"strconv"
	"time"
)

/*
Function used to print texts in color
We use the ANSI Escape Sequences
*/
func PrintColor(s, color string) {
	color_map := map[string]string{ //map to store the corresponding color's ansi code
		"Black":  "\033[0;30m",
		"Red":    "\033[0;31m",
		"Green":  "\033[0;32m",
		"Yellow": "\033[0;33m",
		"Blue":   "\033[0;34m",
		"Purple": "\033[0;35m",
		"Cyan":   "\033[0;36m",
		"White":  "\033[0;37m"}
	fmt.Printf("%v%v", color_map[color], s)
	fmt.Printf("%v", color_map["White"])
}

/*
Function that displays the current state of the hangman according to the number of errors
*/
func DisplayHangman(numberError int) {
	file := GetFileLineInSlice("../Templates/hangman/hangman.txt")
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
func DisplayModLetter(save *Save, color string, withSpace, input bool) {
	pathDisplay := "../Templates/policies/" + save.TemplatesNames[save.DisplayMode] + ".txt"
	multiple := CalculateLinesBtwChar(pathDisplay)
	for i := 1; i <= multiple; i++ { // loop each line
		l := ""                                           // store line
		for j := 0; j < len(save.CurrentStateWord); j++ { // loop for each character
			l += GetFileLineInSlice(pathDisplay)[33*multiple+i+Getpositioninalphabet(save.CurrentStateWord[j])*multiple]
			if j != len(save.CurrentStateWord)-1 && withSpace {
				l += " " // add a space after the letter except the last one
			}
		}
		PrintColor(l+"\n", "White")
	}
	if input {
		PrintColor("Choose: ", "White") //Input display
	}
}

/*
Func that makes a loding bar

[░░░░░░░░░░░░░░░░░░░░░░░░░░░]

[███████████░░░░░░░░░░░░░░░░]

[███████████████████████████]

░ 176
█ 219
*/
func LoadingBar(t time.Duration) { //t = Duration in millisecond
	ldbar := "[░░░░░░░░░░░░░░░░░░░░░░░░░░░]"
	ldbarRune := []rune(ldbar)
	for i := 1; i < len(ldbarRune)-1; i++ {
		ClearConsole()
		ldbarRune[i] = rune('█')
		PrintColor(string(ldbarRune), "Red")
		time.Sleep(t * time.Millisecond) //pauses for t milliseconds
	}
	ClearConsole()
	PrintColor(string(ldbarRune), "Green")
	time.Sleep(500 * time.Millisecond)
	ClearConsole()
	PrintColor("Loading Complet\n", "Green")
	time.Sleep(500 * time.Millisecond)
}
