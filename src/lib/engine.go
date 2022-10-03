package lib

import (
	"fmt"
	"math/rand"
	"time"
)

func Engine(words []string) {
	Clear()
	maxError := 10
	numberError := 0
	word_to_guess := ChoseRandomWord(words)
	hiddenWord := ""
	sliceAllChar := []string{}
	for i := 0; i < len(word_to_guess); i++ {
		sliceAllChar = append(sliceAllChar, string(word_to_guess[i]))
	}
	numberLetterRevealed := len(word_to_guess)/2 - 1
	for i := 0; i < len(word_to_guess); i++ {
		hiddenWord += "_"
	}
	slice_byte_hidden := []byte(hiddenWord)
	for i := 0; i < numberLetterRevealed; i++ {
		indexLetterRevealed := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(slice_byte_hidden))
		DiscoverLetter(slice_byte_hidden, string(word_to_guess[indexLetterRevealed]), word_to_guess)
	}
	remainLetter := RemainingLetter(slice_byte_hidden, word_to_guess)
	found := true
	for len(remainLetter) > 0 {
		if numberError >= maxError {
			found = false
			break
		}
		fmt.Println(word_to_guess) // A enlver
		DisplayInput(slice_byte_hidden, numberError)
		var input string
		fmt.Scanln(&input)
		if IsIn(sliceAllChar, input) && IsIn(remainLetter, input) {
			DiscoverLetter(slice_byte_hidden, input, word_to_guess)
			remainLetter = RemainingLetter(slice_byte_hidden, word_to_guess)
		} else {
			numberError++
			DisplayWrongLetter(numberError, maxError)
			DisplayHangman(numberError)
		}
	}
	if found {
		fmt.Println()
		fmt.Println("Congrat !")
	} else {
		fmt.Println("You didn't find the word")
	}
}
