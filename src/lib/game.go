package lib

import (
	"fmt"
	"math/rand"
	"time"
)

func Game(lists_words []string) {
	Clear()
	debug_mod := true
	maxError := 10
	numberError := 0
	word_to_guess := ChoseRandomWord(lists_words)
	hiddenWord := ""
	sliceAllChar := []string{}
	inputHistory := []string{}
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
		inputHistory, _ = Checktwice(string(word_to_guess[indexLetterRevealed]), inputHistory)
	}
	remainLetter := RemainingLetter(slice_byte_hidden, word_to_guess)
	found := true
	var input string
	attempt_number := 0
	invalid_ouput := false
	twice := false
	for len(remainLetter) > 0 {
		attempt_number++
		Clear()
		if debug_mod {
			fmt.Println("Word to find: " + word_to_guess)
			fmt.Printf("Number error max: %v\n", maxError)
			fmt.Printf("Number error: %v\n", numberError)
			fmt.Printf("Attempt number: %v\n", attempt_number)
			if len(inputHistory) > 0 {
				fmt.Print("Input history")
				fmt.Println(inputHistory)
			}
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
			if invalid_ouput {
				PrintColor("Invalid input, only one alphabetical character is supported in entry", "White")
				invalid_ouput = false
			} else if twice {
				PrintColor("Already try this letter", "White")
				twice = false
			} else {
				DisplayWrongLetter(numberError, maxError)
			}
			DisplayHangman(numberError)
		}
		DisplayInput(slice_byte_hidden, numberError)
		fmt.Scanln(&input)
		// if len(input) < 1 {
		// 	numberError++
		// 	invalid_ouput = true
		// 	continue
		// }
		if len(input) > 1 {
			if input == word_to_guess {
				break
			} else {
				numberError += 2
				continue
			}
		}
		if input >= "0" && input <= "9" {
			numberError++
			invalid_ouput = true
			continue
		}
		inputHistory, twice = Checktwice(input, inputHistory)
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
		PrintColor("Congrat !\nYou've found the word\nThe word was: "+word_to_guess+"\n\n", "Green")
	} else {
		DisplayHangman(numberError)
		PrintColor("You didn't find the word !\nThe word was: "+word_to_guess+"\n\n", "Red")
	}
}
