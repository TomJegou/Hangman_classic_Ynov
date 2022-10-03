package lib

import "fmt"

func Engine(words []string) {
	word_to_guess := ChosseRandomWord(words)
	fmt.Println(word_to_guess)
}
