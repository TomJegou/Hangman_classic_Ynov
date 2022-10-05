package lib

import (
	"bufio"
	"fmt"
	"os"
)

func DisplayAsciiWord(word string) { //test
	for i := 0; i < len(word); i++ {
		w := []byte(string(word[i]))
		DisplayAsciiLetter(w)
	}
}

func DisplayAsciiLetter(word []byte) {
	file, err := os.Open("../Templates/standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	draw := []string{}
	scanner := bufio.NewScanner(file)
	pos := 0
	for k := 0; k < len(word); k++ {
		word[k] = word[k] - 32
		for i := 0; i < 9*(int(word[k]))+1; i++ {
			for scanner.Scan() {
				if pos < 9*(int(word[k]))-1 {
					pos += 1
				} else if pos <= 9*(int(word[k]))+9 {
					draw = append(draw, scanner.Text())
					draw = append(draw, "\n")
					pos += 1
					break
				}

			}
		}
	}
	for i := 0; i < len(draw); i++ {
		PrintColor(draw[i], "White")
	}
}
