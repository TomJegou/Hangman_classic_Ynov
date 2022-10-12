package lib

import (
	"encoding/json"
	"os"
)

/*
Data Save structure
*/

type Save struct {
	Debug                                                bool
	NumberError, MaxError, AttemptNumber                 int
	InputHistory, SliceAllChar, ListsWords, RemainLetter []string
	CurrentStateWord                                     []byte
	WordToGess, DisplayMode                              string
	TemplatesNames, DictionnaryNames                     map[string]string
}

/*
Function that load a Save
*/

func LoadSave() *Save {
	save := &Save{}
	bytevalue, err := os.ReadFile("saves/save.json")
	if err != nil {
		return save
	}
	json.Unmarshal(bytevalue, &save)
	return save
}

/*
Function that save a game
*/

func SaveGame(save *Save) {
	byteValue, err := json.MarshalIndent(save, "", "    ")
	if err != nil {
		PrintColor("Faild to save the game", "Red")
	}
	os.WriteFile("saves/save.json", byteValue, 0644)
}

/*
Function that reset the word ,nb error , attempt ,from the save
*/

func ResetWordFromSave(save *Save) {
	save.AttemptNumber = 0
	save.NumberError = 0
	save.SliceAllChar = []string{}
	save.InputHistory = []string{}
	save.WordToGess = ChoseRandomWord(save.ListsWords)
}
