package lib

// art from http://www.ascii-art.de/ascii/s/stickman.txt and https://ascii.co.uk/art/hangman
func EasterEgg(filename string) {
	file := GetFileInLine("../Templates/.foreasteregg/" + filename + ".txt")
	for i := 0; i < len(file); i++ {
		PrintColor(file[i]+"\n", "White")
	}
}
