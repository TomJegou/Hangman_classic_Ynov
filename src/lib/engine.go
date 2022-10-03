package lib

import (
	"fmt"
	"math/rand"
	"time"
)

func Engine(words []string) {
	Clear()
	word_to_guess := ChosseRandomWord(words)
	hiddenWord := ""
	numberLetterRevealed := len(word_to_guess)/2 - 1
	for i := 0; i < len(word_to_guess); i++ {
		if i == len(word_to_guess)-1 {
			hiddenWord += "_"
		} else {
			hiddenWord += "_"
		}
	}
	slice_byte_hidden := []byte(hiddenWord)
	for i := 0; i < numberLetterRevealed; i++ {
		indexLetterRevealed := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(slice_byte_hidden))
		slice_byte_hidden[indexLetterRevealed] = word_to_guess[indexLetterRevealed]
	}
	finalHiddenWord := ""
	for i := 0; i < len(slice_byte_hidden); i++ {
		if i == len(slice_byte_hidden)-1 {
			finalHiddenWord += string(slice_byte_hidden[i])
		} else {
			finalHiddenWord += string(slice_byte_hidden[i]) + " "
		}
	}
	fmt.Println(finalHiddenWord)
}
