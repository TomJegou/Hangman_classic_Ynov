package lib

import (
	"bufio"
	"fmt"
	"os"
)

func DisplayInput(t []byte, numberError int) {
	currentStateWord := ""
	if numberError == 0 {
		fmt.Println("Good Luck, you have 10 attempts.")
		fmt.Println()
	}
	for i := 0; i < len(t); i++ {
		if i == len(t)-1 {
			currentStateWord += string(t[i])
		} else {
			currentStateWord += string(t[i]) + " "
		}
	}
	fmt.Println(currentStateWord + "\n")
	fmt.Print("Choose: ")
}

func DisplayHangman(attempt int) {
	file, err := os.Open("../Templates/hangman.txt")
	if err != nil {
		fmt.Println(err)
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
		fmt.Print(draw[i])
	}
}

func DisplayWrongLetter(numberError, maxError int) {
	fmt.Printf("Not present in the word, %v attempts remaining\n", maxError-numberError)
}