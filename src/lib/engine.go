package lib

import (
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
	letterRemaining := []string{}
	for i := 0; i < len(slice_byte_hidden); i++ {
		if string(slice_byte_hidden[i]) == "_" {
			letterRemaining = append(letterRemaining, string(word_to_guess[i]))
		}
	}
	currentStateWord := finalHiddenWord
	DisplayInput(currentStateWord, attemptRemaining)

}
