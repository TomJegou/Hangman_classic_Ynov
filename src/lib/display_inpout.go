package lib

import "fmt"

func DisplayInput(currentStateWord string, attempt_remaining int) {
	if attempt_remaining == 10 {
		fmt.Println("Good Luck, you have 10 attempts.")
		fmt.Println()
	}
	fmt.Println(currentStateWord + "\n")
	fmt.Println("Choose: ")
}
