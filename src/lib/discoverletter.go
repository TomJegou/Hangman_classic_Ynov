package lib

/*
Fuction used to change the "_" to the corresponding letter from the word to guess
*/

func DiscoverLetter(t []byte, letter string, word string) {
	for i := 0; i < len(word); i++ {
		if string(word[i]) == letter {
			t[i] = word[i]
		}
	}
}
