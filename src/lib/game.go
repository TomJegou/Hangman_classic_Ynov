package lib

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Game(lists_words []string, display_mod string) {
	Clear()
	debug_mod := false // Boolean to activate the debug mod
	maxError := 10
	numberError := 0
	word_to_guess := ChoseRandomWord(lists_words)
	hiddenWord := ""
	sliceAllChar := []string{}
	inputHistory := []string{}
	for i := 0; i < len(word_to_guess); i++ { // append all charcater into a slice in order to be read by the Isin function
		sliceAllChar = append(sliceAllChar, string(word_to_guess[i]))
	}
	// the programm will reveal n random letters in the word, where n is the len(word) / 2 - 1
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
	// Create a slice of the remaining letter to guess
	remainLetter := RemainingLetter(slice_byte_hidden, word_to_guess)
	found := true //boolean wich determine if the player found the word
	var input string
	attempt_number := 0
	invalid_ouput := false
	twice := false
	for len(remainLetter) > 0 {
		Clear()
		if numberError == 0 {
			PrintColor("Good Luck, you have ", "White")
			PrintColor("10 ", "Green")
			PrintColor("attempts.\n\n", "White")
		}
		attempt_number++
		if debug_mod { // output usefull variables in case of debug mod
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
				PrintColor("Invalid input, only one alphabetical character is supported as entry\n", "White")
				invalid_ouput = false
			} else if twice {
				PrintColor("Already try this letter\n", "White")
				twice = false
			} else {
				DisplayWrongLetter(numberError, maxError)
			}
			DisplayHangman(numberError)
		}
		DisplayModLetter(slice_byte_hidden, display_mod)
		fmt.Scanln(&input) // get the intput player
		// check the input validity
		if len(input) < 1 {
			numberError++
			invalid_ouput = true
			continue
		}
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
		//check if the input as been already played
		inputHistory, twice = Checktwice(input, inputHistory)
		if IsIn(sliceAllChar, input) && IsIn(remainLetter, input) { // in case of good answer
			DiscoverLetter(slice_byte_hidden, input, word_to_guess)          // reveal the letter
			remainLetter = RemainingLetter(slice_byte_hidden, word_to_guess) // reffresh the slice of remanig letter
		} else {
			numberError++
		}
	}
	// Display endgame message
	if found {
		Clear()
		DisplayHangman(numberError)
		DisplayModLetter(slice_byte_hidden, display_mod)
		PrintColor("\nCongrat !\nYou've found the word\nThe word was: "+word_to_guess+"\n\n", "Green")
	} else {
		DisplayHangman(numberError)
		DisplayModLetter(slice_byte_hidden, display_mod)
		PrintColor("\nYou didn't find the word !\nThe word was: "+word_to_guess+"\n\n", "Red")
	}
	//loop to ask the player to keep playing or not
	loop := true
	for loop {
		if invalid_ouput {
			PrintColor("Invalid output\n", "White")
			invalid_ouput = false
		}
		PrintColor("[c]continue [q]quit [b]Back\n\n", "White")
		PrintColor("Choose: ", "White")
		fmt.Scanln(&input)
		if input == "c" {
			Clear()
			PrintColor("Starting new game...", "White")
			time.Sleep(1 * time.Second)
			Game(lists_words, display_mod)
		} else if input == "q" {
			Clear()
			PrintColor("Thanks for playing !", "White")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		} else if input == "b" {
			Clear()
			loop = false
		} else {
			Clear()
			invalid_ouput = true
		}
	}
}
