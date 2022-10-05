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
			return filescan.Text()
		} else {
			i++
		}
	}
	return ""
}

func DisplayStickLetter(t []byte) {
	for i := 0; i < 9; i++ {
		l := ""
		for j := 0; j < len(t); j++ {
			l += getline(297 + i + getpositioninalphabet(t[j])*9)
		}
		fmt.Println(l)
	}
}
