package lib

import (
	"strconv"
)

/*
Function that display the word in normal mod
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

func DisplayHangman(numberError int) {
	file := GetFileInLine("../Templates/hangman.txt")
	for i := 0; i < 8; i++ {
		if i != 7 {
			PrintColor(file[i+(numberError-1)*8]+"\n", "White")
		} else {
			PrintColor(file[i+(numberError-1)*8], "White")
		}
	}
}

func DisplayWrongLetter(numberError, maxError int) {
	PrintColor("Not present in the word, ", "White")
	PrintColor(strconv.Itoa(maxError-numberError), "Red")
	PrintColor(" attempts remaining\n", "White")
}

func DisplayModLetter(t []byte, template_mod string) {
	if template_mod == "0" {
		DisplayClassic(t)
		return
	}
	template_name := map[string]string{
		"1": "standard",
		"2": "thinkertoy"}
	for i := 1; i <= 9; i++ {
		l := ""
		for j := 0; j < len(t); j++ {
			l += GetFileInLine("../Templates/" + template_name[template_mod] + ".txt")[297+i+Getpositioninalphabet(t[j])*9]
			if j != len(t)-1 {
				l += GetFileInLine("../Templates/" + template_name[template_mod] + ".txt")[i]
			}
		}
		PrintColor(l+"\n", "White")
	}
	PrintColor("Choose: ", "White")
}
