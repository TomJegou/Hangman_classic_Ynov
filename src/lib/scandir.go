package lib

import (
	"os"
)

func Scandir(pathdir string) map[string]string {
	data := map[string]string{}
	files, err := os.ReadDir(pathdir)
	if err != nil {
		PrintColor("Dictionnary Error", "Red")
	}
	for i := 0; i < len(files); i++ {
		data[string(byte(i+49))] = files[i].Name()[:len(files[i].Name())-4]
	}
	return data
}
