package characters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// LoadSheet reads the player data from a json file
func LoadSheet(player *Player, name string) (err error) {
	absPath, _ := filepath.Abs(fmt.Sprintf("../sheet/%v.json", name))
	jsonFile, err := os.Open(absPath)
	if err != nil {
		return
	}
	defer jsonFile.Close()
	sheetData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(sheetData, player)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
