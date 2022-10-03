package lib

import "strings"

func RemainingLetter(currentStateWord string, completeWord string, t []string) {
	temp_slice := strings.Split(currentStateWord, " ")
	for i := 0; i < len(temp_slice); i++ {
		if string(temp_slice[i]) == " " {
			continue
		}
		if !IsIn(t, string(temp_slice[i])) && string(temp_slice[i]) == "_" {
			t = append(t, string(completeWord[i]))
		}
	}
}
