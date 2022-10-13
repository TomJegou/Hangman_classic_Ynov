package lib

import (
	"os"
)

/*
Function that scans a directory and return a map of all file's names in the dir as the value and an index as key
*/

func Scandir(pathdir string) map[string]string {
	data := map[string]string{}
	files, err := os.ReadDir(pathdir)
	if err != nil {
		PrintColor("Dictionnary Error", "Red")
	}
	for i := 0; i < len(files); i++ {
		data[string(byte(i+49))] = files[i].Name()[:len(files[i].Name())-4] // only add the name of the file without the .txt
	}
	return data
}
