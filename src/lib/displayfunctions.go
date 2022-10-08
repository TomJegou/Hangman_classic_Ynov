package lib

import (
	"bufio"
	"os"
	"strconv"
)

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

func DisplayHangman(attempt int) {
	file, err := os.Open("../Templates/hangman.txt")
	if err != nil {
		PrintColor("File doesn't exist", "Red")
	}
	draw := []string{}
	scanner := bufio.NewScanner(file)
	pos := 0
	for i := 0; i <= (8*attempt)+1; i++ {
		for scanner.Scan() {
			if pos < 8*(attempt-1) {
				pos += 1
				break
			} else if pos < 8*attempt {
				draw = append(draw, scanner.Text())
				draw = append(draw, "\n")
				pos += 1
				break
			}
		}
	}
	for i := 0; i < len(draw)-1; i++ {
		PrintColor(draw[i], "White")
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
			l += Getline(297+i+Getpositioninalphabet(t[j])*9, template_name[template_mod])
			if j != len(t)-1 {
				l += Getline(i, template_name[template_mod])
			}
		}
		PrintColor(l+"\n", "White")
	}
	PrintColor("Choose: ", "White")
}
