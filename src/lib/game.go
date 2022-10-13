package lib

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
Game function, main loop of the game with all the settings already set by the player before the function call
*/

func Game(save *Save, new bool) {
	Clear()
	if new {
		hiddenWord := ""
		for i := 0; i < len(save.WordToGess); i++ { // append all charcater into a slice in order to be read by the Isin function
			save.SliceAllChar = append(save.SliceAllChar, string(save.WordToGess[i]))
		}
		// the programm will reveal n random letters in the word, where n is the len(word) / 2 - 1
		numberLetterRevealed := len(save.WordToGess)/2 - 1
		for i := 0; i < len(save.WordToGess); i++ {
			hiddenWord += "_"
		}
		save.CurrentStateWord = []byte(hiddenWord)
		for i := 0; i < numberLetterRevealed; i++ {
			indexLetterRevealed := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(save.CurrentStateWord))
			DiscoverLetter(save.CurrentStateWord, string(save.WordToGess[indexLetterRevealed]), save.WordToGess)
			save.InputHistory, _ = Checktwice(string(save.WordToGess[indexLetterRevealed]), save.InputHistory)
		}
		// Create a slice of the remaining letter to guess
		save.RemainLetter = RemainingLetter(save.CurrentStateWord, save.WordToGess)
	}
	SaveGame(save)
	found := true //boolean wich determine if the player found the word
	var input string
	invalid_ouput := false
	twice := false
	for len(save.RemainLetter) > 0 {
		Clear()
		if save.NumberError == 0 {
			PrintColor("Good Luck, you have ", "White")
			PrintColor("10 ", "Green")
			PrintColor("attempts.\n\n", "White")
		}
		save.AttemptNumber++
		if save.Debug { // output usefull variables in case of debug mod
			fmt.Println("Word to find: " + save.WordToGess)
			fmt.Printf("Number error max: %v\n", save.MaxError)
			fmt.Printf("Number error: %v\n", save.NumberError)
			fmt.Printf("Attempt number: %v\n", save.AttemptNumber)
			if len(save.InputHistory) > 0 {
				fmt.Print("Input history")
				fmt.Println(save.InputHistory)
			}
			if len(input) > 0 {
				fmt.Println("Last input player: " + input)
			}
			fmt.Println()
		}
		if save.NumberError >= save.MaxError {
			found = false
			break
		}
		if save.NumberError > 0 {
			if invalid_ouput {
				PrintColor("Invalid input, only one alphabetical character is supported as entry\n", "White")
				invalid_ouput = false
			} else if twice {
				PrintColor("Already try this letter\n", "White")
				twice = false
			} else {
				DisplayWrongLetter(save.NumberError, save.MaxError)
			}
			DisplayHangman(save.NumberError)
		}
		DisplayModLetter(save, "White", true, true)
		fmt.Scanln(&input) // get the intput player
		/*all easter egg*/
		if input == "noHolidays" {
			EasterEgg("island")
			PrintColor("Press enter to continue \n", "Green")
			fmt.Scanln()
			continue
		}
		/*end easter egg*/
		// check the input validity
		if input == "STOP" {
			SaveGame(save)
			Clear()
			PrintColor("Thanks for playing !", "White")
			os.Exit(0)
		}
		if len(input) < 1 {
			save.NumberError++
			invalid_ouput = true
			continue
		}
		if len(input) > 1 {
			if input == save.WordToGess {
				break
			} else {
				save.NumberError += 2
				continue
			}
		}
		if input >= "0" && input <= "9" {
			save.NumberError++
			invalid_ouput = true
			continue
		}
		//check if the input as been already played
		save.InputHistory, twice = Checktwice(input, save.InputHistory)
		if IsIn(save.SliceAllChar, input) && IsIn(save.RemainLetter, input) { // in case of good answer
			DiscoverLetter(save.CurrentStateWord, input, save.WordToGess)               // reveal the letter
			save.RemainLetter = RemainingLetter(save.CurrentStateWord, save.WordToGess) // reffresh the slice of remanig letter
		} else {
			save.NumberError++
		}
		SaveGame(save)
	}
	// Display endgame message
	if found {
		Clear()
		if save.NumberError > 0 {
			DisplayHangman(save.NumberError)
		}
		DisplayModLetter(save, "White", true, true)
		PrintColor("\nCongrat !\nYou've found the word\nThe word was: "+save.WordToGess+"\n\n", "Green")
	} else {
		DisplayHangman(save.NumberError)
		DisplayModLetter(save, "White", true, true)
		PrintColor("\nYou didn't find the word !\nThe word was: "+save.WordToGess+"\n\n", "Red")
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
		input = strings.ToLower(input)
		if input == "c" {
			Clear()
			PrintColor("Starting new game...", "White")
			time.Sleep(1 * time.Second)
			ResetWordFromSave(save)
			Game(save, true)
		} else if input == "q" {
			Clear()
			PrintColor("Thanks for playing !", "White")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		} else if input == "b" {
			Clear()
			loop = false
			ResetWordFromSave(save)
		} else {
			Clear()
			invalid_ouput = true
		}
	}
}
