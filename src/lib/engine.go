package lib

import (
	"fmt"
	"math/rand"
	"time"
)

func Engine(words []string) {
	Clear()
	debug_mod := true
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
	var input string
	attempt_number := 0
	for len(remainLetter) > 0 {
		attempt_number++
		Clear()
		if debug_mod {
			fmt.Println("Word to find: " + word_to_guess)
			fmt.Printf("Number error max: %v\n", maxError)
			fmt.Printf("Number error: %v\n", numberError)
			fmt.Printf("Attempt number: %v\n", attempt_number)
			if len(input) > 0 {
				fmt.Println("Last input player: " + input)
			}
			fmt.Println()
		}
		if numberError >= maxError {
			found = false
			break
		}
		if numberError > 0 {
			DisplayWrongLetter(numberError, maxError)
			DisplayHangman(numberError)
		}
		DisplayInput(slice_byte_hidden, numberError)
		fmt.Scanln(&input)
		if IsIn(sliceAllChar, input) && IsIn(remainLetter, input) {
			DiscoverLetter(slice_byte_hidden, input, word_to_guess)
			remainLetter = RemainingLetter(slice_byte_hidden, word_to_guess)
		} else {
			numberError++
		}
	}
	if found {
		Clear()
		DisplayHangman(numberError)
		fmt.Println("Congrat !")
		fmt.Println("You've found the word !")
		fmt.Println("The word was: " + word_to_guess)
	} else {
		DisplayHangman(numberError)
		fmt.Println("You didn't find the word !")
		fmt.Println("The word was: " + word_to_guess)
	}
}
