package lib

func DiscoverLetter(t []byte, letter string, word string) {
	for i := 0; i < len(word); i++ {
		if string(word[i]) == letter {
			t[i] = word[i]
		}
	}
}
