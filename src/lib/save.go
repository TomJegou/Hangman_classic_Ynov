package lib

import (
	"encoding/json"
	"os"
)

type Save struct {
	NumberError      int
	MaxError         int
	InputHistory     []string
	CurrentStateWord []byte
	WordToGess       string
	DisplayMod       string
	ListsWords       []string
	TemplatesNames   map[string]string
	SliceAllChar     []string
}

func LoadSave() (Save, error) {
	var save Save
	bytevalue, err := os.ReadFile("saves/save.json")
	if err != nil {
		return save, err
	}
	json.Unmarshal(bytevalue, &save)
	return save, err
}

func SaveGame() {

}
