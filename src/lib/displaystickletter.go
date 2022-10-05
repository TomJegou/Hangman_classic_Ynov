package lib

import (
	"bufio"
	"fmt"
	"os"
)

func getline(line_number int, file *os.File) string {
	filescan := bufio.NewScanner(file)
	filescan.Split(bufio.ScanLines)
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

func DisplayStickLetter() {
	file, err := os.Open("../Templates/thinkertoy.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getline(13, file))
}
