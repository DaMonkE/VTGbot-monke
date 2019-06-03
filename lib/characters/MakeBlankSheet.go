package characters

import (
	"fmt"
	"os"
	"path/filepath"
)

// MakeBlankSheet initializes a json file for the player character
func MakeBlankSheet(name string) string {
	// Check if a sheet already exists
	filename, _ := filepath.Abs(fmt.Sprintf("../sheet/%v.json", name))
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// Create the path to make the sheet
		player := Player{Name: name}
		err = SaveSheet(&player, name)
	} else {
		return fmt.Sprintf("Sheet already exists for %v, please use different name...", name)
	}
	return fmt.Sprintf("Sheet created for %v, please use /set to fill in the dots...", name)
}
