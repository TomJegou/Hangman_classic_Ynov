package lib

import (
	"bufio"
	"fmt"
	"os"
)

func getpositioninalphabet(letter byte) int {
	return int(letter) - 97
}

func getline(line_number int) string {
	file, err := os.Open("../Templates/thinkertoy.txt")
	if err != nil {
		fmt.Println(err)
	}
	filescan := bufio.NewScanner(file)
	i := 1
	for filescan.Scan() {
		if i == line_number {
			return filescan.Text()[:len(filescan.Text())]
		} else {
			i++
		}
	}
	return ""
}

func DisplayStickLetter(t []byte) {
	for i := 0; i < 9; i++ {
		fmt.Println(getline(297 + i + getpositioninalphabet(t[0])*9))
	}
}
