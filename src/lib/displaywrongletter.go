package lib

import "fmt"

func DisplayWrongLetter(numberError, maxError int) {
	fmt.Printf("Not present in the word, %v attempts remaining\n", maxError-numberError)
}
