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

func LoadSave() Save {
	var save Save
	bytevalue, err := os.ReadFile("saves/save.json")
	if err != nil {
		return save
	}
	json.Unmarshal(bytevalue, &save)
	return save
}

func SaveGame(
	numberError int,
	maxError int,
	inputHistory []string,
	currentStateWord []byte,
	wordToGess string,
	displayMod string,
	listsWords []string,
	templatesNames map[string]string,
	sliceAllChar []string,
	save *Save) {
	save.NumberError = numberError
	save.MaxError = maxError
	save.InputHistory = inputHistory
	save.CurrentStateWord = currentStateWord
	save.WordToGess = wordToGess
	save.DisplayMod = displayMod
	save.ListsWords = listsWords
	save.TemplatesNames = templatesNames
	save.SliceAllChar = sliceAllChar
	byteValue, err := json.MarshalIndent(save, "", "    ")
	if err != nil {
		PrintColor("Faild to save the game", "Red")
	}
	os.WriteFile("save/save.json", byteValue, 0644)
}
