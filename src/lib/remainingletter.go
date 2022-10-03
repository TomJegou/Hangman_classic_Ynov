package lib

func RemainingLetter(t []byte, completeWord string) []string {
	result := []string{}
	for i := 0; i < len(t); i++ {
		if !IsIn(result, string(t[i])) && string(t[i]) == "_" && !IsIn(result, string(completeWord[i])) {
			result = append(result, string(completeWord[i]))
		}
	}
	return result
}
