package characters

import "reflect"

// GetDots retrieves the number of dots character has in the indicated slot
func GetDots(name, attr string) (num int, err error) {
	// Open the character sheet
	var player Player
	err = LoadSheet(&player, name)
	if err != nil {
		return 0, err
	}
	// Read the value
	num = int(reflect.Indirect(reflect.ValueOf(&player)).FieldByName(attr).Int())
	return
}
