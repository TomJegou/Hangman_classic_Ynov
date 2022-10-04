package lib

import "fmt"

func Checktwice(letter string) {
	sliceChar := []string{}
	if IsIn(sliceChar, letter) {
		fmt.Print("You already try this letter ")
	} else {
		sliceChar = append(sliceChar, letter)
	}
}
