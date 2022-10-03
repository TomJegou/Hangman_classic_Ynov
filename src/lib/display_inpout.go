package lib

import "fmt"

func DisplayInput(t []byte, attempt_remaining int) {
	currentStateWord := ""
	if attempt_remaining == 10 {
		fmt.Println("Good Luck, you have 10 attempts.")
		fmt.Println()
	}
	for i := 0; i < len(t); i++ {
		if i == len(t)-1 {
			currentStateWord += string(t[i])
		} else {
			currentStateWord += string(t[i]) + " "
		}
	}
	fmt.Println(currentStateWord + "\n")
	fmt.Println("Choose: ")
}
