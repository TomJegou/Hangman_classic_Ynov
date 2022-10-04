package lib

import "fmt"

func Checktwice(letter string, t []string) []string {
	if IsIn(t, letter) {
		fmt.Println("You already try this letter ")
	} else {
		t = append(t, letter)
	}
	return t
}
