package characters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// SaveSheet saves the player info into a json file
func SaveSheet(player *Player, name string) (err error) {
	absPath, _ := filepath.Abs(fmt.Sprintf("../sheet/%v.json", name))
	sheetByte, err := json.Marshal(player)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(absPath, sheetByte, 0777)
	if err != nil {
		return
	}
	return
}
