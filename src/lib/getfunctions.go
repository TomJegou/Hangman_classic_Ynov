package lib

import (
	"bufio"
	"os"
)

/*
Function that open a file and return a slice of all lines
*/

func GetFileInLine(file_path string) []string {
	data := []string{} // slice wich sotre all lines
	file, err := os.Open(file_path)
	if err != nil {
		Clear()
		PrintColor("This file doesn't exist\n", "Red")
		os.Exit(0)
	} else {
		defer file.Close() // close the file at the end of the function
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
	}
	return data
}

/*
Function that return the postition of the letter in the alphabet
*/

func Getpositioninalphabet(letter byte) int {
	return int(letter) - 65
}

func GetInputUser(options []string) (string, bool) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input, IsIn(options, input)
}
