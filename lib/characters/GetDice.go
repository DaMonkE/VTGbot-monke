package characters

import "reflect"

// GetDice returns the number of dice the character can roll
func GetDice(name, attr, skill string, mod int) (diceNum int) {
	// Open the character sheet
	var player Player
	err := LoadSheet(&player, name)
	if err != nil {
		return 0
	}
	// Determine the dice needed
	attrDice := int(reflect.Indirect(reflect.ValueOf(&player)).FieldByName(attr).Int())
	skillDice := int(reflect.Indirect(reflect.ValueOf(&player)).FieldByName(skill).Int())
	diceNum = attrDice + skillDice + mod
	return
}
