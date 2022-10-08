package lib

import (
	"bufio"
	"os"
)

/*
File that contains all the get functions
*/

func GetFileInLine(file_path string) []string {
	data := []string{}
	file, err := os.Open(file_path)
	if err != nil {
		Clear()
		PrintColor("This file doesn't exist\n", "Red")
	} else {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
	}
	return data
}

func Getpositioninalphabet(letter byte) int {
	return int(letter) - 65
}

func Getline(line_number int, template_name string) string {
	return GetFileInLine("../Templates/" + template_name + ".txt")[line_number]
}
