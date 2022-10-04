package lib

import (
	"bufio"
	"fmt"
	"os"
)

func DisplayAsciiWord(word string) {
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
		kIntager := int(word[k])
		for i := 0; i < 9*(kIntager)+1; i++ {
			for scanner.Scan() {
				if pos < 9*(kIntager)-1 {
					pos += 1
				} else if pos < 9*(kIntager)+9 {
					draw = append(draw, scanner.Text())
					draw = append(draw, "\n")
					pos += 1
					break
				}
			}
		}
	}
	for i := 0; i < len(draw); i++ {
		fmt.Print(draw[i])
	}
}
