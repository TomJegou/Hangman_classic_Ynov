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
	for i := 0 + attempt; i <= (9 * attempt); i++ {
		for scanner.Scan() {
			draw = append(draw, scanner.Text())
			break
		}
	}
	fmt.Print(draw)
}
