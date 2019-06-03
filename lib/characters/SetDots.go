package characters

import "reflect"

// SetDots loads a character sheet, sets the attribute/skill value, and saves the sheet
func SetDots(name, slot string, num int64) (err error) {
	// Open the character sheet
	var player Player
	err = LoadSheet(&player, name)
	if err != nil {
		return
	}
	// Assign dots to the designated skill
	reflect.ValueOf(&player).Elem().FieldByName(slot).SetInt(num)
	// Save the new sheet
	err = SaveSheet(&player, name)
	if err != nil {
		return
	}
	return
}
