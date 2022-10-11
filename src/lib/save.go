package lib

import (
	"encoding/json"
	"os"
)

type Save struct {
	Debug                                                bool
	NumberError, MaxError, Attempt_number                int
	InputHistory, SliceAllChar, ListsWords, RemainLetter []string
	CurrentStateWord                                     []byte
	WordToGess, DisplayMode                              string
	TemplatesNames                                       map[string]string
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

func SaveGame(save Save) {
	byteValue, err := json.MarshalIndent(save, "", "    ")
	if err != nil {
		PrintColor("Faild to save the game", "Red")
	}
	os.WriteFile("save/save.json", byteValue, 0644)
}
