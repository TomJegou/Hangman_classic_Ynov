package lib

import (
	"bufio"
	"fmt"
	"os"
)

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
			if attempt == 1 {
				draw = append(draw, scanner.Text())
				break
			}
			if pos < 8 {
				pos += 1
				break
			}
			if pos < 8*attempt {
				draw = append(draw, scanner.Text())
				draw = append(draw, "\n")
				break
			}
		}
	}
	fmt.Print(draw)
}
