package lib

import "fmt"

/*
Function that clears the terminal output
We use the ANSI Escape Sequences
*/
func Clear() {
	fmt.Print("\033[H\033[2J") // the \033[h is used to moove the cursor to the 0 0 positon and \033[j to clear the terminal
}
