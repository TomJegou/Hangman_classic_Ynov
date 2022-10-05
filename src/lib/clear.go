package lib

import "fmt"

/*
The function clear is use to clear the consol to play
For that we used the ANSI Escape Sequences
*/
func Clear() {
	fmt.Print("\033[H\033[2J") // the \033[h is used to moove the cursor to the 0 0 positon and \033[j to clear the consol
}
