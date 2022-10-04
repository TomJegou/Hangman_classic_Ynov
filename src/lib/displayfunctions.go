package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func DisplayInput(t []byte, numberError int) {
	currentStateWord := ""
	if numberError == 0 {
		PrintColor("Good Luck, you have ", "White")
		PrintColor("10 ", "Green")
		PrintColor("attempts.\n\n", "White")
	}
	for i := 0; i < len(t); i++ {
		if i == len(t)-1 {
			currentStateWord += string(t[i])
		} else {
			currentStateWord += string(t[i]) + " "
		}
	}
	fmt.Println(currentStateWord + "\n")
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
	for i := 0; i < len(draw); i++ {
		PrintColor(draw[i], "White")
	}
}

func DisplayWrongLetter(numberError, maxError int) {
	PrintColor("Not present in the word, ", "White")
	PrintColor(strconv.Itoa(maxError-numberError), "Red")
	PrintColor(" attempts remaining\n", "White")
}
