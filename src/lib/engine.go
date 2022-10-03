package lib

import (
	"fmt"
	"math/rand"
	"time"
)

func Engine(words []string) {
	Clear()
	attemptRemaining := 10
	word_to_guess := ChoseRandomWord(words)
	hiddenWord := ""
	numberLetterRevealed := len(word_to_guess)/2 - 1
	for i := 0; i < len(word_to_guess); i++ {
		hiddenWord += "_"
	}
	slice_byte_hidden := []byte(hiddenWord)
	for i := 0; i < numberLetterRevealed; i++ {
		indexLetterRevealed := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(slice_byte_hidden))
		DiscoverLetter(slice_byte_hidden, string(word_to_guess[indexLetterRevealed]), word_to_guess)
	}
	fmt.Println(word_to_guess)
	fmt.Println(RemainingLetter(slice_byte_hidden, word_to_guess))
	DisplayInput(slice_byte_hidden, attemptRemaining)
}
