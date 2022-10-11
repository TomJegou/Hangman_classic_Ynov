package lib

import (
	"encoding/json"
	"os"
)

type Save struct {
	Debug                                                bool
	NumberError, MaxError, AttemptNumber                 int
	InputHistory, SliceAllChar, ListsWords, RemainLetter []string
	CurrentStateWord                                     []byte
	WordToGess, DisplayMode                              string
	TemplatesNames, DictionnaryNames                     map[string]string
}

func LoadSave() *Save {
	save := &Save{}
	bytevalue, err := os.ReadFile("saves/save.json")
	if err != nil {
		return save
	}
	json.Unmarshal(bytevalue, &save)
	return save
}

func SaveGame(save *Save) {
	byteValue, err := json.MarshalIndent(save, "", "    ")
	if err != nil {
		PrintColor("Faild to save the game", "Red")
	}
	os.WriteFile("saves/save.json", byteValue, 0644)
}

func ResetWordFromSave(save *Save) {
	save.AttemptNumber = 0
	save.NumberError = 0
	save.SliceAllChar = []string{}
	save.InputHistory = []string{}
	save.WordToGess = ChoseRandomWord(save.ListsWords)
}
