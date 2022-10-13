package lib

func EasterEgg(filename string) {
	file := GetFileInLine("../Templates/.foreasteregg/" + filename + ".txt")
	for i := 0; i < len(file); i++ {
		PrintColor(file[i]+"\n", "White")
	}
}
